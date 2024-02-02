package gateway

import (
	"fmt"
	"net/http"
	"os"

	"github.com/EMSI-zero/go-chat/controller/rest"
	"github.com/EMSI-zero/go-chat/gateway/user"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/override/docs"
)

var ListenAddress = "LISTEN_ADDRESS"

//	@title			go-chat API
//	@version		1.0
//	@description	Webshop Admin Panel API Service
//	@termsOfService	http://swagger.io/terms/

//	@contact.email	emad.fdws@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func CreateHTTPServer(r registry.ServiceRegistry) (*http.Server, error) {

	engine := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	docs.SwaggerInfo.BasePath = "/api/v1"

	// TODO: fix for production
	corsHandler := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	engine.Use(corsHandler)

	AddAdminRoutes(engine, rest.NewController(r))

	if os.Getenv("ENVIRONMENT") != "production" {
		engine.GET("/admin/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	address, err := GetListenAddressEnv("admin")
	if err != nil {
		return nil, err
	}

	server := &http.Server{Addr: address, Handler: engine}
	return server, nil
}

func GetListenAddressEnv(serv string) (string, error) {
	var address string
	address = os.Getenv(ListenAddress)
	return address, nil
}

func AddAdminRoutes(e *gin.Engine, c *rest.Controller) {
	r := e.Group(fmt.Sprintf("/api/v1/%s", "admin"))
	r.Use(c.HandleError())
	r.Use(c.AuthController.Authenticate())
	user.AddRoutes(r, c.UserController, c.AuthController.SetByPassPolicy)
}
