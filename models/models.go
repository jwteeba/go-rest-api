package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Movie struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string			  `json:"title" bson:"title,omitempty"`
	Year   string             `json:"year" bson:"year,omitempty"`
	Rating string             `json:"rating" bson:"rating,omitempty"`
	Genre  string             `json:"genre" bson:"genre,omitempty"`
	Description string        `json:"description" bson:"description,omitempty"`
	Time string               `json:"time" bson:"time,omitempty"`
	Star   string             `json:"star" bson:"star,omitempty"`
	Movie_Image  string       `json:"movie_image" bson:"movie_image,omitempty"`
}

