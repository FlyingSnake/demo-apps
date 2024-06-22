package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/users", getUsers)
	e.GET("/sleep/:seconds", sleep)
	e.GET("/status/random", randomStatus)
	e.GET("/exception", exceptionHandler)

	e.Start(":80")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World (Golang / Echo)")
}

func getUsers(c echo.Context) error {
	db, err := sql.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_DATABASE"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection error"})
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Query execution error"})
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Row scanning error"})
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func sleep(c echo.Context) error {
	secondsStr := c.Param("seconds")
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid number of seconds"})
	}

	time.Sleep(time.Duration(seconds) * time.Second)
	return c.String(http.StatusOK, "sleep "+secondsStr+"s")
}

func randomStatus(c echo.Context) error {
	rand.Seed(time.Now().UnixNano())
	statuses := []struct {
		code    int
		message string
	}{
		{http.StatusOK, "OK"},
		{http.StatusCreated, "Created"},
		{http.StatusAccepted, "Accepted"},
		{http.StatusNoContent, "No Content"},
		{http.StatusBadRequest, "Bad Request"},
		{http.StatusUnauthorized, "Unauthorized"},
		{http.StatusForbidden, "Forbidden"},
		{http.StatusNotFound, "Not Found"},
		{http.StatusInternalServerError, "Internal Server Error"},
		{http.StatusNotImplemented, "Not Implemented"},
		{http.StatusBadGateway, "Bad Gateway"},
		{http.StatusServiceUnavailable, "Service Unavailable"},
	}

	randomStatus := statuses[rand.Intn(len(statuses))]
	return c.String(randomStatus.code, randomStatus.message)
}

func exceptionHandler(c echo.Context) error {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	errMessage := "exception called at " + currentTime
	log.Printf("[ERROR] %s", errMessage)
	panic(errMessage)
}
