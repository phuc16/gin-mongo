package user

import (
	"context"
	model "gin-mongo/src/models"
	"gin-mongo/src/mongoDb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var timeFormat = "02/01/2006:15:04:05 -0700"

func CreateUserNew(ctx context.Context, req *UserRegisterReq) UserRegisterResp {
	//Check validate
	resp := UserRegisterResp{}
	if req.Name == "" || req.Age == 0 || req.Password == "" {
		resp.Code = 400
		resp.Message = "Invalid data"
		return resp
	}
	//

	// log.Println(req)

	//Kiểm tra user đã tồn tài chưa

	_, err := mongoDb.GetUserByName(ctx, bson.M{"name": req.Name, "status": "active"})

	if err != nil && err != mongo.ErrNoDocuments {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	if !(err == mongo.ErrNoDocuments) {
		resp.Code = 409
		resp.Message = "User is exist"
		return resp
	}

	// Insert DB
	userMongo := model.User{
		Name:      req.Name,
		Age:       req.Age,
		Password:  req.Password,
		Status:    "active",
		CreatedAt: time.Now().Format(timeFormat),
	}

	err = mongoDb.CreateUserNew(ctx, userMongo)
	// if insert db err
	if err != nil {
		resp.Code = 409
		resp.Message = err.Error()
		return resp
	}

	//if success
	resp.Message = "Success"
	return resp

}

func GetAllUsers(ctx context.Context, req *UserGetAllReq) UserGetAllResp {
	resp := UserGetAllResp{}

	if req.ToDate == "" {
		tmpFromDate, err := time.Parse(timeFormat, timeFormat)

		if err != nil {
			resp.Code = 400
			resp.Message = "Bad request"
			return resp
		}

		req.FromDate = tmpFromDate.Format(timeFormat)
	}

	if req.ToDate == "" {
		// resp.Code = 400
		// resp.Message = "Invalid date"
		// return resp
		req.ToDate = time.Now().Format(timeFormat)
	}

	// log.Println(req)

	res, err := mongoDb.GetAllUsers(ctx, req.FromDate, req.ToDate)

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	resp.Code = 200
	resp.Message = "Success"
	resp.Data = res
	return resp
}

func GetUserById(ctx context.Context, req *UserGetByIdReq) UserGetByIdResp {
	resp := UserGetByIdResp{}
	// log.Println("userId", req.Id)

	if req.Id == "" {
		resp.Code = 400
		resp.Message = "Id không tồn tại"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	res, err := mongoDb.GetUserById(ctx, bson.M{"_id": objId, "status": "active"})

	if err == mongo.ErrNoDocuments {
		resp.Code = 404
		resp.Message = "Not found"
		return resp
	} else if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	resp.Code = 200
	resp.Message = "Success"
	resp.Data = res
	return resp
}

func UpdateUserById(ctx context.Context, req *UserUpdateByIdReq) UserUpdateByIdResp {
	resp := UserUpdateByIdResp{}
	// log.Println(req)

	if req.Id == "" {
		resp.Code = 400
		resp.Message = "Id không tồn tại"
		return resp
	}
	if req.Name == "" || req.Age == 0 {
		resp.Code = 400
		resp.Message = "Trường Name hoặc Age không hợp lệ"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	newUser := model.User{
		Name:     req.Name,
		Password: req.Password,
		Age:      req.Age,
		Status:   "active",
	}

	res, err := mongoDb.UpdateUserById(ctx, bson.M{"_id": objId, "status": "active"}, newUser)

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = 404
		resp.Message = "Not found"
		return resp
	}

	resp.Code = 200
	resp.Message = "Success"
	return resp
}

func DeleteUserById(ctx context.Context, req *UserDeleteByIdReq) UserDeleteByIdResp {
	resp := UserDeleteByIdResp{}
	// log.Println(req)

	if req.Id == "" {
		resp.Code = 400
		resp.Message = "Id không tồn tại"
		return resp
	}

	objId, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	res, err := mongoDb.DeleteUserById(ctx, bson.M{"_id": objId, "status": "active"}, bson.M{"status": "deleted", "created_at": time.Now().Format(timeFormat)})

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = 404
		resp.Message = "Not found"
		return resp
	}

	resp.Code = 200
	resp.Message = "Success"
	return resp
}

func Login(ctx context.Context, req *UserLoginReq) UserLoginResp {
	resp := UserLoginResp{}

	if req.Name == "" || req.Password == "" {
		resp.Code = 400
		resp.Message = "Invalid data"
		return resp
	}

	user, err := mongoDb.GetUserByName(ctx, bson.M{"name": req.Name, "status": "active"})

	if err == mongo.ErrNoDocuments {
		resp.Code = 401
		resp.Message = "User hasn't been registered"
		return resp
	} else if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	if user.Password != req.Password {
		resp.Code = 401
		resp.Message = "Password is incorrect"
		return resp
	}

	res, err := mongoDb.UserLogin(ctx, bson.M{"name": user.Name, "is_logged": false})

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = 409
		resp.Message = "User is logged"
		return resp
	}

	resp.Code = 200
	resp.Message = "Success"
	return resp
}

func Logout(ctx context.Context, req *UserLogoutReq) UserLogoutResp {
	resp := UserLogoutResp{}

	if req.Name == "" {
		resp.Code = 400
		resp.Message = "Invalid data"
		return resp
	}

	res, err := mongoDb.UserLogout(ctx, bson.M{"name": req.Name, "status": "active", "is_logged": true})

	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return resp
	}

	if res < 1 {
		resp.Code = 401
		resp.Message = "User hasn't been registered or logged"
		return resp
	}

	resp.Code = 200
	resp.Message = "Success"
	return resp

}