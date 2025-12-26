package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title           Swagger Example API
// @version         1.0
// @description     这是一个 Swag 示例项目
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	r := gin.Default()

	// 注册 Swagger 路由
	// 访问地址将会是：http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/user/:name", GetUser)
	}

	r.Run(":8080")
}

// GetUser 这是一个示例 Handler
// @Summary      获取用户信息
// @Description  根据用户名获取详细信息
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        name   path      string  true  "用户名"
// @Success      200    {object}  UserResponse
// @Failure      400    {object}  ErrorResponse
// @Router       /user/{name} [get]
func GetUser(c *gin.Context) {
	name := c.Param("name")

	// 模拟业务逻辑
	if name == "admin" {
		c.JSON(http.StatusOK, UserResponse{
			Name: "Admin User",
			Age:  30,
		})
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    400,
			Message: "User not found",
		})
	}
}

// --- 定义用于文档展示的结构体 (DTO/VO) ---

type UserResponse struct {
	Name string `json:"name" example:"John Doe"`
	Age  int    `json:"age" example:"25"`
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
