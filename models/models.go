package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Email string             `bson:"email,omitempty"`
	Phone string             `bson:"phone,omitempty"`
}

type Book struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Owner           primitive.ObjectID `bson:"owner,omitempty"`
	Name            string             `bson:"name,omitempty"`
	Author          string             `bson:"author,omitempty"`
	PublicationDate string             `bson:"publicationDate,omitempty"`
	Sinopse         string             `bson:"sinopse,omitempty"`
}
