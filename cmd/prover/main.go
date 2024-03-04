package main

import (
	"fmt"
	"use-gnark/circuit"
	"use-gnark/utils"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	_ "github.com/consensys/gnark/std/math/bits"
)

func main() {
	cs := groth16.NewCS(ecc.BN254)
	utils.ReadFile(cs, "01-cs.bin")

	pk := groth16.NewProvingKey(ecc.BN254)
	utils.ReadFile(pk, "01-pk.bin")

	assignment := &circuit.CheckBalanceCircuit{
		Balance: 10001,
	}
	witness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}

	proof, err := groth16.Prove(cs, pk, witness)
	if err != nil {
		panic(err)
	}
	utils.WriteFile(proof, "02-proof.bin")

	publicWitness, err := witness.Public()
	if err != nil {
		panic(err)
	}
	utils.WriteFile(publicWitness, "02-public_witness.bin")

	fmt.Printf("proof sha256: %x\n", utils.FileSHA256("02-proof.bin"))
	fmt.Printf("public witness sha256: %x\n", utils.FileSHA256("02-public_witness.bin"))
}
