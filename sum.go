package main

import (
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// Circuit defines a sum circuit
type Circuit struct {
	A, B, C frontend.Variable
}

// declares the circuit's constraints
func (circuit *Circuit) Define(api frontend.API) error {
	// Ensure that a + b == c
	sum := api.Add(circuit.A, circuit.B)
	api.AssertIsEqual(sum, circuit.C)
	return nil
}

func main() {
	// Define the circuit
	var circuit Circuit

	// Compile the circuit into a R1CS
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		fmt.Println("Error compiling circuit:", err)
		return
	}

	// Save the R1CS to a file
	file, err := os.Create("circuit.r1cs")
	if err != nil {
		fmt.Println("Error creating R1CS file:", err)
		return
	}
	defer file.Close()

	// Use WriteTo method to serialize the R1CS to the file
	_, err = ccs.WriteTo(file)
	if err != nil {
		fmt.Println("Error writing R1CS file:", err)
	}
}
