package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
)

type Users struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" json:"name" validate:"required,min=3"`
	Nickname string             `bson:"nickname" json:"nickname" validate:"required,min=5"`
	Avatar   image.Image        `bson:"avatar" json:"avatar"`
	TgAcc    string             `bson:"tg_acc" json:"tg_acc" validate:"required""`
	Gender   string             `bson:"gender" json:"gender" validate:"required"`
	Status   string             `bson:"status" json:"status"`
	Children bool               `bson:"children" json:"children"`
	Pets     bool               `bson:"pets" json:"pets"`
	Location string             `bson:"location" json:"location" validate:"required"`
	Password string             `bson:"password" json:"password" validate:"required"`
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
