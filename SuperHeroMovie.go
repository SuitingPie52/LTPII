package main

type SuperHeroMovie struct {
	Movie
}

func (shm SuperHeroMovie) CollectRatings() (float32, error) {

	return 10, nil

}
