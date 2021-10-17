package main

import (
	"github.com/cdecoux/golang-markvon-chain/markov"
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


	mkvChain := markov.NewMarkovChain("a", "b", "c")

	_ = mkvChain.SetWeight("a", "b", 1)
	_ = mkvChain.SetWeight("a", "a", 5)
	_ = mkvChain.SetWeight("b", "c", 1)
	_ = mkvChain.SetWeight("c", "c", 10)
	_ = mkvChain.SetWeight("c", "a", 1)


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
