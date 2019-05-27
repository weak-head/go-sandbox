# gobox

[![Go Report Card](https://goreportcard.com/badge/github.com/weak-head/gobox?style=flat-square)](https://goreportcard.com/report/github.com/weak-head/gobox)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/weak-head/gobox)
[![GoDoc](https://godoc.org/github.com/weah-head/gobox?status.svg)](https://godoc.org/github.com/weak-head/gobox)

Simple Go Patterns and some basic examples of usage.

## Basic Patterns

- [generator](pkg/patterns/generator.go)
- [multiplexing](pkg/patterns/multiplexing.go)
- [select](pkg/patterns/selecttimeout.go)
- [daisy chain](pkg/patterns/daisychain.go)
- [quit channel](pkg/patterns/quitchannel.go)
- [restore sequence](pkg/patterns/restoreseqeunce.go)

## Tiny Examples

- [mutex](pkg/concur/counter/counter.go)
- [wait group #1](pkg/concur/chans/chans.go)
- [wait group #2](pkg/concur/chans/subred.go)
- [synced map](pkg/concur/crawler/crawler.go)
- [url poll](pkg/concur/poller/poller.go)
- [compare binary trees](pkg/concur/tree/tree.go)

## Complete Examples

- [search](pkg/example/search)
- [feed reader](pkg/example/feedreader)
