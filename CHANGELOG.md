
*CURRENT*
---

- 2020-01-22 (Laisky) ci: upgrade go-utils
- 2020-01-21 (Laisky) fix: auto gc ratio bug
- 2020-01-21 (Laisky) ci: upgrade patch
- 2020-01-21 (Laisky) fix: upgrade go-utils
- 2020-01-21 (Laisky) feat: add `AutoGC`
- 2020-01-21 (Laisky) ci: upgrade patch
- 2020-01-21 (Laisky) fix: typo
- 2020-01-21 (Laisky) fix: reduce duplicate alert
- 2020-01-21 (Laisky) feat: add journal monitor
- 2020-01-21 (Laisky) feat: enable gz
- 2020-01-21 (Laisky) feat: set config file path(not dir path)
- 2020-01-16 (Laisky) docs: update changelog
- 2020-01-16 (Laisky) fix: producer sender error
- 2020-01-16 (Laisky) fix: producer sender error
- 2020-01-16 (Laisky) fix: fluentd child sender
- 2020-01-16 (Laisky) fix: producer sender typo
- 2020-01-16 (Laisky) fix: producer sender typo
- 2020-01-16 (Laisky) fix: producer monitor typo
- 2020-01-16 (Laisky) feat: improve producer

*v1.12.5*
---

- 2020-01-15 (Laisky) feat: improve performance
- 2020-01-15 (Laisky) fix
- 2020-01-14 (Laisky) fix: add `journal.gc_inteval_sec`
- 2020-01-14 (Laisky) fix: journal config typo
- 2020-01-14 (Laisky) fix: dispatcher panic
- 2020-01-14 (Laisky) docs: update readme
- 2020-01-14 (Laisky) ci: upgrade go 1.13.6
- 2020-01-14 (Laisky) docs: update readme
- 2020-01-14 (Laisky) docs: update settings
- 2020-01-14 (Laisky) fix: some logs loss

*v1.12.3*
---

- 2020-01-10 (Laisky) ci: upgrade patch
- 2020-01-10 (Laisky) docs: update readme
- 2020-01-10 (Laisky) docs: add more comment
- 2020-01-10 (Laisky) docs: typo
- 2020-01-10 (Laisky) style: change some vars name
- 2020-01-07 (Laisky) fix: detect log-alert config
- 2020-01-07 (Laisky) feat: enable & telegram alert
- 2020-01-07 (Laisky) fix: alert pusher
- 2020-01-06 (Laisky) feat: add alert wechat service
- 2020-01-03 (Laisky) fix: ignore es type conflict
- 2020-01-03 (Laisky) fix: handle es return error

*v1.12.1*
---

- 2019-12-27 (Laisky) fix: support es6
- 2019-12-12 (Laisky) ci: upgrade go-utils
- 2019-12-12 (Laisky) ci: upgrade go-utils

*v1.12.0*
---

- 2019-12-03 (Laisky) ci: remove marathon build
- 2019-12-03 (Laisky) docs: update docs
- 2019-12-03 (Laisky) fix: add more log
- 2019-12-02 (Laisky) feat(paas-445): support k8s fluent-bit daemonset
- 2019-11-29 (Laisky) ci: upgrade golang v1.13.4
- 2019-11-28 (Laisky) feat(paas-441): support k8s
- 2019-11-28 (Laisky) fix: rewrite rsys log
- 2019-11-28 (Laisky) fix: do not reset rsyslog tag
- 2019-11-28 (Laisky) fix: do not reset rsyslog tag
- 2019-11-28 (Laisky) feat(paas-444): rsyslog support custom tag
- 2019-11-12 (Laisky) fix: need root
- 2019-11-12 (Laisky) ci: set user
- 2019-11-12 (Laisky) fix: avoid empty es body
- 2019-11-11 (Laisky) feat: upgrade zap
- 2019-11-08 (Laisky) fix: zap fields parse
- 2019-11-08 (Laisky) fix: upgrade zap
- 2019-11-06 (Laisky) feat: add logger alert
- 2019-10-23 (Laisky) perf: disable gin log
- 2019-10-18 (Laisky) ci: upgrade golang v1.13.3
- 2019-10-18 (Laisky) fix: upgrade go-utils v1.8.1

*v1.11.5*
---

- 2019-10-16 (Laisky) docs: update changelog
- 2019-10-16 (Laisky) perf: tidy
- 2019-10-15 (Laisky) test: fix ci lint
- 2019-10-09 (Laisky) docs: update example settings
- 2019-10-09 (Laisky) fix: fluentd decodeMsg with context

