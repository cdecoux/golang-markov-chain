package markov

import (
	"errors"
	backpack "github.com/cdecoux/golang-backpack/pkg"
	log "github.com/sirupsen/logrus"
)

type State interface {}

/*
	chain is a 2d Map that declares weights for stepping from one object to the next
 */
type markovChainStruct struct {
	chain map[interface{}]map[interface{}]int
}

func NewMarkovChain(initialStates ...State)  *markovChainStruct {

	chain := make(map[interface{}]map[interface{}]int)

	for _, src := range initialStates {
		chain[src] = make(map[interface{}]int)

		for _, dst := range initialStates {
			chain[src][dst] = 0
		}
	}

	log.Debug(chain)

	markovChain := &markovChainStruct{
		chain: chain,
	}

	return markovChain
}

func (self *markovChainStruct) AddStates(states ...State)  {
	for _, state := range states {
		// Check if state already exists, so we don't remake slice
		if _, exists := self.chain[state]; !exists {
			self.chain[state] = make(map[interface{}]int)
		}
	}
}

func (self *markovChainStruct) SetOrCreateWeight(src State, dst State, weight int) {
	self.chain[src][dst] = weight
	if weights, ok := self.chain[src]; ok {
		weights[dst] = weight
	} else {
		self.chain[src] = make(map[interface{}]int)
		self.chain[src][dst] = weight
	}
}

func (self *markovChainStruct) SetWeight(src State, dst State, weight int) error {
	// Check if src/dst are in chain.
	// If they don't exists, then throw errors
	if weights, ok := self.chain[src]; ok {
		if _, ok := weights[dst]; ok {
			weights[dst] = weight
		}

		return errors.New("dst was not in existing chain")
	}

	return errors.New("src was not in existing chain")
}

func (self *markovChainStruct) step(src State, selector backpack.DistributionSelector) State {

	var selection State
	var err error
	if selector != nil {
		selection, err = selector.SelectRandom()
		if err != nil {
			log.Error(err)
			return nil
		}
	}
	return selection
}

func (self *markovChainStruct) Step(src State) State {
	selector, err := backpack.NewDistributionSelector(self.chain[src])
	if err != nil {
		log.Error(err)
		return nil
	}

	return self.step(src, selector)
}

/*
	Does N steps and returns an ordered slice of steps.
	Should the markov chain not be cyclic, <nil> will be used for the terminator.
	This function will immediately return the results once <nil> is hit regardless of 'n'
 */
func (self *markovChainStruct) StepN(src State, n int) []State {

	// Create State slice for returning
	results := make([]State, 0, n)

	// Create cache for distribution selectors
	selectorCache := make(map[State]backpack.DistributionSelector)
	for state, distributionMap := range self.chain {
		selector, _ := backpack.NewDistributionSelector(distributionMap)
		selectorCache[state] = selector
	}

	currentState := src

	for i := 0; i < n; i++ {
		if currentState == nil {
			return results
		}
		state := self.step(currentState, selectorCache[currentState])
		results = append(results, state)
		currentState = state
	}

	return results
}

