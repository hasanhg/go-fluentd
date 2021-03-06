package libs

//go:generate msgp

// FluentMsg is the structure of fluent message
type FluentMsg struct {
	Tag     string
	Message map[string]interface{}
	Id      int64
	ExtIds  []int64
}

type FluentBatchMsg []interface{}
