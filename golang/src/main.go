package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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

	e.Logger.Fatal(e.Start(":80"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World (Golang / Echo)")
}

func getUsers(c echo.Context) error {
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_DATABASE")
	mysqlEndpoint := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName
	db, err := otelsql.Open("mysql", mysqlEndpoint, otelsql.WithAttributes(semconv.DBSystemMySQL), otelsql.WithDBName(dbName))
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	rows, err := db.QueryContext(c.Request().Context(), "SELECT * FROM user")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		users = append(users, map[string]interface{}{
			"id":    id,
			"name":  name,
			"email": email,
		})
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
	errMessage := "Internal Server Error - Manual Exception"
	log.Printf(errMessage)
	return c.JSON(http.StatusInternalServerError, map[string]string{"timestamp": currentTime, "status": "500", "error": errMessage})
}
