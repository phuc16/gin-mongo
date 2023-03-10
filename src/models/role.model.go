package model

type Role struct {
	Id       string `bson:"_id" json:"-"`
	RoleCode string `bson:"role_code" json:"-"`
	RoleName string `bson:"role_name" json:"-"`
}
