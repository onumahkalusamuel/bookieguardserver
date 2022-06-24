package main

import (
	"log"
	"net"
	"os"

	"bookieguardserver/config"
	"bookieguardserver/internal/account"
	"bookieguardserver/internal/admin"
	"bookieguardserver/internal/api"
	"bookieguardserver/internal/db"
	"bookieguardserver/internal/middleware"
	"bookieguardserver/internal/public"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// prepare all variables
	config.SetUpEnv()

	// get gin
	r := gin.Default()

	// set session
	store := cookie.NewStore([]byte("bookieguard"))
	r.Use(sessions.Sessions("bookieguard", store))

	// setup the database
	db.Init()

	// set route for static files
	r.Static("/assets", "./web/assets")

	// set template for all web pages
	r.LoadHTMLGlob("./web/**/*.html")

	// api routes
	apiRoutes := r.Group("/api/")
	{
		apiRoutes.Use(middleware.ApiRequest())
		apiRoutes.POST("activation", api.Activation, middleware.ApiResponse())
		apiRoutes.POST("update", api.Update, middleware.ApiResponse())
		apiRoutes.POST("download-updates", api.DownloadUpdates, middleware.ApiResponse())
		apiRoutes.POST("upload-hosts", api.UploadHosts, middleware.ApiResponse())
		apiRoutes.POST("system-status", api.SystemStatus, middleware.ApiResponse())
	}

	// public routes
	publicRoutes := r.Group("/")
	{
		publicRoutes.GET("/", public.Index)
		publicRoutes.GET("/how-it-works", public.HowItWorks)
		publicRoutes.GET("/contact", public.Contact)
		publicRoutes.GET("/pricing", public.Pricing)

	}

	// admin routes
	adminRoutes := r.Group("/admin/")
	{
		adminRoutes.GET("/login", admin.Login)
		adminRoutes.POST("/login", admin.Login)

		adminRoutes.Use(middleware.AdminAuth())
		adminRoutes.GET("/", admin.Dashboard)
		adminRoutes.GET("/dashboard", admin.Dashboard)
		adminRoutes.GET("/logout", admin.Logout)

		adminRoutes.GET("/users", admin.Users)
		adminRoutes.POST("/users", admin.Users)
		adminRoutes.GET("/users/:user_id", admin.User)
		adminRoutes.GET("/users/:user_id/delete", admin.UserDelete)
		adminRoutes.GET("/users/:user_id/block-groups", admin.UserBlockGroups)
		adminRoutes.POST("/users/:user_id/block-groups", admin.UserBlockGroups)
		adminRoutes.GET("/users/:user_id/block-groups/:blockgroup_id/settings", admin.UserBlockGroupSettings)
		adminRoutes.POST("/users/:user_id/block-groups/:blockgroup_id/settings", admin.UserBlockGroupSettings)
		adminRoutes.GET("/users/:user_id/block-groups/:blockgroup_id/settings/:action/:action_id", admin.UserBlockGroupSettingsAction)

		adminRoutes.GET("/blocklist-categories", admin.BlocklistCategories)
		adminRoutes.POST("/blocklist-categories", admin.BlocklistCategories)
		adminRoutes.GET("/blocklist-categories/:category_id/delete", admin.BlocklistCategoriesAction)

		adminRoutes.GET("/blocklists", admin.Blocklists)
		adminRoutes.POST("/blocklists", admin.Blocklists)
		adminRoutes.GET("/blocklists/:blocklist_id/delete", admin.BlocklistAction)

		adminRoutes.GET("/hosts", admin.Hosts)
		adminRoutes.GET("/hosts/:host_id/:action", admin.HostAction)

		adminRoutes.GET("/settings", admin.Settings)
		adminRoutes.POST("/settings", admin.Settings)

	}

	// user routes
	accountRoutes := r.Group("/account")
	{
		accountRoutes.GET("/login", account.Login)
		accountRoutes.POST("/login", account.LoginPost)
		accountRoutes.GET("/register", account.Register)
		accountRoutes.POST("/register", account.RegisterPost)
		accountRoutes.GET("/forgot-password", account.ForgotPassword)
		accountRoutes.POST("/forgot-password", account.ForgotPassword)
		accountRoutes.GET("/paystack-callback", account.PaystackCallBack)
		// authenticated routes
		accountRoutes.Use(middleware.AccountAuth())
		accountRoutes.GET("/", account.Dashboard)
		accountRoutes.GET("/dashboard", account.Dashboard)
		accountRoutes.GET("/logout", account.Logout)
		// block groups
		accountRoutes.GET("/block-groups", account.BlockGroups)
		accountRoutes.POST("/block-groups", account.BlockGroups)
		accountRoutes.GET("/block-groups/:blockgroup_id", account.BlockGroup)
		accountRoutes.GET("/block-groups/:blockgroup_id/settings", account.BlockGroupSettings)
		accountRoutes.POST("/block-groups/:blockgroup_id/settings", account.BlockGroupSettings)
		accountRoutes.GET("/block-groups/:blockgroup_id/settings/:action/:action_id", account.BlockGroupSettingsAction)
		// settings
		accountRoutes.GET("/settings", account.Settings)
		accountRoutes.POST("/settings", account.Settings)
	}

	// start server
	r.Run(net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT")))

}
