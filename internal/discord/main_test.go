package discord

import "math/rand"

//go:generate mockgen -source=contract.go -destination=mocks/mock.go

func init() {
	rand.Seed(0)
}
