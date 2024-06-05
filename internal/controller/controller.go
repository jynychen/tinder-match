package controller

import "tinder-match/internal/service"

type ControllerV1 struct {
	matchService service.MatchService
}

func NewControllerV1(matchService service.MatchService) *ControllerV1 {
	return &ControllerV1{
		matchService: matchService,
	}
}
