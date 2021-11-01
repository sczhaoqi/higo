package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	ID   uint64 `form:"id" binding:"required,gt=0"`
	Name string `form:"name" binding:"required"`
}

func main() {
	hello()
	r := gin.Default()
	users := map[uint64]User{1: {ID: 1, Name: "张三"}, 2: {ID: 2, Name: "李四"}}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello", hello())
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		//var user User
		//
		//if err := c.ShouldBind(&user); err != nil {
		//	fmt.Printf(fmt.Sprint(err))
		//	return
		//}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, "id not exists.")
			return
		}
		user, ok := users[id]
		if ok {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusNotFound, "not found")
		}

	})
	r.Run(":8080")
}
