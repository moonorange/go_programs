package main

import (
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) Handle(msg kafka.Message) {
	var m Msg
	json.Unmarshal([]byte(msg.Value), &m)
	logrus.Infof("Handled msg: %+v", msg)
}
