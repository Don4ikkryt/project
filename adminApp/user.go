package main

type User struct {
	Name     string `json:"name" form:"name" bson:"name"`
	Age      int    `json:"age" form:"age" bson:"age"`
	Gender   string `json:"gender" form:"gender" bson:"gender"`
	ID       uint   `json:"id" form:"id" bson:"_id"`
	Cats     []Cat  `json:"cats" form:"cats" bson:"cats"`
	Username string `json:"username" form:"username" bson:"username"`
	Password string `json:"password" form:"password" bson:"password"`
	Message  string `json:"message" form:"message" bson:"message"`
}
