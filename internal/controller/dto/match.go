package dto

type AddSinglePersonAndMatchRequest struct {
	Person
}

type AddSinglePersonAndMatchResponse struct {
	MatchedPeople []Person `json:"matched"`
}

type RemoveSinglePersonRequest struct {
	Name string `uri:"name" binding:"required"`
}

type RemoveSinglePersonResponse struct{}

type QuerySinglePeopleRequest struct {
	Limit int `form:"limit" binding:"required,min=1"`
}

type QuerySinglePeopleResponse struct {
	People []Person `json:"people"`
}
