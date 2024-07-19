package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/user", func(c echo.Context) error {
		id := c.QueryParam("id")

		return c.String(http.StatusOK, id)
	})
	e.GET("/user/heatmap", func(c echo.Context) error {
		// This should give the days when the user has played like the github contribution graph
		id := c.QueryParam("id")

		response := "Hello" + id
		return c.String(http.StatusOK, response)
	})
	e.GET("/user/trend/trophies", func(c echo.Context) error {
		// This should give the trophies from the user since the first sign up on the site. To make a graph of the trophies
		// id := c.QueryParam("id")

		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start(":8123"))
}

// func database() {
// 	db, err := sql.Open("sqlite3", "./test.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// }
