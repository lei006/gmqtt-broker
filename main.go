package main

import (
	"os"
	"os/signal"

	"github.com/lei006/gmqtt-broker/broker"
	"github.com/lei006/gmqtt-broker/logger"
	"go.uber.org/zap"
)

var log = logger.Get()

func main() {
	config, err := broker.ConfigureConfig(os.Args[1:])
	if err != nil {
		log.Fatal("configure broker config error", zap.Error(err))
	}

	b, err := broker.NewBroker(config)
	if err != nil {
		log.Fatal("New Broker error: ", zap.Error(err))
	}
	b.StartServer(":1883")

	s := waitForSignal()
	log.Info("signal received, broker closed.", zap.Any("signal", s))
}

func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
