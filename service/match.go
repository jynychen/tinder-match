package service

import (
	"tinder-match/model"
)

type MatchService interface {
	AddSinglePersonAndMatch(*model.Person) error
	RemoveSinglePerson(name string) error
	QuerySinglePeople() ([]model.Person, error)
}

func NewMatchService() MatchService {
	return matchService{}
}

type matchService struct {
}

func (m matchService) AddSinglePersonAndMatch(person *model.Person) error {
	return nil
}

func (m matchService) RemoveSinglePerson(name string) error {
	return nil
}

func (m matchService) QuerySinglePeople() ([]model.Person, error) {
	return nil, nil
}
