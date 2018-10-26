package tagFilters

import (
	"regexp"
	"sync"
	"time"

	"github.com/Laisky/go-concator/libs"
	utils "github.com/Laisky/go-utils"
	"go.uber.org/zap"
)

type ConcatorCfg struct {
	Tag        string
	OutChan    chan<- *libs.FluentMsg
	MsgKey     string
	Identifier string
	PMsgPool   *sync.Pool
	Regexp     *regexp.Regexp
}

// Concator work for one tag, contains many identifier("container_id")
type Concator struct {
	*ConcatorCfg
	slot   map[string]*PendingMsg
	maxLen int
}

// PendingMsg is the message wait tobe concatenate
type PendingMsg struct {
	msg   *libs.FluentMsg
	lastT time.Time
}

// NewConcator create new Concator
func NewConcator(cfg *ConcatorCfg) *Concator {
	utils.Logger.Debug("create new concator",
		zap.String("tag", cfg.Tag),
		zap.String("identifier", cfg.Identifier),
		zap.String("msgKey", cfg.MsgKey))

	return &Concator{
		ConcatorCfg: cfg,
		slot:        map[string]*PendingMsg{},
		maxLen:      utils.Settings.GetInt("settings.max_msg_length"),
	}
}

// Run starting Concator to concatenate messages,
// you should not run concator directly,
// it's better to create and run Concator by ConcatorFactory
//
// TODO: concator for each tag now,
// maybe set one concator for each identifier in the future for better performance
func (c *Concator) Run(inChan <-chan *libs.FluentMsg) {
	var (
		msg        *libs.FluentMsg
		pmsg       *PendingMsg
		identifier string
		log        []byte
		ok         bool

		now             time.Time
		initWaitTs      = 20 * time.Millisecond
		maxWaitTs       = 500 * time.Millisecond
		waitTs          = initWaitTs
		nWaits          = 0
		nWaitsToDouble  = 2
		concatTimeoutTs = 5 * time.Second
		timer           = libs.NewTimer(libs.NewTimerConfig(initWaitTs, maxWaitTs, waitTs, concatTimeoutTs, nWaits, nWaitsToDouble))
	)

	for {
		select {
		case msg = <-inChan:
			now = time.Now()
			timer.Reset(now)

			// unknown identifier
			switch msg.Message[c.Identifier].(type) {
			case []byte:
				identifier = string(msg.Message[c.Identifier].([]byte))
			case string:
				identifier = msg.Message[c.Identifier].(string)
			default:
				utils.Logger.Warn("unknown identifier or unknown type",
					zap.String("tag", msg.Tag),
					zap.String("identifier", c.Identifier))
				c.OutChan <- msg
				continue
			}

			// unknon msg key
			switch msg.Message[c.MsgKey].(type) {
			case []byte:
				log = msg.Message[c.MsgKey].([]byte)
			case string:
				log = []byte(msg.Message[c.MsgKey].(string))
			default:
				utils.Logger.Warn("unknown msg key or unknown type",
					zap.String("tag", msg.Tag),
					zap.String("msg_key", c.MsgKey))
				c.OutChan <- msg
				continue
			}

			pmsg, ok = c.slot[identifier]
			// new identifier
			if !ok {
				utils.Logger.Debug("got new identifier", zap.String("identifier", identifier), zap.ByteString("log", msg.Message[c.MsgKey].([]byte)))
				pmsg = c.PMsgPool.Get().(*PendingMsg)
				pmsg.lastT = now
				pmsg.msg = msg
				c.slot[identifier] = pmsg
				continue
			}

			// old identifer
			if c.Regexp.Match(log) { // new line
				utils.Logger.Debug("got new line",
					zap.ByteString("log", msg.Message[c.MsgKey].([]byte)),
					zap.String("tag", msg.Tag))
				c.OutChan <- c.slot[identifier].msg
				c.slot[identifier].msg = msg
				c.slot[identifier].lastT = now
				continue
			}

			// need to concat
			utils.Logger.Debug("concat lines", zap.ByteString("log", msg.Message[c.MsgKey].([]byte)))
			c.slot[identifier].msg.Message[c.MsgKey] =
				append(c.slot[identifier].msg.Message[c.MsgKey].([]byte), '\n')
			c.slot[identifier].msg.Message[c.MsgKey] =
				append(c.slot[identifier].msg.Message[c.MsgKey].([]byte), msg.Message[c.MsgKey].([]byte)...)
			if c.slot[identifier].msg.ExtIds == nil {
				c.slot[identifier].msg.ExtIds = []int64{} // create ids, wait to append tail-msg's id
			}
			c.slot[identifier].msg.ExtIds = append(c.slot[identifier].msg.ExtIds, msg.Id)
			c.slot[identifier].lastT = now

			// too long to send
			if len(c.slot[identifier].msg.Message[c.MsgKey].([]byte)) >= c.maxLen {
				utils.Logger.Debug("too long to send", zap.String("msgKey", c.MsgKey), zap.String("tag", msg.Tag))
				c.OutChan <- c.slot[identifier].msg
				c.PMsgPool.Put(c.slot[identifier])
				delete(c.slot, identifier)
			}

		default: // check timeout
			now = time.Now()
			for identifier, pmsg = range c.slot {
				if now.Sub(pmsg.lastT) > concatTimeoutTs { // timeout to flush
					utils.Logger.Debug("timeout flush", zap.ByteString("log", pmsg.msg.Message[c.MsgKey].([]byte)))
					c.OutChan <- pmsg.msg
					c.PMsgPool.Put(pmsg)
					delete(c.slot, identifier)
				}
			}

			timer.Sleep()
		}
	}
}

type ConcatorFactCfg struct {
	ConcatorCfgs map[string]*libs.ConcatorTagCfg
}

// ConcatorFactory can spawn new Concator
type ConcatorFactory struct {
	*ConcatorFactCfg
	pMsgPool *sync.Pool
}

// NewConcatorFact create new ConcatorFactory
func NewConcatorFact(cfg *ConcatorFactCfg) *ConcatorFactory {
	utils.Logger.Info("create concatorFactory")
	return &ConcatorFactory{
		ConcatorFactCfg: cfg,
		pMsgPool: &sync.Pool{
			New: func() interface{} {
				return &PendingMsg{}
			},
		},
	}
}

func (cf *ConcatorFactory) GetName() string {
	return "concator_tagfilter"
}

func (cf *ConcatorFactory) IsTagSupported(tag string) bool {
	_, ok := cf.ConcatorCfgs[tag]
	return ok
}

// Spawn create and run new Concator for new tag
func (cf *ConcatorFactory) Spawn(tag string, outChan chan<- *libs.FluentMsg) chan<- *libs.FluentMsg {
	inChan := make(chan *libs.FluentMsg, 1000)
	concator := NewConcator(&ConcatorCfg{
		Tag:        tag,
		OutChan:    outChan,
		MsgKey:     cf.ConcatorCfgs[tag].MsgKey,
		Identifier: cf.ConcatorCfgs[tag].Identifier,
		PMsgPool:   cf.pMsgPool,
		Regexp:     cf.ConcatorCfgs[tag].Regexp,
	})
	go concator.Run(inChan)
	return inChan
}
