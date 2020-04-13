package main

type Cat struct {
	Name   string `json:"name" form:"name" bson:"name"`
	Age    int    `json:"age" form:"age" bson:"age"`
	Breed  string `json:"breed" form:"breed" bson:"breed"`
	Gender string `json:"gender" form:"gender" bson:"gender"`
	Color  string `json:"color" form:"color" bson:"color"`
	ID     uint   `json:"id" form:"id" bson:"_id"`
}
