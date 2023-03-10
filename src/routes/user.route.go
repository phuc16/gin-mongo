package routes

import (
	"fmt"
	"gin-mongo/src/controllers/user"
	middlewares "gin-mongo/src/middlewares"
	"gin-mongo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ResCode = utils.ResCode

func UserRoutes(routes *gin.Engine) {
	notAuthRoutes := routes.Group("api/user")
	{
		notAuthRoutes.POST("/register", createUserNew) //nên để tên api đàng hoàng nha em . /user/createUser
		notAuthRoutes.GET("/getAll", getAllUsers)      //user/getAll
		notAuthRoutes.GET("/getById", getUserById)     // hạn chế sử dụng kiểu url như này nha em, để key - value 1 phần sẽ bảo mật hơn vì có key nữa, và no cung k ảnh hưởng trực tiếp đến url,
		// // với ko ai để id lên param nhen, thường sẽ để trong body. //sua lai query param giup anh nhen

		notAuthRoutes.PUT("/updateById", updateUserById)  //sua lai query param giup anh nhen
		notAuthRoutes.POST("/deleteById", deleteUserById) //sua lai query param giup anh nhen
		notAuthRoutes.POST("/changePassword", changePassword)

		notAuthRoutes.POST("/login", login)
		notAuthRoutes.GET("/search", getUserByKey)
		// notAuthRoutes.GET("/getRole", getRole)
		notAuthRoutes.GET("/getRolesList", getRolesList)
	}

	authRoutes := routes.Group("api/user", middlewares.Authenticate())
	{
		authRoutes.POST("/logout", logout)
		authRoutes.GET("/profile", getUserProfile)
		authRoutes.PUT("/updateRole", updateRole)
	}
}

func createUserNew(c *gin.Context) {
	var (
		request  = &user.UserRegisterReq{}
		response = user.UserRegisterResp{}
		r        = c.Request
	)
	if err := c.Bind(&request); err != nil {
		//
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.CreateUserNew(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func getAllUsers(c *gin.Context) {
	var (
		request  = &user.UserGetAllReq{}
		response = user.UserGetAllResp{}
		r        = c.Request
	)

	if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.GetAllUsers(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func getUserById(c *gin.Context) {
	var (
		request  = &user.UserGetByIdReq{}
		response = user.UserGetByIdResp{}
		r        = c.Request
	)

	if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.GetUserById(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func updateUserById(c *gin.Context) {
	var (
		request  = &user.UserUpdateByIdReq{}
		response = user.UserUpdateByIdResp{}
		r        = c.Request
	)

	if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.UpdateUserById(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func deleteUserById(c *gin.Context) {
	var (
		request  = &user.UserDeleteByIdReq{}
		response = user.UserDeleteByIdResp{}
		r        = c.Request
	)
	if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.DeleteUserById(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func login(c *gin.Context) {
	var (
		request  = &user.UserLoginReq{}
		response = user.UserLoginResp{}
		r        = c.Request
	)

	if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.Login(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func logout(c *gin.Context) {
	var (
		request  = &user.UserLogoutReq{}
		response = user.UserLogoutResp{}
		r        = c.Request
	)

	claims, err := utils.ExtractTokenId(c)
	request.Token = utils.ExtractToken(c)
	request.Id = fmt.Sprintf("%v", claims["userId"])

	if err != nil {
		response.Code = ResCode.Unauthorized
		response.Message = "Unauthorized"
		c.JSON(response.Code, response)
	} else {
		response = user.Logout(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func getUserByKey(c *gin.Context) {
	var (
		request  = &user.UserGetByKeyReq{}
		response = user.UserGetByKeyResp{}
		r        = c.Request
	)

	if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.GetUserByKey(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func getUserProfile(c *gin.Context) {
	var (
		request  = &user.UserGetProfileReq{}
		response = user.UserGetProfileResp{}
		r        = c.Request
	)

	claims, err := utils.ExtractTokenId(c)
	request.Id = fmt.Sprintf("%v", claims["userId"])

	if err != nil {
		response.Code = ResCode.Unauthorized
		response.Message = "Unauthorized"
		c.JSON(response.Code, response)
	} else {
		response = user.GetUserProfile(r.Context(), request)
		c.JSON(response.Code, response)
	}
}

func updateRole(c *gin.Context) {
	var (
		request  = &user.UserUpdateRoleReq{}
		response = user.UserUpdateRoleResp{}
		r        = c.Request
	)

	claims, err := utils.ExtractTokenId(c)

	if err != nil {
		response.Code = ResCode.Unauthorized
		response.Message = "Unauthorized"
		c.JSON(response.Code, response)
	} else if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		request.Id = fmt.Sprintf("%v", claims["userId"])
		roleCode, err := strconv.Atoi(fmt.Sprintf("%v", claims["roleCode"]))

		if err != nil {
			response.Code = ResCode.BadRequest
			response.Message = err.Error()
			c.JSON(response.Code, response)
		} else {
			request.RoleCode = roleCode
			response = user.UpdateRole(r.Context(), request)
			c.JSON(response.Code, response)
		}
	}
}

// func getRole(c *gin.Context) {
// 	var (
// 		request  = &user.UserGetRoleReq{}
// 		response = user.UserGetRoleResp{}
// 		r        = c.Request
// 	)

// 	if err := c.Bind(&request); err != nil {
// 		response.Code = ResCode.BadRequest
// 		response.Message = err.Error()
// 		c.JSON(response.Code, response)
// 	} else {
// 		response = user.GetRole(r.Context(), request)
// 		c.JSON(response.Code, response)
// 	}
// }

func getRolesList(c *gin.Context) {
	var (
		response = user.UserGetRolesListResp{}
		r        = c.Request
	)

	response = user.GetRolesList(r.Context())
	c.JSON(response.Code, response)
}

func changePassword(c *gin.Context) {
	var (
		request  = &user.UserChangePasswordReq{}
		response = user.UserChangePasswordResp{}
		r        = c.Request
	)

	if err := c.Bind(&request); err != nil {
		response.Code = ResCode.BadRequest
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		response = user.ChangePassword(r.Context(), request)
		c.JSON(response.Code, response)
	}
}
