package model

type Movie struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Year        string `json:"year" bson:"year"`
	Rating      string `json:"rating" bson:"rating"`
	RatingCount string `json:"rating_count" bson:"rating_count"`
	Runtime     string `json:"runtime" bson:"runtime"`
	Genres      string `json:"genres" bson:"genres"`
	Budget      string `json:"budget" bson:"budget"`
	Gross       string `json:"gross" bson:"gross"`
	Director    string `json:"director" bson:"director"`
	Stars       string `json:"stars" bson:"stars"`
	Country     string `json:"country" bson:"country"`
	StoryLine   string `json:"story_line" bson:"story_line"`
}