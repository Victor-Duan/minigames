package cpu

import (
	"math/rand"
)

type CPU struct {
	Hand []int
}

func (c *CPU) PlayCard() int {
	cardIndex := rand.Intn(len(c.Hand))
	value := c.Hand[cardIndex]

	// Remove the selected card from the players Hand
	c.Hand = append(c.Hand[:cardIndex], c.Hand[cardIndex + 1:]...)
	return value
}
