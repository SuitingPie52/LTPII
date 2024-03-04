package main

type Movie struct {
	Title     string
	Date      string
	Genre     string
	AvgRating float32
	Reviews   []Review
}

func (m Movie) PrintAllReviews() {

	for _, r := range m.Reviews {

		r.PrintReview()

	}

}
