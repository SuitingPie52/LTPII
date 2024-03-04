package main

import "fmt"

func main() {

	shm := SuperHeroMovie{
		Movie{
			Title:     "Avengers",
			Date:      "2010",
			Genre:     "Action",
			AvgRating: 0,
			Reviews:   nil,
		},
	}

	cm := CultMovie{
		Movie{
			Title:     "Mulholland Drive",
			Date:      "2001",
			Genre:     "Mystery",
			AvgRating: 0,
			Reviews:   nil,
		},
	}

	r := Review{
		Movie:  cm.Movie,
		Critic: Critic{},
		Rating: 5,
		Text:   "muito foda",
	}

	r2 := Review{
		Movie:  shm.Movie,
		Critic: Critic{},
		Rating: 5,
		Text:   "muito paia",
	}

	cm.Reviews = append(cm.Reviews, r)
	shm.Reviews = append(shm.Reviews, r2)

	var putz error
	cm.AvgRating, putz = cm.CollectRatings()

	fmt.Println(cm.AvgRating, putz)

	shm.AvgRating, putz = shm.CollectRatings()

	fmt.Println(shm.AvgRating, putz)
}
