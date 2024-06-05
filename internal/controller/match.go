package controller

import (
	"errors"
	"net/http"
	"tinder-match/internal/controller/dto"
	"tinder-match/internal/service"

	"github.com/gin-gonic/gin"
)

// AddSinglePersonAndMatchHandler
// @Id AddSinglePersonAndMatch
// @Summary      add a new user and find any possible matches
// @Description  Add a new user to the matching system and find any possible matches for the new user
// @Accept       json
// @Produce      json
// @Param person body dto.AddSinglePersonAndMatchRequest true "person"
// @Success      200          {object}  dto.AddSinglePersonAndMatchResponse
// @Failure      400,409,500  {object}  gin.Error
// @Router       /api/v1/person [post]
func (c *ControllerV1) AddSinglePersonAndMatchHandler(ctx *gin.Context) {
	req := dto.AddSinglePersonAndMatchRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.Error{
			Err: err,
		})
		return
	}

	matched, err := c.matchService.AddSinglePersonAndMatch(req.Person.ToModel())
	if err != nil {
		if errors.Is(err, service.ErrPersonAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.Error{
				Err: err,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.Error{
			Err: err,
		})
		return
	}

	res := dto.AddSinglePersonAndMatchResponse{MatchedPeople: []dto.Person{}}
	for _, each := range matched {
		p := dto.Person{}
		p.FromModel(&each)
		res.MatchedPeople = append(res.MatchedPeople, p)
	}

	ctx.JSON(http.StatusOK, res)
}

// RemoveSinglePersonHandler
// @Id RemoveSinglePerson
// @Summary      Remove a user from the matching system
// @Description  Remove a user from the matching system so that the user cannot be matched anymore
// @Produce      json
// @Param name path string true "name"
// @Success      200          {object}  dto.RemoveSinglePersonResponse
// @Failure      400,500      {object}  gin.Error
// @Router       /api/v1/person/{name} [delete]
func (c *ControllerV1) RemoveSinglePersonHandler(ctx *gin.Context) {
	req := dto.RemoveSinglePersonRequest{}
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.Error{
			Err: err,
		})
		return
	}

	err = c.matchService.RemoveSinglePerson(req.Name)
	if err != nil {
		if errors.Is(err, service.ErrPersonNotFound) {
			ctx.JSON(http.StatusOK, dto.RemoveSinglePersonResponse{})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.Error{
			Err: err,
		})

		return
	}

	ctx.JSON(http.StatusOK, dto.RemoveSinglePersonResponse{})
}

// QuerySinglePeopleHandler
// @Id QuerySinglePeople
// @Summary      Query a list of users from the matching system
// @Description  Find the most N possible matched single people, where N is a request parameter
// @Produce      json
// @Param limit query int false "limit"
// @Success      200          {object}  dto.QuerySinglePeopleResponse
// @Failure      400,500      {object}  gin.Error
// @Router       /api/v1/person [get]
func (c *ControllerV1) QuerySinglePeopleHandler(ctx *gin.Context) {
	req := dto.QuerySinglePeopleRequest{}
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.Error{
			Err: err,
		})
		return
	}

	people, err := c.matchService.QuerySinglePeople(req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{
			Err: err,
		})
		return
	}

	res := dto.QuerySinglePeopleResponse{People: []dto.Person{}}
	for _, each := range people {
		p := dto.Person{}
		p.FromModel(&each)
		res.People = append(res.People, p)
	}

	ctx.JSON(http.StatusOK, res)
}
