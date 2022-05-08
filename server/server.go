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
}

func NewServer() *Server {
	return &Server{
		Logic: logic.NewLogic(),
	}
}

func (s *Server) Run() {
	r := gin.New()
	//close ip
	r.SetTrustedProxies(nil)
	//session
	store := cookie.NewStore([]byte("licat23333333"))
	r.Use(sessions.Sessions("sunylivesession", store))
	//middleware
	r.Use(gin.Recovery())
	//static file
	r.LoadHTMLGlob("view/*")
	//router
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", s.Logic.IndexPage)
	r.GET("/login", s.Logic.LoginPage)
	api := r.Group("/api")
	{
		api.POST("/login", s.Logic.Login)
		api.GET("/images", s.Logic.Images)
		api.POST("/addUrl", s.Logic.AddUrls)
	}

	if err := r.Run(global.Config.Addr); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
