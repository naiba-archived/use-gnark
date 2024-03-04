package circuit

import "github.com/consensys/gnark/frontend"

type CheckBalanceCircuit struct {
	Balance frontend.Variable `gnark:"balance"`
}

func (circuit *CheckBalanceCircuit) Define(api frontend.API) error {
	ret := api.Cmp(10000, circuit.Balance)
	api.AssertIsEqual(ret, -1)
	return nil
}