*v1.11.4*
---

- 2019-10-08 (Laisky) fix(paas-420): - journal og ttl - duplicate id - kafka
- 2019-09-30 (Laisky) fix: upgrade go-utils v1.7.8
- 2019-09-30 (Laisky) fix: rotate counter start at 1
- 2019-09-30 (Laisky) fix: fix duplicate in counter
- 2019-09-29 (Laisky) fix(paas-408): keep one legacy buf file to descend duplicate'
- 2019-09-29 (Laisky) fix: use normal counter to fix duplicate
- 2019-09-23 (Laisky) build: upgrade go-utils v1.7.7
- 2019-09-18 (Laisky) build: upgrade go-utils
- 2019-09-16 (Laisky) build: go mod
- 2019-09-16 (Laisky) feat(paas-412): support k8s log
- 2019-09-12 (Laisky) feat: preallocate file size in journal

*v1.11.1*
---

- 2019-09-06 (Laisky) docs: update changelog
- 2019-09-06 (Laisky) ci: upgrade go-utils v1.7.6
- 2019-09-05 (Laisky) fix: refact contronllor's ctx
- 2019-09-05 (Laisky) fix: use stopChan replace cancel
- 2019-09-05 (Laisky) fix: kafka recv after closed

*v1.11.0*
---

- 2019-09-05 (Laisky) perf: reduce goroutines
- 2019-09-05 (Laisky) docs: update changelog
- 2019-09-05 (Laisky) ci: upgrade go-utils v1.7.5
- 2019-09-05 (Laisky) fix: shrink deps
- 2019-09-05 (Laisky) ci: upgrade go-utils v1.7.4
- 2019-09-04 (Laisky) fix: upgrade go-utils
- 2019-09-04 (Laisky) fix: add context in postpipeline
- 2019-09-04 (Laisky) build: upgrade go v1.13.0
- 2019-09-04 (Laisky) fix: add context in accptor
- 2019-09-04 (Laisky) fix: add context in acceptorPipeline
- 2019-09-04 (Laisky) fix: add context in producer
- 2019-09-04 (Laisky) fix: refactor tagPipeline
- 2019-09-03 (Laisky) ci: upgrade go-utils v1.7.3
- 2019-09-03 (Laisky) docs: update changelog

*v1.10.8*
---

- 2019-09-03 (Laisky) fix: ctx.Done return

*v1.10.7*
---

- 2019-09-02 (Laisky) ci: upgrade go-utils v1.7.0
- 2019-09-02 (Laisky) feat: add context to control journal
- 2019-09-02 (Laisky) feat(paas-405): upgrade go-utils, use new ids set
- 2019-08-29 (Laisky) docs: update example settings

*v1.10.6*
---

- 2019-08-29 (Laisky) fix: blocking when commitChan is full

*v1.10.5*
---

- 2019-08-28 (Laisky) docs: update changelog
- 2019-08-28 (Laisky) fix: commit chan should not discard
- 2019-08-28 (Laisky) docs: update changelog
- 2019-08-28 (Laisky) perf: periodic gc

*v1.10.4*
---

- 2019-08-28 (Laisky) perf: disable jj gc

*v1.10.3*
---

- 2019-08-28 (Laisky) docs: update changelog
- 2019-08-28 (Laisky) ci: upgrade go-utils
- 2019-08-27 (Laisky) fix(paas-403): journal roll bug
- 2019-08-27 (Laisky) build: upgrade go-utils
- 2019-08-27 (Laisky) build: `http_proxy` should in lowercase
- 2019-08-27 (Laisky) fix: upgrade to go v1.12.9
- 2019-08-27 (Laisky) fix: upgrade to go v1.12.9
- 2019-08-27 (Laisky) perf: use `NewMonotonicCounterFromN`
- 2019-08-26 (Laisky) feat(paas-398): split journal into different directory by tag
- 2019-08-23 (Laisky) fix(paas-397): msg disorder after acceptorPipeline, then cause concator error

*v1.10.2*
---

- 2019-08-21 (Laisky) build: upgrade go-utils to v1.6.2
- 2019-08-21 (Laisky) fix: format warn
- 2019-08-21 (Laisky) fix(paas-397): missing content

*v1.10.1*
---

- 2019-08-20 (Laisky) fix: ignore es conflict error

*v1.10.0*
---

