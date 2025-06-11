// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

//	func main() {
//		r := gin.Default()
//		r.GET("/ping", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"message": "pong ping",
//			})
//		})
//		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//	}

// package main

// import (
// 	databaseConn "Go_Watchlist/dbConn"
// 	"Go_Watchlist/handler"

// 	"github.com/gin-gonic/gin"
// )

//	func main() {
//		databaseConn.Init()
//		r := gin.Default()
//		r.POST("/register", handler.RegisterUserHandler)
//		r.Run(":8080")
//	}
package main

import (
	databaseConn "Go_Watchlist/dbConn"
	"Go_Watchlist/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	databaseConn.Init()
	r := gin.Default()

	r.POST("/register/add", handler.RegisterUserHandler)
	r.POST("/register/remove", handler.RemoveRegisterUser)

	// Watchlist APIs
	r.POST("/watchlist/add", handler.AddToWatchlistHandler)
	r.POST("/watchlist/remove", handler.RemoveFromWatchlistHandler)
	// r.GET("/watchlist", handler.GetWatchlistHandler)

	r.Run(":8080")
}
