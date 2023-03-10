package user

import (
	"context"
	model "gin-mongo/src/models"
	"gin-mongo/src/mongoDb"
	utils "gin-mongo/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var timeFormat = utils.TimeFormat
var ResCode = utils.ResCode

func CreateUserNew(ctx context.Context, req *UserRegisterReq) UserRegisterResp {
	//Check validate
	resp := UserRegisterResp{}
	if req.Name == "" || req.FullName == "" || req.Age == 0 || req.Password == "" || req.RoleCode == 0 {
		resp.Code = ResCode.BadRequest
		resp.Message = "Invalid data"
		return resp
	}
	//

	//Kiểm tra user đã tồn tài chưa

	_, err := mongoDb.GetUserByName(ctx, bson.M{"name": req.Name, "status": "active"})

	if err != nil && err != mongo.ErrNoDocuments {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if !(err == mongo.ErrNoDocuments) {
		resp.Code = ResCode.Conflict
		resp.Message = "User is exist"
		return resp
	}

	password, err := utils.HashPassword(req.Password)

	if err != nil {
		resp.Code = ResCode.ServerError
		resp.Message = "Internal server error"
		return resp
	}

	// Insert DB
	userMongo := model.User{
		Name:      req.Name,
		FullName:  req.FullName,
		Age:       req.Age,
		RoleCode:  req.RoleCode,
		Password:  password,
		Status:    "active",
		IsLogged:  false,
		CreatedAt: time.Now().Format(timeFormat),
		UpdatedAt: time.Now().Format(timeFormat),
	}

	err = mongoDb.CreateUserNew(ctx, userMongo)
	// if insert db err
	if err != nil {
		resp.Code = ResCode.Conflict
		resp.Message = err.Error()
		return resp
	}

	//if success
	resp.Code = ResCode.Success
	resp.Message = "Success"
	return resp

}

func GetAllUsers(ctx context.Context, req *UserGetAllReq) UserGetAllResp {
	resp := UserGetAllResp{}

	if req.FromDate == "" {
		tmpFromDate, err := time.Parse(timeFormat, timeFormat)

		if err != nil {
			resp.Code = ResCode.BadRequest
			resp.Message = "Bad request"
			return resp
		}

		req.FromDate = tmpFromDate.Format(timeFormat)
	}

	if req.ToDate == "" {
		// resp.Code = ResCode.BadRequest
		// resp.Message = "Invalid date"
		// return resp
		req.ToDate = time.Now().Format(timeFormat)
	}

	// log.Println(req)

	res, err := mongoDb.GetAllUsers(ctx, req.FromDate, req.ToDate)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	resp.Data = res
	return resp
}

func GetUserById(ctx context.Context, req *UserGetByIdReq) UserGetByIdResp {
	resp := UserGetByIdResp{}
	// log.Println("userId", req.Id)

	if req.Id == "" {
		resp.Code = ResCode.BadRequest
		resp.Message = "Id không tồn tại"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	res, err := mongoDb.GetUserById(ctx, objId)

	if err == mongo.ErrNoDocuments {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	} else if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	resp.Data = res
	return resp
}

func UpdateUserById(ctx context.Context, req *UserUpdateByIdReq) UserUpdateByIdResp {
	resp := UserUpdateByIdResp{}
	// log.Println(req)

	if req.Id == "" {
		resp.Code = ResCode.BadRequest
		resp.Message = "Id không tồn tại"
		return resp
	}
	if req.FullName == "" || req.Password == "" || req.Age == 0 {
		resp.Code = ResCode.BadRequest
		resp.Message = "Invalid data"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	password, err := utils.HashPassword(req.Password)

	if err != nil {
		resp.Code = ResCode.ServerError
		resp.Message = "Internal server error"
		return resp
	}

	newUser := model.User{
		FullName:  req.FullName,
		Password:  password,
		Age:       req.Age,
		UpdatedAt: time.Now().Format(timeFormat),
	}

	filter := bson.M{
		"_id":    objId,
		"status": "active",
	}

	res, err := mongoDb.UpdateUserById(ctx, filter, newUser)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	return resp
}

func DeleteUserById(ctx context.Context, req *UserDeleteByIdReq) UserDeleteByIdResp {
	resp := UserDeleteByIdResp{}
	// log.Println(req)

	if req.Id == "" {
		resp.Code = ResCode.BadRequest
		resp.Message = "Id không tồn tại"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	res, err := mongoDb.DeleteUserById(ctx, objId)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	return resp
}

func Login(ctx context.Context, req *UserLoginReq) UserLoginResp {
	resp := UserLoginResp{}

	if req.Name == "" || req.Password == "" {
		resp.Code = ResCode.BadRequest
		resp.Message = "Invalid data"
		return resp
	}

	user, err := mongoDb.GetUserByName(ctx, bson.M{"name": req.Name, "status": "active"})

	if err == mongo.ErrNoDocuments {
		resp.Code = ResCode.Unauthorized
		resp.Message = "User hasn't been registered"
		return resp
	} else if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if utils.CheckPasswordHash(req.Password, user.Password) == false {
		resp.Code = ResCode.Unauthorized
		resp.Message = "Password is incorrect"
		return resp
	}

	res, err := mongoDb.UserLogin(ctx, bson.M{"name": user.Name, "is_logged": false})

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = ResCode.Conflict
		resp.Message = "User is logged"
		return resp
	}

	// log.Println(user)
	token, err := utils.GenerateToken(user)

	if err != nil {
		resp.Code = ResCode.Conflict
		resp.Message = err.Error()
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(user.Id)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	res, err = mongoDb.CreateNewToken(ctx, objId, token)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = ResCode.ServerError
		resp.Message = "Internal server error"
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	resp.Data = token
	return resp
}

func Logout(ctx context.Context, req *UserLogoutReq) UserLogoutResp {
	resp := UserLogoutResp{}

	if req.Id == "" {
		resp.Code = ResCode.BadRequest
		resp.Message = "Invalid data"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	res, err := mongoDb.UserLogout(ctx, objId)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = ResCode.Unauthorized
		resp.Message = "User hasn't been registered or logged"
		return resp
	}

	res, err = mongoDb.DeleteToken(ctx, req.Token)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = ResCode.Unauthorized
		resp.Message = "User hasn't been registered or logged"
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	return resp

}

func GetUserByKey(ctx context.Context, req *UserGetByKeyReq) UserGetByKeyResp {
	resp := UserGetByKeyResp{}

	res, err := mongoDb.GetUserByKey(ctx, req.Search)

	if err == mongo.ErrNoDocuments {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	} else if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	resp.Data = res
	return resp
}

func GetUserProfile(ctx context.Context, req *UserGetProfileReq) UserGetProfileResp {
	resp := UserGetProfileResp{}
	// log.Println("userId", req.Id)

	if req.Id == "" {
		resp.Code = ResCode.BadRequest
		resp.Message = "Id không tồn tại"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	res, err := mongoDb.GetUserById(ctx, objId)

	if err == mongo.ErrNoDocuments {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	} else if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	resp.Data = res
	return resp
}

func UpdateRole(ctx context.Context, req *UserUpdateRoleReq) UserUpdateRoleResp {
	resp := UserUpdateRoleResp{}

	if req.RoleCode < 2 {
		resp.Code = ResCode.Forbidden
		resp.Message = "Forbidden"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.UpdatedId)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	res, err := mongoDb.UpdateRole(ctx, objId, req.UpdatedRoleCode)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	return resp
}

func GetRole(ctx context.Context, req *UserGetRoleReq) UserGetRoleResp {
	resp := UserGetRoleResp{}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	user, err := mongoDb.GetUserById(ctx, objId)

	if err == mongo.ErrNoDocuments {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	} else if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	role, err := mongoDb.GetRole(ctx, user.RoleCode)

	if err == mongo.ErrNoDocuments {
		resp.Code = ResCode.NotFound
		resp.Message = "Not found"
		return resp
	} else if err != nil {
		resp.Code = ResCode.BadRequest
		resp.Message = err.Error()
		return resp
	}

	resp.Code = ResCode.Success
	resp.Message = "Success"
	resp.Data = role
	return resp
}
