package producer

import "fmt"

type Log struct {
	Key  string
	Msg  string
	Code int
	Err  error
}

type Logger interface {
	Populate(k, m string, c int, e error)
}

type GenericError struct {
	Message string `json:"message"`
}

type GenericMessage struct {
	Message string `json:"message"`
}

func (l *Log) Populate(k, m string, c int, e error) {
	l.Key = k
	l.Msg = m
	l.Code = c
	l.Err = e
	fmt.Println(l)
}
