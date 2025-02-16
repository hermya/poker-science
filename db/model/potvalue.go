package model

type PotValue struct {
	PotPriority int8    // Once one or more people match the all-in value in a round, PotPriority identifies side pot
	PotTotal    float64 // Value of this pot
}
