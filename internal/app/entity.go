package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name,omitempty" json:"name" validate:"min=5"`
	Nickname  string             `bson:"nickname,omitempty" json:"nickname" validate:"min=5"`
	Gender    string             `bson:"gender" json:"gender"`
	Status    string             `bson:"status" json:"status"`
	Childrens bool               `bson:"childrens" json:"childrens"`
	Pets      bool               `bson:"pets" json:"pets"`
	Location  string             `bson:"location" json:"location"`
	Password  string             `bson:"password" json:"password"`
}
type Specialists struct {
	ID        primitive.ObjectID `bson:"_id"`
	Firstname string             `bson:"firstname" json:"firstname" validate:"min=2"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname" validate:"min=3"`
	Gender    string             `bson:"gender" json:"gender"`
	Location  string             `bson:"location,omitempty" json:"location"`
	Sphere    string             `bson:"sphere" json:"sphere"`
	Password  string             `bson:"password" json:"password"`
}
type Support struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" json:"name" validate:"min=3"`
	Nickname string             `bson:"nickname,omitempty" json:"nickname" validate:"min=3"`
	Gender   string             `bson:"gender" json:"gender"`
	Password string             `bson:"password" json:"password"`
}

type Volunteers struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name,omitempty" json:"name" validate:"min=3"`
	Nickname string             `bson:"nickname" json:"nickname" validate:"min=3"`
	Gender   string             `bson:"gender" json:"gender"`
	Password string             `bson:"password" json:"password"`
}
