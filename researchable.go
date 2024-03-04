package main

type Researchable interface {
	CollectRatings() (float32, error)
}
