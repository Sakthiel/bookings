package helpers

import (
	"math/rand"
)

type Users struct{
	Firstname string
	Lastname string
}

func RandomGen(n int) int {

	value := rand.Intn(n)

	return  value
}