- 2019-08-20 (Laisky) ci: upgrade go-utils
- 2019-08-20 (Laisky) ci: upgrade go-utils
- 2019-08-19 (Laisky) docs: update quickstart
- 2019-08-19 (Laisky) feat: enable gz in journal
- 2019-08-19 (Laisky) style: more log

*v1.9.3*
---

- 2019-08-15 (Laisky) build: upgrade go-utils to v1.5.4
- 2019-08-15 (Laisky) fix: double default postfilter
- 2019-08-14 (Laisky) docs: update changelog
- 2019-08-14 (Laisky) feat: support `@RANDOM_STRING`
- 2019-08-14 (Laisky) build: upgrade to go 1.12.7
- 2019-08-14 (Laisky) feat(paas-390): add wuling mapping
- 2019-07-23 (Laisky) perf: upgrade go-utils to v1.5.3
- 2019-07-17 (Laisky) fix: optimize regexp
- 2019-07-17 (Laisky) docs: update example settings
- 2019-06-26 (Laisky) build: upgrade go-utils to v1.5.1
- 2019-06-26 (Laisky) fix: go mod conflict
- 2019-06-21 (Laisky) build: upgrade to golang:1.12.6
- 2019-06-12 (Laisky) ci: disable vendor cache
- 2019-06-12 (Laisky) test: fix test docket

*v1.9.1*
---

- 2019-06-12 (Laisky) test: fix test docket
- 2019-06-12 (Laisky) test: fix test docket
- 2019-06-12 (Laisky) ci: update dockerfile
- 2019-06-12 (Laisky) ci: update dockerfile
- 2019-06-12 (Laisky) docs: update readme
- 2019-06-12 (Laisky) build: fix package conflict
- 2019-06-12 (Laisky) ci: update travis
- 2019-06-12 (Laisky) ci: update cache
- 2019-06-10 (Laisky) ci: fix dependencies error
- 2019-06-10 (Laisky) fix: add uint32set & improve rotate
- 2019-06-04 (Laisky) docs: update readme

*v1.9.0*
---

- 2019-06-04 (Laisky) style: add some comment
- 2019-06-04 (Laisky) ci: add prometheus
- 2019-06-04 (Laisky) ci: upgrade go-utils v1.4.0
- 2019-06-04 (Laisky) fix(paas-361): replace iris by gin
- 2019-05-31 (Laisky) fix(paas-360): corrupt file panic
- 2019-05-31 (Laisky) ci: upgrade go-utils v1.3.8
- 2019-05-31 (Laisky) fix(paas-357): reduce memory use
- 2019-05-29 (Laisky) ci: upgrade go-utils v1.3.6
- 2019-05-29 (Laisky) docs: add comment

*v1.8.9*
---

- 2019-05-29 (Laisky) ci: upgrade go-utils
- 2019-05-29 (Laisky) fix: compatable with smaller rotate id
- 2019-05-28 (Laisky) docs: update readme
- 2019-05-28 (Laisky) fix: recvs check active env
- 2019-05-28 (Laisky) feat: add null sender
- 2019-05-28 (Laisky) +null
- 2019-05-27 (Laisky) ci: upgrade go-utils
- 2019-05-27 (Laisky) docs: update readme

*v1.8.8*
---

- 2019-05-24 (Laisky) fix(paas-357): journal memory leak

*v1.8.7*
---

- 2019-05-23 (Laisky) ci: update gomod
- 2019-05-23 (Laisky) docs: update readme

*v1.8.6*
---

- 2019-05-23 (Laisky) perf: reduce alloc in es sender
- 2019-05-22 (Laisky) ci: upgrade golang to 1.12.5
- 2019-05-22 (Laisky) perf: only allow one rotate waiting
- 2019-05-22 (Laisky) perf: only load all ids once
- 2019-05-22 (Laisky) fix: upgrade go-utils to v1.3.1

*1.8.5*
---

- 2019-05-15 (Laisky) docs: update readme
- 2019-05-14 (Laisky) ci: improve docker cache
- 2019-05-08 (Laisky) feat(paas-344): add prometheus metrics at `/metrics`

*1.8.4*
---

- 2019-04-26 (Laisky) build: no neee do `go mod download`
- 2019-04-26 (Laisky) perf: improve `runLB`

*1.8.3*
---

- 2019-04-17 (Laisky) build: fix bin file path
- 2019-04-17 (Laisky) fix: change ci to gomod
- 2019-04-17 (Laisky) build: replace glide by gomod in dockerfile
- 2019-04-17 (Laisky) build: replace glide by gomod
- 2019-04-16 (Laisky) docs: fix dockerfiles

*1.8.2*
---

