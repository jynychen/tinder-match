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
