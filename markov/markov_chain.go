package markov

import (
	"errors"
	backpack "github.com/cdecoux/golang-backpack/pkg"
)

/*
	chain is a 2d Map that declares weights for stepping from one object to the next
 */
type markovChainStruct struct {
	chain map[interface{}]map[interface{}]int
}

func NewMarkovChain(initialStates ...interface{})  *markovChainStruct {

	chain := make(map[interface{}]map[interface{}]int)

	for _, state := range initialStates {
		chain[state] = make(map[interface{}]int)
	}

	markovChain := &markovChainStruct{
		chain: chain,
	}

	return markovChain
}

func (self *markovChainStruct) SetOrCreateWeight(src interface{}, dst interface{}, weight int) {
	self.chain[src][dst] = weight
	if weights, ok := self.chain[src]; ok {
		weights[dst] = weight
	} else {
		self.chain[src] = make(map[interface{}]int)
		self.chain[src][dst] = weight
	}
}

func (self *markovChainStruct) SetWeight(src interface{}, dst interface{}, weight int) error {
	self.chain[src][dst] = weight
	if weights, ok := self.chain[src]; ok {
		weights[dst] = weight
		return nil
	} else {
		return errors.New("src was not in existing chain")
	}
}

func (self *markovChainStruct) Step(src interface{}) interface{} {
	selector, err := backpack.NewDistributionSelector(self.chain[src])
	if err != nil {
		return nil
	}

	selection, err := selector.SelectRandom()
	if err != nil {
		return nil
	}

	return selection
}
