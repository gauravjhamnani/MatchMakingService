package main

type subjectIdList struct {
	subjectIds []string
}

type MatchMakingGameEngine struct {
	game *Game
	// both of these ratingMap are for 1V1, we can create similarly for 1VN
	ratingMapForIndividual []subjectIdList
	ratingMapForTeam []subjectIdList   
	// pendingRequests []PairingRequestWithRating  this would have been useful if there was a background threads for each MatchMaking gameEngine
}

func NewMatchMakingGameEngine(game *Game) *MatchMakingGameEngine {
	var obj MatchMakingGameEngine
	obj.game = game
	obj.ratingMapForIndividual = make([]subjectIdList, game.maxRating + 1)
	obj.ratingMapForTeam = make([]subjectIdList, game.maxRating + 1)
	return &obj
}

func (obj *MatchMakingGameEngine) pairUpUsers(request *PairingRequestWithRating) []string {
	// For 1V1
	var ratingMapObj []subjectIdList
	if request.pairingReqest.gameFormat == GAMEFORMAT(INDIVIDUAL) {
		ratingMapObj = obj.ratingMapForIndividual
	} else {
		ratingMapObj = obj.ratingMapForTeam
	}
	ratingMapObj[request.rating].subjectIds = append(ratingMapObj[request.rating].subjectIds, 
		request.pairingReqest.id)
	return searchForIds(ratingMapObj, request.pairingReqest.id, request.rating)
}

func searchForIds(ratingMapObj []subjectIdList, pairingId string, rating int) []string {
	var theta int = 1
	var cnt int = 0
	for {
		cnt = 0;
		for cnt < theta {
			if ((rating + cnt) < maxRating && (rating + cnt) > 0) //if in range 
				subjectIdList[rating + cnt]
			   // look up at ratingMapObj[rating + cnt] and 
			   //  if subjectIds are not null(excluding pairingId) remove first one from the list and return it
			if ((rating - cnt) < len(subjectIdList) && (rating - cnt) > 0) //if in range 
			   // look up at ratingMapObj[rating + cnt] and 
			   //  if subjectIds are not null(excluding pairingId) remove first one from the list and return it
			cnt++;
		}
		theta *= 2
		time.sleep(50);
	}
}