package model

type User struct {
	Name      string `bson:"name" json:"name"`
	Password  string `bson:"password" json:"password"`
	Age       int    `bson:"age" json:"age"`
	Status    string `bson:"status" json:"-"`
	CreatedAt string `bson:"created_at" json:"-"`
}
