package model

type Gender string

const (
	GenderMale   = Gender("male")
	GenderFemale = Gender("female")
)

type Person struct {
	Name        string
	Height      int
	Gender      Gender
	WantedDates int
}
