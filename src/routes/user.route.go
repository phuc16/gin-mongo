package routes

import (
	"gin-mongo/src/controllers/user"
	middlewares "gin-mongo/src/middlewares"
	token "gin-mongo/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	notAuthRoutes := routes.Group("api/user")
	{
		notAuthRoutes.POST("/register", createUserNew) //nên để tên api đàng hoàng nha em . /user/createUser
		notAuthRoutes.GET("/getAll", getAllUsers)      //user/getAll
		notAuthRoutes.GET("/getById", getUserById)     // hạn chế sử dụng kiểu url như này nha em, để key - value 1 phần sẽ bảo mật hơn vì có key nữa, và no cung k ảnh hưởng trực tiếp đến url,
		// // với ko ai để id lên param nhen, thường sẽ để trong body. //sua lai query param giup anh nhen

		notAuthRoutes.PUT("/updateById", updateUserById)  //sua lai query param giup anh nhen
		notAuthRoutes.POST("/deleteById", deleteUserById) //sua lai query param giup anh nhen

		notAuthRoutes.POST("/login", login)
		notAuthRoutes.GET("/search", getUserByKey)
	}
	authRoutes := routes.Group("api/user", middlewares.Authenticate())
	{
		authRoutesg.POST("/logout", logout)
		authRoutes.GET("/profile", getUserProfile)
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
		response.Code = 400
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
		response.Code = 400
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		log.Println(request)
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
		response.Code = 400
		response.Message = err.Error()
		c.JSON(response.Code, response)
	} else {
		log.Println(request)
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
		response.Code = 400
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
		response.Code = 400
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
		response.Code = 400
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
		err      error
	)
	request.Id, err = token.ExtractTokenName(c)
	request.Token = token.ExtractToken(c)

	if err != nil {
		response.Code = 401
		response.Message = "Unauthorized"
		c.JSON(response.Code, response)
	}

	response = user.Logout(r.Context(), request)
	c.JSON(response.Code, response)
}

func getUserByKey(c *gin.Context) {
	var (
		request  = &user.UserGetByKeyReq{}
		response = user.UserGetByKeyResp{}
		r        = c.Request
	)

	if err := c.Bind(&request); err != nil {
		response.Code = 400
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
		err      error
	)

	log.Println("aaa")
	request.Id, err = token.ExtractTokenName(c)

	if err != nil {
		response.Code = 401
		response.Message = "Unauthorized"
		c.JSON(response.Code, response)
	}

	response = user.GetUserProfile(r.Context(), request)
	c.JSON(response.Code, response)
}
