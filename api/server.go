package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/machingclee/2023-11-04-go-gin/internal/db"
	"github.com/machingclee/2023-11-04-go-gin/token"
	"github.com/machingclee/2023-11-04-go-gin/util"
)

type Server struct {
	config     *util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config *util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}
func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.POST("/transfers", server.createTransfer)
	router.POST("/user", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
}
