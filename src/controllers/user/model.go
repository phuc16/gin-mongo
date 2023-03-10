package user

import (
	model "gin-mongo/src/models"
)

type UserRegisterReq struct { //request API
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	RoleCode int    `json:"roleCode"`
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
	Id       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Password string `json:"password"`
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
	Name     string `json:"name"`
	Password string `json:"password"`
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
