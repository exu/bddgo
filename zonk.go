package bddgo

import (
	"math/rand"
)

type Zonk struct {
	Doors   [3]int
	Opened  [3]int
	Choice  bool
	Choices [3]int
	Records [3]int
}

func (this *Zonk) Init() {
	this.Doors = [3]int{0, 0, 0}
	this.Doors[rand.Intn(2)] = 1
}
