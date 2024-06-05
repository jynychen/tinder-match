package dto

import (
	"tinder-match/model"
)

type Person struct {
	Name        string       `json:"name"`
	Height      int          `json:"height"`
	Gender      model.Gender `json:"gender"`
	WantedDates int          `json:"wanted_dates"`
}

func (p *Person) FromModel(m *model.Person) {
	p.Name = m.Name
	p.Height = m.Height
	p.Gender = m.Gender
	p.WantedDates = m.WantedDates
}

func (p *Person) ToModel() *model.Person {
	return &model.Person{
		Name:        p.Name,
		Height:      p.Height,
		Gender:      p.Gender,
		WantedDates: p.WantedDates,
	}
}
