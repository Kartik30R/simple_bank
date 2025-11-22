package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/kartik30R/simple_bank/db/sqlc"
	"github.com/kartik30R/simple_bank/token"
	"github.com/kartik30R/simple_bank/utils"
)

type Server struct {
	config utils.Config
	store *db.Store
	tokenMaker token.Maker
	router *gin.Engine
}


func NewServer(config utils.Config,store *db.Store) (*Server, error ){
tokenMaker , err := token.NewJWTMaker(config.TokenSymmetricKey)
if err!=nil{
	return nil, fmt.Errorf("can not create token %d",err)
}

	server := &Server{
		config: config,
		store:store,
		tokenMaker: tokenMaker,
	}
 
	if v, ok :=binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency", currencyValidator)
	}

 server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/token/renew_access", server.renewAccessToken)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
return server.router.Run(address);
}


func errorResponse(err error) gin.H {
    return gin.H{"error": err.Error()}
}
