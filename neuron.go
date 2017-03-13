// A neuron has a difference in voltage between inside and outside.
//
// This difference is called membrane potential.
//
// Each neuron has a specific negative membrane potential at rest which is
// the same for each neuron and is called resting potential.
//
// Neurons have connections to each other.
//
// A neuron receives signals via its dendrites.
//
// A neuron sends signals via its axon.
//
// A signal sent from one neuron's axon to another neuron's dendrite can activate or inhibit the receiving neuron.
//
// An excitatory neuron causes a receiving neuron's membrane potential to become more positive.
//
// An inhibitory neuron causes a receiving neuron's membrane potential to become more negative.
//
// For a neuron to be activated, its potential needs to reach a threshold of -55mV. This is threshold potential.
//
// If the threshold potential is reached, a neuron produces an action potential.
//
// An action potential travels along a neuron's axon and can cause the activation or inhibition of other neurons via their dendrites.
//
// Scenario: Let's have three neurons: an excitatory neuron, an interneuron, an a receiving neuron which is currently at rest.
//
// We want to know what happens when a neuron is excited, inhibited or both.
//
// Let's assume that neuron 1 is already active and it sends a signal to neuron 2.
//
// We want to calculate when an action potential occurs.

// n1 := neuron{..}
// n2 := neuron{..}
// n1.setAxon(n2)
// n2.addDendrite(n1)
// n2.ChangePotential(5.0) // this could cause n1 to fire, which can activate n2
package main

import (
	"fmt"
)

type (
	// A potential is a difference in voltage between the inside and the outside of a cell.
	potential int
	// A neuron is a cell that receives and sends signals.
	neuron struct{
		// axon
		// dendrite
		potential potential
	}

	// threshold := potential(3)
	// resting := potential(1)
	receptorPotential  struct {
		amplitude int
		duration  int64 // Ask Lars: there is an actual type duration. How to use?
	}
	actionPotential struct {
		frequency int
		number    int
	}
	neurotransmitter struct {
		kind  string // E.g., dopamine, serotonin.
		amount int
	}
	// neurotransmitter{kind: "serotonin", amount: 100}
)

func (r potential) reachThreshold() string {
	if r <= -5 {
		return "Threshold for action potential is not reached."
	}
	return "Threshold for action potential is reached."
}

func main() {
	r := potential(-69)
	fmt.Println(r.reachThreshold())

}

// func (a receptorPotential) increaseRestingPotential() string {

