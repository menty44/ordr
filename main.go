package main

import (
	"github.com/gin-gonic/gin"
	"github.com/menty44/ordr/controller"
	"github.com/menty44/ordr/service"
	"net/http"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})
	// godotenv package
	port := goDotEnvVariable("PORT")
	//port := os.Getenv("PORT")
	// Elastic Beanstalk forwards requests to port default : 5000
	if port == "" {
		port = "5000"
	}
	r.Run(":" + port)
}
