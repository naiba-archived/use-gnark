package main

import (
	"fmt"
	"use-gnark/circuit"
	"use-gnark/utils"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func main() {
	var circuit circuit.CheckBalanceCircuit
	cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}

	pk, vk, err := groth16.Setup(cs)
	if err != nil {
		return
	}

	utils.WriteFile(cs, "01-cs.bin")
	utils.WriteFile(pk, "01-pk.bin")
	utils.WriteFile(vk, "02-vk.bin")

	fmt.Printf("cs sha256: %x\n", utils.FileSHA256("01-cs.bin"))
	fmt.Printf("pk sha256: %x\n", utils.FileSHA256("01-pk.bin"))
	fmt.Printf("vk sha256: %x\n", utils.FileSHA256("02-vk.bin"))
}
