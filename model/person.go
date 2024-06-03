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

func (p Person) Key() int {
	switch p.Gender {
	case GenderMale:
		return p.Height
	case GenderFemale:
		return -p.Height
	default:
		return 0
	}
}