- 2019-04-12 (Laisky) fix: ts format only support `.`
- 2019-04-12 (Laisky) fix: dispatch lock corner bug
- 2019-04-12 (Laisky) docs: add dockerfile relations

*1.8.1*
---

- 2019-04-09 (Laisky) fix(paas-320): dispatcher conflict with inChanForEachTag

*1.8*
---

- 2019-04-04 (Laisky) feat(paas-320): let dispatcher & tagpipeline parallel
- 2019-03-05 (Laisky) style: format improve

*1.7.2*
---

- 2019-03-04 (Laisky) docs: update changelog
- 2019-03-04 (Laisky) feat(paas-312): postfilter support plugins
- 2019-03-01 (Laisky) docs: add push at docker doc
- 2019-03-01 (Laisky) build: upgrade to golang:1.12
- 2019-03-01 (Laisky) build: upgrade to golang:1.12
- 2019-03-01 (Laisky) build: upgrade go-utils to v9.1
- 2019-03-01 (Laisky) fix(paas-312): add some tests and fix some bugs
- 2019-02-28 (Laisky) docs: update changelog

*1.7.1*
---

- 2019-02-28 (Laisky) docs: fix settings demo
- 2019-02-28 (Laisky) docs: add quick start
- 2019-02-21 (Laisky) fix: more details in log
- 2019-02-21 (Laisky) feat(paas-288): add fluentd-forward log
- 2019-02-20 (Laisky) ci: add id into docker tag
- 2019-02-20 (Laisky) fix
- 2019-02-20 (Laisky) fix
- 2019-02-19 (Laisky) test: fix test readme
- 2019-02-15 (Laisky) ci: fix test
- 2019-02-15 (Laisky) ci: fix test
- 2019-02-15 (Laisky) ci: fix test
- 2019-02-15 (Laisky) ci: add test
- 2019-02-15 (Laisky) docs: update changelog

*1.7*
---

- 2019-02-15 (Laisky) build: upgrade to alpine3.9
- 2019-02-14 (Laisky) perf: reduce memory usage in recv
- 2019-02-14 (Laisky) ci: add ci retry
- 2019-02-14 (Laisky) perf: replace all encoding/json
- 2019-02-13 (Laisky) feat(paas-294): replace codec by msgp
- 2019-02-12 (Laisky) perf: replace builtin json by `github.com/json-iterator/go` in recvs
- 2019-02-12 (Laisky) perf: replace builtin json by `github.com/json-iterator/go`
- 2019-02-12 (Laisky) style: fix some format

*1.6.8*
---

- 2019-02-13 (Laisky) fix(paas-294): compatable to 0.9 journal data file

*1.6.7*
---

- 2019-02-11 (Laisky) build: upgrade go-utils

*1.6.6*
---

- 2019-02-01 (Laisky) fix: RegexNamedSubMatch should remove empty key
- 2019-02-01 (Laisky) docs: update readme
- 2019-02-01 (Laisky) ci: add ci var
- 2019-01-31 (Laisky) fix: remove empty key and replace '.' in key
- 2019-01-29 (Laisky) test: add more test case

*1.6.5*
---

- 2019-01-30 (Laisky) fix(paas-287): should set payload to nil before decode fluentd's msg
- 2019-01-30 (Laisky) fix: add more logs

*1.6.4*
---

- 2019-01-29 (Laisky) fix: add `max_allowed_ahead_sec`

*1.6.3*
---

- 2019-01-29 (Laisky) fix: fluentd producer tags should not append env

*1.6.2*
---

- 2019-01-25 (Laisky) docs: update CHANGELOG
- 2019-01-25 (Laisky) perf: upgrade zap
- 2019-01-25 (Laisky) build: upgrade go-utils
- 2019-01-25 (Laisky) perf: use uniform clock in go-utils
- 2019-01-24 (Laisky) fix(paas-284): reduce debug log to reduce cpu
- 2019-01-24 (Laisky) fix: remove useless pprof entrypoint
- 2019-01-24 (Laisky) fix: check timer's ts
- 2019-01-24 (Laisky) build: upgrade iris
- 2019-01-24 (Laisky) perf: reduce invoke time.Now
- 2019-01-24 (Laisky) build: upgrade golang v1.11.5
- 2019-01-24 (Laisky) docs: update docs
- 2019-01-23 (Laisky) fix(paas-282): some bugs during testing
- 2019-01-23 (Laisky) fix: fluentd sender missing env
- 2019-01-17 (Laisky) docs: update docs

