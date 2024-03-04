package main

import (
	"use-gnark/utils"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
)

func main() {
	vk := groth16.NewVerifyingKey(ecc.BN254)
	utils.ReadFile(vk, "02-vk.bin")

	proof := groth16.NewProof(ecc.BN254)
	utils.ReadFile(proof, "02-proof.bin")

	publicWitness, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}
	utils.ReadFile(publicWitness, "02-public_witness.bin")

	if err := groth16.Verify(proof, vk, publicWitness); err != nil {
		panic(err)
	}
}
