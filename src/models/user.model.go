package model

type User struct {
	Name      string `bson:"name" json:"name"`
	FullName  string `bson:"full_name" json:"fullName"`
	Password  string `bson:"password" json:"password"`
	Age       int    `bson:"age" json:"age"`
	Status    string `bson:"status" json:"-"`
	CreatedAt string `bson:"created_at" json:"-"`
	UpdatedAt string `bson:"updated_at" json:"-"`
}
