package main

import "fmt"

type Review struct {
	Movie  Movie
	Critic Critic
	Rating float32
	Text   string
}

func (r Review) PrintReview() {

	fmt.Printf("Filme: %s\nCrítico: %s\nNota: %f\nTexto: %s\n\n", r.Movie.Title, r.Critic.Name, r.Rating, r.Text)

}
