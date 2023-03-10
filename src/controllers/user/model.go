package user

import (
	model "gin-mongo/src/models"
)

type UserRegisterReq struct { //request API
	Name     string `form:"name" json:"name"`
	FullName string `form:"fullName" json:"fullName"`
	Age      int    `form:"age" json:"age"`
	Password string `form:"password" json:"password"`
	RoleCode int    `form:"roleCode" json:"roleCode"`
}

type UserRegisterResp struct { //response API
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserGetAllReq struct {
	FromDate string `form:"fromDate"`
	ToDate   string `form:"toDate"`
}

type UserGetAllResp struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []model.User `json:"data"`
}

type UserGetByIdReq struct {
	Id string `form:"id"`
}

type UserGetByIdResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    model.User `json:"data"`
}

type UserUpdateByIdReq struct {
	Id       string `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	FullName string `form:"fullName" json:"fullName"`
	Age      int    `form:"age" json:"age"`
}

type UserUpdateByIdResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserDeleteByIdReq struct {
	Id string `json:"id"`
}

type UserDeleteByIdResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserLoginReq struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
}

type UserLoginResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type UserLogoutReq struct {
	Id    string
	Token string
}

type UserLogoutResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserGetByKeyReq struct {
	Search string `form:"search"`
}

type UserGetByKeyResp struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []model.User `json:"data"`
}
type UserChangePasswordReq struct {
	Name        string `form:"name" json:"name"`
	OldPassword string `form:"oldPassword" json:"oldPassword"`
	NewPassword string `form:"newPassword" json:"newPassword"`
}
type UserChangePasswordResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type UserGetProfileReq struct {
	Id string
}

type UserGetProfileResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    model.User `json:"data"`
}
type UserUpdateRoleReq struct {
	Id              string `form:"-" json:"-"`                             // id of the user
	UpdatedId       string `form:"updatedId" json:"updatedId"`             // id of the user need to update the role
	RoleCode        int    `json:"-"`                                      // role of the user
	UpdatedRoleCode int    `form:"updatedRoleCode" json:"updatedRoleCode"` // role of the user need to update the role
}

type UserUpdateRoleResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserGetRoleReq struct {
	Id string `form:"id" json:"id"`
}

type UserGetRoleResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    model.Role `json:"data"`
}
