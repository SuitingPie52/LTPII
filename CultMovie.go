package main

import "errors"

type CultMovie struct {
	Movie
}

func (cm CultMovie) CollectRatings() (float32, error) {

	var sum_ratings, total_reviews float32
	sum_ratings = 0

	if cm.Reviews == nil {

		return 0, errors.New("não há nenhuma crítica")

	}

	for _, r := range cm.Reviews {

		sum_ratings += r.Rating
		total_reviews++

	}

	return sum_ratings / total_reviews, nil

}
