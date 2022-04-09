package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type Users struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" json:"name"`
	Nickname string             `bson:"nickname" json:"nickname" validate:"required,min=3"`
	TgAcc    string             `bson:"tg_acc" json:"tg_acc"`
	Gender   string             `bson:"gender" json:"gender"`
	Status   string             `bson:"status" json:"status"`
	Children bool               `bson:"children" json:"children"`
	Pets     bool               `bson:"pets" json:"pets"`
	Location string             `bson:"location" json:"location"`
	Password string             `bson:"password" json:"password" validate:"required"`
}
type Specialists struct {
	ID        primitive.ObjectID `bson:"_id"`
	Firstname string             `bson:"firstname" json:"firstname" validate:"required,min=2"`
	Lastname  string             `bson:"lastname" json:"lastname" validate:"min=3"`
	Nickname  string             `bson:"nickname" json:"nickname" validate:"required,min=3"`
	Gender    string             `bson:"gender" json:"gender"`
	Location  string             `bson:"location" json:"location"`
	Sphere    string             `bson:"sphere" json:"sphere"`
	Password  string             `bson:"password" json:"password" validate:"required"`
}
type Support struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" json:"name" validate:"min=3"`
	Nickname string             `bson:"nickname" json:"nickname" validate:"required,min=3"`
	Gender   string             `bson:"gender" json:"gender"`
	Location string             `bson:"location" json:"location"`
	Password string             `bson:"password" json:"password" validate:"required"`
}

type Volunteers struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" json:"name"`
	Nickname string             `bson:"nickname" json:"nickname" validate:"required,min=3"`
	Password string             `bson:"password" json:"password" validate:"required"`
	Location string             `bson:"location" json:"location"`
	Gender   string             `bson:"gender" json:"gender"`
}

func ValidateUserStruct(user Users) error {
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}
	return nil
}

func ValidateVolunteerStruct(volunteer Volunteers) error {
	validate = validator.New()
	err := validate.Struct(volunteer)
	if err != nil {
		return err
	}
	return nil
}

func ValidateSpecialistStruct(specialist Specialists) error {
	validate = validator.New()
	err := validate.Struct(specialist)
	if err != nil {
		return err
	}
	return nil
}

func ValidateSupportStruct(support Support) error {
	validate = validator.New()
	err := validate.Struct(support)
	if err != nil {
		return err
	}
	return nil
}
