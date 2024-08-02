package main

type Movie struct {
	title  string
	rating float64
}

type MovieList struct {
	Movies []Movie
}

func (ml *MovieList) AddMovie(title string, rating float64) {
	movie := Movie{
		title:  title,
		rating: rating,
	}

	ml.Movies = append(ml.Movies, movie)
}

func (ml MovieList) AverageRating() float64 {
	var totalRating float64
	for _, v := range ml.Movies {
		totalRating += v.rating
	}
	return totalRating / float64(len(ml.Movies))
}