*1.6.1*
---

- 2019-01-17 (Laisky) fix: `append_time_zone`
- 2019-01-17 (Laisky) fix: kafka recv error
- 2019-01-17 (Laisky) docs: add docker readme
- 2019-01-17 (Laisky) docs: add pipeline label
- 2019-01-17 (Laisky) ci: `IMAGE_NAME` in gitlab-ci
- 2019-01-17 (Laisky) ci: fix image bug
- 2019-01-16 (Laisky) ci: fix image bug
- 2019-01-16 (Laisky) ci: add proxy
- 2019-01-16 (Laisky) ci: add golang-stretch
- 2019-01-16 (Laisky) ci: enable marathon
- 2019-01-16 (Laisky) ci: add strecth-mfs
- 2019-01-16 (Laisky) ci: fix ci yml
- 2019-01-16 (Laisky) fix(paas-281): add gitlab-ci
- 2019-01-15 (Laisky) docs: update TOC
- 2019-01-15 (Laisky) docs: add cn doc
- 2019-01-15 (Laisky) style: rename dockerfile
- 2019-01-15 (Laisky) feat(paas-276): able to load configuration from config-server
- 2019-01-14 (Laisky) feat(paas-275): refactory controllor - parse settings more flexible

*1.6*
---

- 2019-01-11 (Laisky) build: upgrade go-utils
- 2019-01-10 (Laisky) fix: flatten delimiter
- 2019-01-10 (Laisky) feat: flatten httprecv json
- 2019-01-10 (Laisky) fix: no not panic when got incorrect fluentd msg
- 2019-01-10 (Laisky) fix(paas-272): do not retry all es msgs when part of mesaages got error
- 2019-01-10 (Laisky) fix: add more err log
- 2019-01-10 (Laisky) style: less hardcoding
- 2019-01-10 (Laisky) feat(paas-270): add signature in httprecv
- 2019-01-09 (Laisky) feat(paas-259): support external log api
- 2019-01-09 (Laisky) build: upgrade glide & marathon
- 2019-01-07 (Laisky) feat(paas-261): add http json recv
- 2019-01-09 (Laisky) feat(paas-265): use stretch to mount mfs

*1.5.4*
---

- 2019-01-07 (Laisky) fix(paas-258): backup logs missing time
- 2019-01-02 (Laisky) docs: update changelog

*1.5.3*
---

- 2018-12-27 (Laisky) feat(paas-249): set default_field `message` default to ""
- 2018-12-26 (Laisky) fix: add panic logging in goroutines
- 2018-12-26 (Laisky) docs: update changelog

*1.5.2*
---

- 2018-12-25 (Laisky) fix(paas-245): esdispatcher error
- 2018-12-25 (Laisky) fix(paas-246): journal data should rewrite into journal
- 2018-12-25 (Laisky) perf(paas-245): remove kafka buffer

*1.5.1*
---

- 2018-12-24 (Laisky) build: add http proxy
- 2018-12-24 (Laisky) fix(paas-244): upgrade go-utils

*1.5*
---

- 2018-12-24 (Laisky) feat(paas-212): async dispatcher
- 2018-12-21 (Laisky) fix(paas-212): some errors during testing
- 2018-12-19 (Laisky) Revert "Revert "Merge branch 'feature/paas-212-fluent-buf' into develop""
- 2018-12-19 (Laisky) Revert "Merge branch 'feature/paas-212-fluent-buf' into develop"
- 2018-12-18 (Laisky) feat(paas-212): replace fluentd-buf completely

*1.4.2*
---

- 2018-12-12 (Laisky) fix: upgrade go-syslog to fix bug
- 2018-12-12 (Laisky) feat(paas-225): support BLB health check
- 2018-12-11 (Laisky) style: clean code

*1.4.1*
---

- 2018-12-12 (Laisky) fix(paas-231): trim blank in fields

*1.4*
---

- 2018-12-11 (Laisky) fix: do not send geely backup in go-fluentd now
- 2018-12-06 (Laisky) fix: do not retry when sender is too busy
- 2018-12-06 (Laisky) fix(paas-222): race in dispatcher
- 2018-12-06 (Laisky) fix: kafka_recvs config
- 2018-12-06 (Laisky) fix: typo
- 2018-12-06 (Laisky) fix: typo
- 2018-12-06 (Laisky) feat(paas-220): support rsyslog recv
- 2018-12-04 (Laisky) fix(paas-208): change flatten delimiter to "__"
- 2018-12-04 (Laisky) build: rename to .com`
