package main

import (
	"fmt"
	"strings"
)

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (nl *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range nl.Neurons {
		result = append(result, &nl.Neurons[i])
	}
	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{Neurons: make([]Neuron, count)}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

// VisualizeNetwork creates a better visual representation of the neural network.
func VisualizeNetwork(layers ...NeuronInterface) {
	fmt.Println("Neural Network Visualization:")

	// Display neurons layer by layer
	layerStrs := []string{}
	for _, layer := range layers {
		layerStr := ""
		for range layer.Iter() {
			layerStr += "(O) " // each neuron represented as (O)
		}
		layerStrs = append(layerStrs, strings.TrimSpace(layerStr))
	}

	// Draw the layers
	for i := 0; i < len(layerStrs); i++ {
		fmt.Println(layerStrs[i])

		// Draw connections between layers if there’s a next layer
		if i < len(layerStrs)-1 {
			connStr := ""
			// Draw lines to simulate connections from each neuron in the current layer
			for j := 0; j < len(layerStrs[i]); j += 4 {
				connStr += "|   "
			}
			fmt.Println(connStr)

			// Draw diagonal lines between neurons in layers
			diagStr := ""
			for j := 0; j < len(layerStrs[i]); j += 4 {
				diagStr += "/ \\ "
			}
			fmt.Println(diagStr)
		}
	}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1 := NewNeuronLayer(3)
	layer2 := NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)

	// Visualize the network
	VisualizeNetwork(neuron1, layer1, layer2, neuron2)
}
