package model

type User struct {
	Id       string `bson:"_id" json:"id"`
	Name     string `bson:"name" json:"name"`
	FullName string `bson:"full_name" json:"fullName"`
	Password string `bson:"password" json:"password"`
	Age      int    `bson:"age" json:"age"`

	Status    string `bson:"status" json:"-"`
	IsLogged  bool   `bson:"is_logged" json:"-"`
	CreatedAt string `bson:"created_at" json:"createdAt"`
	UpdatedAt string `bson:"updated_at" json:"updatedAt"`
}
