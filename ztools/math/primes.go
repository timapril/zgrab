package math

import "math/big"

type Prime struct {
	Number  *big.Int `json:"number,omitempty"`
	IsPrime bool     `json:"is_prime"`
	Safe    bool     `json:"safe"`
}

const CONFIDENCE int = 128

func CheckPrime(candidate *big.Int) *Prime {
	out := new(Prime)
	out.IsPrime = candidate.ProbablyPrime(CONFIDENCE)
	q := big.NewInt(0)
	q.Sub(candidate, big.NewInt(1))
	q.Div(q, big.NewInt(2))
	out.Safe = candidate.ProbablyPrime(CONFIDENCE)
	return out
}

func CheckPrimeAndAssign(candidate *big.Int) *Prime {
	out := CheckPrime(candidate)
	out.Number = candidate
	return out
}
