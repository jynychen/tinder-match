package service

import (
	"testing"
	"tinder-match/model"
)

func TestAddSinglePersonAndMatch(t *testing.T) {
	ms := NewMatchService()

	// validation
	result, err := ms.AddSinglePersonAndMatch(nil)
	if err != ErrPersonValidation {
		t.Errorf("Expected validation error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
	result, err = ms.AddSinglePersonAndMatch(&model.Person{Name: "John", Height: 0, Gender: model.GenderMale, WantedDates: 3})
	if err != ErrPersonValidation {
		t.Errorf("Expected validation error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
	result, err = ms.AddSinglePersonAndMatch(&model.Person{Name: "John", Height: 2, Gender: model.GenderMale, WantedDates: 0})
	if err != ErrPersonValidation {
		t.Errorf("Expected validation error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	person := &model.Person{Name: "John", Height: 180, Gender: model.GenderMale, WantedDates: 3}
	result, err = ms.AddSinglePersonAndMatch(person)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	result, err = ms.AddSinglePersonAndMatch(person)
	if err != ErrPersonAlreadyExists {
		t.Errorf("Expected duplicate person error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	// match
	female := &model.Person{Name: "Jane", Height: 170, Gender: model.GenderFemale, WantedDates: 2}
	result, err = ms.AddSinglePersonAndMatch(female)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %v", result)
	}
	if result[0].Name != "John" {
		t.Errorf("Expected John, got %v", result[0].Name)
	}
	if person.WantedDates != 2 || female.WantedDates != 1 {
		t.Errorf("Matching failed: John's or Jane's WantedDates incorrect")
	}

	male := &model.Person{Name: "Mike", Height: 175, Gender: model.GenderMale, WantedDates: 1}
	result, err = ms.AddSinglePersonAndMatch(male)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %v", result)
	}
	if result[0].Name != "Jane" {
		t.Errorf("Expected Jane, got %v", result[0].Name)
	}
	if female.WantedDates != 0 || male.WantedDates != 0 {
		t.Errorf("Matching failed: Jane's or Mike's WantedDates incorrect")
	}

	// no match
	newPerson := &model.Person{Name: "Bob", Height: 160, Gender: model.GenderMale, WantedDates: 1}
	result, err = ms.AddSinglePersonAndMatch(newPerson)
	if err != nil {
		t.Errorf("Unexpected error when adding unmatched person: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
	if newPerson.WantedDates != 1 {
		t.Errorf("Expected WantedDates 1, got %v", newPerson.WantedDates)
	}
}

func TestAddSinglePersonAndMatch1(t *testing.T) {
	ms := NewMatchService()

	user1 := &model.Person{Name: "user1", Height: 180, Gender: model.GenderMale, WantedDates: 3}
	result, err := ms.AddSinglePersonAndMatch(user1)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	user2 := &model.Person{Name: "user2", Height: 170, Gender: model.GenderMale, WantedDates: 3}
	result, err = ms.AddSinglePersonAndMatch(user2)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	user3 := &model.Person{Name: "user3", Height: 175, Gender: model.GenderMale, WantedDates: 3}
	result, err = ms.AddSinglePersonAndMatch(user3)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	user4 := &model.Person{Name: "user4", Height: 173, Gender: model.GenderFemale, WantedDates: 3}
	result, err = ms.AddSinglePersonAndMatch(user4)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 1 result, got %v", result)
	}
	if result[0].Name != "user3" {
		t.Errorf("Expected user3, got %v", result[0].Name)
	}
	if result[1].Name != "user1" {
		t.Errorf("Expected user1, got %v", result[1].Name)
	}

	user5 := &model.Person{Name: "user5", Height: 175, Gender: model.GenderFemale, WantedDates: 3}
	result, err = ms.AddSinglePersonAndMatch(user5)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %v", result)
	}
	if result[0].Name != "user1" {
		t.Errorf("Expected user1, got %v", result[0].Name)
	}
}

func TestRemoveSinglePerson(t *testing.T) {
	ms := NewMatchService()

	err := ms.RemoveSinglePerson("John")
	if err != ErrPersonNotFound {
		t.Errorf("Expected person not found error, got %v", err)
	}

	person := &model.Person{Name: "John", Height: 180, Gender: model.GenderMale, WantedDates: 3}
	_, err = ms.AddSinglePersonAndMatch(person)
	if err != nil {
		t.Errorf("Failed to add new person: %s", err)
	}
	err = ms.RemoveSinglePerson("John")
	if err != nil {
		t.Errorf("Failed to remove person: %s", err)
	}

	err = ms.RemoveSinglePerson("John")
	if err != ErrPersonNotFound {
		t.Errorf("Expected person not found error, got %v", err)
	}

	result, err := ms.QuerySinglePeople(1)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}

func TestQuerySinglePeopleOnlyFemale(t *testing.T) {
	ms := NewMatchService()

	result, err := ms.QuerySinglePeople(1)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	persons := []*model.Person{
		{Name: "test1", Height: 180, Gender: model.GenderFemale, WantedDates: 3},
		{Name: "test2", Height: 170, Gender: model.GenderFemale, WantedDates: 2},
		{Name: "test3", Height: 168, Gender: model.GenderFemale, WantedDates: 2},
		{Name: "test4", Height: 2, Gender: model.GenderFemale, WantedDates: 2},
		{Name: "test5", Height: 1, Gender: model.GenderFemale, WantedDates: 2},
		{Name: "test6", Height: 170, Gender: model.GenderFemale, WantedDates: 2},
	}

	for _, p := range persons {
		_, err = ms.AddSinglePersonAndMatch(p)
		if err != nil {
			t.Errorf("Failed to add new person: %s", err)
		}
	}

	result, err = ms.QuerySinglePeople(1)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %v", result)
	}
	if result[0].Name != "test5" {
		t.Errorf("Expected test5, got %v", result[0].Name)
	}

	result, err = ms.QuerySinglePeople(5)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 5 {
		t.Errorf("Expected 5 results, got %v", result)
	}
	if result[0].Name != "test5" {
		t.Errorf("Expected test5, got %v", result[0].Name)
	}
	if result[1].Name != "test4" {
		t.Errorf("Expected test4, got %v", result[1].Name)
	}
	if result[2].Name != "test3" {
		t.Errorf("Expected test3, got %v", result[2].Name)
	}
	if result[3].Height != 170 {
		t.Errorf("Expected Height 170, got %v", result[3])
	}
	if result[4].Height != 170 {
		t.Errorf("Expected Height 170, got %v", result[3])
	}

	result, err = ms.QuerySinglePeople(6)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 6 {
		t.Errorf("Expected 6 results, got %v", result)
	}
	if result[5].Name != "test1" {
		t.Errorf("Expected test1, got %v", result[5].Name)
	}

	result, err = ms.QuerySinglePeople(7)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 6 {
		t.Errorf("Expected 6 results, got %v", result)
	}
	if result[5].Name != "test1" {
		t.Errorf("Expected test1, got %v", result[5].Name)
	}
}

func TestQuerySinglePeopleOnlyMale(t *testing.T) {
	ms := NewMatchService()

	result, err := ms.QuerySinglePeople(1)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	persons := []*model.Person{
		{Name: "test1", Height: 180, Gender: model.GenderMale, WantedDates: 3},
		{Name: "test2", Height: 170, Gender: model.GenderMale, WantedDates: 2},
		{Name: "test3", Height: 168, Gender: model.GenderMale, WantedDates: 2},
		{Name: "test4", Height: 2, Gender: model.GenderMale, WantedDates: 2},
		{Name: "test5", Height: 1, Gender: model.GenderMale, WantedDates: 2},
		{Name: "test6", Height: 170, Gender: model.GenderMale, WantedDates: 2},
	}

	for _, p := range persons {
		_, err = ms.AddSinglePersonAndMatch(p)
		if err != nil {
			t.Errorf("Failed to add new person: %s", err)
		}
	}

	result, err = ms.QuerySinglePeople(1)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %v", result)
	}
	if result[0].Name != "test1" {
		t.Errorf("Expected test1, got %v", result[0].Name)
	}

	result, err = ms.QuerySinglePeople(5)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 5 {
		t.Errorf("Expected 5 results, got %v", result)
	}
	if result[0].Name != "test1" {
		t.Errorf("Expected test1, got %v", result[0].Name)
	}
	if result[1].Height != 170 {
		t.Errorf("Expected Height 170, got %v", result[1])
	}
	if result[2].Height != 170 {
		t.Errorf("Expected Height 170, got %v", result[2])
	}
	if result[3].Name != "test3" {
		t.Errorf("Expected test3, got %v", result[3].Name)
	}
	if result[4].Name != "test4" {
		t.Errorf("Expected test4, got %v", result[4].Name)
	}

	result, err = ms.QuerySinglePeople(6)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 6 {
		t.Errorf("Expected 6 results, got %v", result)
	}
	if result[5].Name != "test5" {
		t.Errorf("Expected test5, got %v", result[5].Name)
	}

	result, err = ms.QuerySinglePeople(7)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 6 {
		t.Errorf("Expected 6 results, got %v", result)
	}
	if result[5].Name != "test5" {
		t.Errorf("Expected test5, got %v", result[5].Name)
	}
}

func TestQuerySinglePeopleMixed(t *testing.T) {
	ms := NewMatchService()

	result, err := ms.QuerySinglePeople(1)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}

	persons := []*model.Person{
		{Name: "test11", Height: 199, Gender: model.GenderMale, WantedDates: 99},
		{Name: "test12", Height: 170, Gender: model.GenderMale, WantedDates: 99},
		{Name: "test16", Height: 170, Gender: model.GenderMale, WantedDates: 99},
		{Name: "test13", Height: 168, Gender: model.GenderMale, WantedDates: 99},
		{Name: "test14", Height: 6, Gender: model.GenderMale, WantedDates: 19900},
		{Name: "test15", Height: 3, Gender: model.GenderMale, WantedDates: 99},
		{Name: "test21", Height: 180, Gender: model.GenderFemale, WantedDates: 99},
		{Name: "test23", Height: 168, Gender: model.GenderFemale, WantedDates: 99},
		{Name: "test22", Height: 150, Gender: model.GenderFemale, WantedDates: 99},
		{Name: "test24", Height: 22, Gender: model.GenderFemale, WantedDates: 99},
		{Name: "test25", Height: 1, Gender: model.GenderFemale, WantedDates: 99},
		{Name: "test26", Height: 1, Gender: model.GenderFemale, WantedDates: 99},
	}

	for _, p := range persons {
		_, err = ms.AddSinglePersonAndMatch(p)
		if err != nil {
			t.Errorf("Failed to add new person: %s", err)
		}
	}

	result, err = ms.QuerySinglePeople(1)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %v", result)
	}
	if result[0].Name != "test11" {
		t.Errorf("Expected test11, got %v", result[0].Name)
	}

	result, err = ms.QuerySinglePeople(2)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 results, got %v", result)
	}
	if result[0].Name != "test11" {
		t.Errorf("Expected test11, got %v", result[0].Name)
	}
	if result[1].Height != 1 {
		t.Errorf("Expected Height 1, got %v", result[1])
	}

	result, err = ms.QuerySinglePeople(3)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("Expected 3 results, got %v", result)
	}
	if result[2].Height != 170 {
		t.Errorf("Expected Height 170, got %v", result[2])
	}

	result, err = ms.QuerySinglePeople(5)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 5 {
		t.Errorf("Expected 5 results, got %v", result)
	}
	if result[4].Height != 170 {
		t.Errorf("Expected Height 170, got %v", result[4])
	}

	result, err = ms.QuerySinglePeople(13)
	if err != nil {
		t.Errorf("Failed to query single people: %v", err)
	}
	if len(result) != 12 {
		t.Errorf("Expected 13 results, got %v", result)
	}
	if result[10].Name != "test15" {
		t.Errorf("Expected test15, got %v", result[11].Name)
	}
	if result[11].Name != "test21" {
		t.Errorf("Expected test21, got %v", result[11].Name)
	}
}
