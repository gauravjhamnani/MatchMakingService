package main

type MatchMakingService interface {
	findMatch(PairingRequest) []string   // this is pairingRequestIds
}

type MatchMakingServiceImpl struct {
	teamIdTeamMap map[string]*Team
	gameIdMatchMakingEngineMap map[string]*MatchMakingGameEngine
	gameIdGameMap map[string]*Game
}

func NewMatchMakingServiceImpl() *MatchMakingServiceImpl {
	var obj MatchMakingServiceImpl
	obj.teamIdTeamMap = make(map[string]*Team)
	obj.gameIdMatchMakingEngineMap = make(map[string]*MatchMakingGameEngine)
	obj.gameIdGameMap = make(map[string]*Game)
	return &obj
}

func (obj *MatchMakingServiceImpl) findMatch(request PairingRequest) []string {
	var ratingServiceImplObj RatingServiceImpl
	var ratingServiceObj RatingService = &ratingServiceImplObj // using Rating service implementation
	var ratingValue int
	if request.gameFormat == GAMEFORMAT(INDIVIDUAL) {
		ratingValue = ratingServiceObj.FetchRating(request.subjectId, request.gameId)
	} else {
		team := obj.teamIdTeamMap[request.subjectId]
		var cummulativeRating int
		for _, userId := range team.userIds {
			cummulativeRating += ratingServiceObj.FetchRating(userId, request.gameId)
		}
		ratingValue = cummulativeRating / len(team.userIds)
	}
	engineObj :=  obj.gameIdMatchMakingEngineMap[request.gameId]
	return engineObj.pairUpUsers(NewPairingRequestWithRating(request, ratingValue))
}