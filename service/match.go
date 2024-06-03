package service

import (
	"errors"
	"sync"
	"tinder-match/model"

	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

type MatchService interface {
	AddSinglePersonAndMatch(person *model.Person) ([]model.Person, error)
	RemoveSinglePerson(name string) error
	QuerySinglePeople(n int) ([]model.Person, error)
}

func NewMatchService() MatchService {
	return &matchService{
		singles: redblacktree.New[int, map[string]*model.Person](),
		peoples: make(map[string]int),
	}
}

var (
	ErrPersonAlreadyExists = errors.New("person already exists")
	ErrPersonNotFound      = errors.New("person not found")
	ErrPersonValidation    = errors.New("person validation failed")
)

type matchService struct {
	singles *redblacktree.Tree[int, map[string]*model.Person] // store single people by person.Key()
	peoples map[string]int                                    // map person name to person.Key()
	rw      sync.RWMutex
}

// AddSinglePersonAndMatch: Add a new user to the matching system and find any possible matches for the new user
func (m *matchService) AddSinglePersonAndMatch(person *model.Person) ([]model.Person, error) {
	if person == nil || person.WantedDates < 1 || person.Height < 1 {
		return nil, ErrPersonValidation
	}

	m.rw.Lock()
	defer m.rw.Unlock()

	if _, ok := m.peoples[person.Name]; ok {
		return nil, ErrPersonAlreadyExists
	}

	matched := m.findAndMatch(person)

	if person.WantedDates > 0 {
		m.store(person)
	}

	return matched, nil
}

// RemoveSinglePerson: Remove a user from the matching system so that the user cannot be matched anymore
func (m *matchService) RemoveSinglePerson(name string) error {
	m.rw.Lock()
	defer m.rw.Unlock()

	return m.remove(name)
}

// QuerySinglePeople: Find the most N possible matched single people, where N is a request parameter
func (m *matchService) QuerySinglePeople(n int) ([]model.Person, error) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	// init iterators
	maleIter := m.maleIterFromHighest()
	femaleIter := m.femaleIterFromShortest()

	result := make([]model.Person, 0, n+1)
	males, females := []*model.Person{}, []*model.Person{}
	for len(result) < n &&
		(maleIter != nil || femaleIter != nil || len(males) > 0 || len(females) > 0) {
		// refill possibleMale
		if maleIter != nil && len(males) == 0 {
			for _, p := range maleIter.Value() {
				males = append(males, p)
			}
			if !maleIter.Prev() || maleIter.Key() <= 0 {
				maleIter = nil
			}
		}
		// refill possibleFemale
		if femaleIter != nil && len(females) == 0 {
			for _, p := range femaleIter.Value() {
				females = append(females, p)
			}
			if !femaleIter.Prev() {
				femaleIter = nil
			}
		}

		// add one male and one female to result
		if len(males) > 0 {
			result, males = append(result, *males[0]), males[1:]
		}
		if len(females) > 0 {
			result, females = append(result, *females[0]), females[1:]
		}
	}

	if len(result) > n {
		result = result[:n]
	}
	return result, nil
}

func (m *matchService) maleIterFromHighest() *redblacktree.Iterator[int, map[string]*model.Person] {
	maleIterStart := m.singles.Right()
	if maleIterStart == nil {
		return nil
	}
	maleIter := m.singles.IteratorAt(maleIterStart)
	if maleIter.Key() <= 0 {
		return nil
	}

	return maleIter
}

func (m *matchService) femaleIterFromShortest() *redblacktree.Iterator[int, map[string]*model.Person] {
	femaleIterStart, found := m.singles.Floor(0)
	if !found {
		return nil
	}

	return m.singles.IteratorAt(femaleIterStart)
}

func (m *matchService) findAndMatch(person *model.Person) (matched []model.Person) {
	node, found := m.singles.Ceiling(-person.Key())
	if !found {
		return
	}
	iter := m.singles.IteratorAt(node)

	// iter to higher
	if iter.Key()+person.Key() == 0 && iter.Key() != 0 {
		found = iter.Next()
	}

	for person.WantedDates > 0 && found && !(person.Key() > 0 && iter.Key() > 0) {
		matchedPeople := iter.Value()
		for name, p := range matchedPeople {
			if p.WantedDates > 0 {
				p.WantedDates--
				person.WantedDates--
				if p.WantedDates == 0 {
					m.shouldRemove(name)
				}
				matched = append(matched, *p)
			}
			if person.WantedDates == 0 {
				break
			}
		}

		// date once, find next possible match
		found = iter.Next()
	}

	return
}

func (m *matchService) store(person *model.Person) {
	persons, found := m.singles.Get(person.Key())
	if !found {
		m.singles.Put(person.Key(), map[string]*model.Person{person.Name: person})
	} else {
		persons[person.Name] = person
	}

	m.peoples[person.Name] = person.Key()
}

func (m *matchService) shouldRemove(name string) {
	if err := m.remove(name); err != nil {
		return
	}
}

func (m *matchService) remove(name string) error {
	key, exists := m.peoples[name]
	if !exists {
		return ErrPersonNotFound
	}
	defer delete(m.peoples, name)

	persons, found := m.singles.Get(key)
	if found {
		delete(persons, name)
		if len(persons) == 0 {
			m.singles.Remove(key)
		}
	}

	return nil
}
