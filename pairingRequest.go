package main

import "time"


type GAMEFORMAT int

const (
	SINGLE GAMEFORMAT = iota   // 1V1
	MULTI					   // 1VN
)


type ParticipantType int

const (
	INDIVIDUAL ParticipantType = iota
	TEAM
)

type PairingRequest struct {
	id string
	subjectId string
	subjectType ParticipantType
	gameId string
	gameFormat GAMEFORMAT
	createdOn time.Time
}

func NewPairingRequest(idParam, subjectIdParam string, subjectTypeParam ParticipantType, gameIdParam string) *PairingRequest {
	var request PairingRequest
	request.id = idParam;
	request.subjectId = subjectIdParam
	request.subjectType = subjectTypeParam
	request.gameId = gameIdParam
	return &request
}

type PairingRequestWithRating struct {
	pairingReqest PairingRequest
	rating int
}

func NewPairingRequestWithRating(probj PairingRequest, ratingValue int) *PairingRequestWithRating {
	var request PairingRequestWithRating 
	request.pairingReqest = probj
	request.rating = ratingValue
	return &request
}