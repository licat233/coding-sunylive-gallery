package server

import (
	"log"
	"sunylive-gallery/global"
	"sunylive-gallery/server/logic"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Logic *logic.Logic
	Gin   *gin.Engine
}

func NewServer() *Server {

	return &Server{
		Logic: logic.NewLogic(),
		Gin:   gin.New(),
	}
}

func (s *Server) Initialize() {
	//set mode
	if global.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	//close ip proxy
	s.Gin.SetTrustedProxies(nil)
	//session
	store := cookie.NewStore([]byte("licat23333333"))
	s.Gin.Use(sessions.Sessions("sunylivesession", store))
	//middleware
	if !global.ReleaseMode {
		s.Gin.Use(gin.Logger())
	}
	s.Gin.Use(gin.Recovery())
	//static file
	s.Gin.LoadHTMLGlob("view/*")
}

func (s *Server) Router() {
	s.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s.Gin.GET("/", s.Logic.IndexPage)
	s.Gin.GET("/login", s.Logic.LoginPage)
	api := s.Gin.Group("/api")
	{
		api.POST("/login", s.Logic.Login)
		api.GET("/images", s.Logic.Images)
	}
	// baseAuth := gin.BasicAuth(gin.Accounts{
	// 	global.Config.Username: global.Config.Password,
	// })
	admin := s.Gin.Group("/admin", AuthMiddleware())
	{
		admin.GET("/", s.Logic.AdminPage)
		admin.GET("/addUrl", s.Logic.AddUrlsPage)
		admin.POST("/addUrl", s.Logic.AddUrls)
	}
}

func (s *Server) Run() {
	s.Initialize()
	s.Router()
	if err := s.Gin.Run(global.Config.Addr); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
