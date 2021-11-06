package main

import (
	"github.com/cdecoux/golang-markov-chain/markov"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	/*
	   Set the Log Level of LogRUS
	*/
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "INFO"
	}


	logLevel, err := log.ParseLevel(lvl)
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)


	mkvChain := markov.NewMarkovChain([]markov.State{"a", nil, "c", "b", 2, 3}...)

	// a
	_ = mkvChain.SetWeight("a", "b", 50)
	_ = mkvChain.SetWeight("a", 1, 1)
	// b
	_ = mkvChain.SetWeight("b", "c", 100)
	_ = mkvChain.SetWeight("b", "b", 5)
	_ = mkvChain.SetWeight("b", 2, 1)
	// c
	_ = mkvChain.SetWeight(nil, "c", 100)
	_ = mkvChain.SetWeight("c", "c", 25)
	_ = mkvChain.SetWeight("c", "a", 15)
	_ = mkvChain.SetWeight("c", 3, 5)


	// Indefinitely loop through chain
	currentState := markov.State("a")
	counter := make(map[interface{}]int)
	ticker := time.NewTicker(time.Second).C

	log.Info(currentState)
	for i := 0; i < 10000; i++ {
		select {
		case <- ticker:
			currentState = mkvChain.Step(currentState)
			counter[currentState]++
			log.Info(currentState)
		}
	}

}
