package main

type RatingService interface {
	FetchRating(string, string) int
}

type RatingServiceImpl struct {

}

func (obj *RatingServiceImpl) FetchRating(userId string, gameId string) int {
	// this can be mocked with random numbers
	return 100;
}