package main

import (
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/internal/post"
)

func main() {
	connectionString := ""
	conn, err := database.NewConnection(connectionString)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	repo := post.Repository{
		Conn: conn,
	}

	repo = repo
	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	g.Run(":3000")
}
