package main

import (
	"github.com/cdecoux/golang-markvon-chain/markov"
	log "github.com/sirupsen/logrus"
	"os"
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


	mkvChain := markov.NewMarkovChain([]markov.State{"a", "b", "c", 1, 2, 3}...)

	// a
	_ = mkvChain.SetWeight("a", "b", 50)
	_ = mkvChain.SetWeight("a", "a", 60)
	_ = mkvChain.SetWeight("a", 1, 1)
	// b
	_ = mkvChain.SetWeight("b", "c", 100)
	_ = mkvChain.SetWeight("b", "b", 5)
	_ = mkvChain.SetWeight("b", 2, 1)
	// c
	_ = mkvChain.SetWeight("c", "c", 100)
	_ = mkvChain.SetWeight("c", "a", 25)
	_ = mkvChain.SetWeight("c", 3, 1)


	steps := mkvChain.StepN("a", 100)

	log.Info(steps)

}
