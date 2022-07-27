package example

// func setupLogOutPut() {
// 	f, _ := os.Create("gin.log")
// 	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
// }

// func main() {
// 	setupLogOutPut()
// 	server := gin.New()
// 	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.MyMiddleware)
// 	server.Static("static", "./templates/static")
// 	server.LoadHTMLGlob("templates/index.html")
// 	router.Router(server)
// 	server.Run("127.0.0.1:9091")
// }

// indexRoutes := server.Group("/")
// {
// 	indexRoutes.GET("", VideoController.ShowAll)
// }

// server.GET("/videos", func(ctx *gin.Context) {
// 	ctx.JSON(200, VideoController.FindAll())
// })

// server.POST("/videos", func(ctx *gin.Context) {
// 	err := VideoController.Save(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	} else {
// 		ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
// 	}
// })
