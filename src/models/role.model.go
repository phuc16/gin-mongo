package model

type Role struct {
	Id       string `bson:"_id" json:"-"`
	RoleCode int    `bson:"role_code" json:"roleCode"`
	RoleName string `bson:"role_name" json:"roleName"`
}
