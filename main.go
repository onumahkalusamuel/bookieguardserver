package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/api"
	"github.com/onumahkalusamuel/bookieguardserver/internal/api/middleware"
	"github.com/onumahkalusamuel/bookieguardserver/internal/db"
	"github.com/onumahkalusamuel/bookieguardserver/internal/web"
)

func main() {

	// http.HandleFunc("/activation", cmd.Activation)

	// http.HandleFunc("/hosts-upload", cmd.HostsUpload)

	// http.HandleFunc("/update", cmd.Update)

	// http.ListenAndServe("localhost:8888", nil)

	r := gin.Default()
	db.Init()

	r.Static("/assets", "./web/assets")

	// api routes
	apiRoutes := r.Group("/api")
	{
		apiRoutes.Use(middleware.ApiRequest())
		apiRoutes.POST("/activation", api.Activation, middleware.ApiResponse())
		apiRoutes.POST("/update", api.Update, middleware.ApiResponse())
		apiRoutes.POST("/download-updates", api.DownloadUpdates, middleware.ApiResponse())
		apiRoutes.POST("/upload-hosts", api.UploadHosts, middleware.ApiResponse())
		apiRoutes.POST("/system-status", api.SystemStatus, middleware.ApiResponse())
	}

	r.LoadHTMLGlob("web/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Bookie Guard Server",
		})
	})

	r.GET("/admin/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_index.html", gin.H{
			"title": "Bookie Guard Server",
		})
	})

	// r.GET("/users/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
	// 		"title": "Users",
	// 	})
	// })
	// website routes
	// r.GET("/", web.Index)

	// login
	r.GET("/login", web.Login)
	r.POST("/login", web.LoginPost)

	r.POST("/logout", web.Logout)

	// r.GET("/dashboard", web.Dashboard, middleware.AuthRequired())

	r.Run(net.JoinHostPort(config.SERVER_HOST, config.SERVER_PORT))

}
