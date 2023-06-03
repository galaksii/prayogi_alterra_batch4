package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id           int `gorm:"primary_key autoIncrement"`
	Name         string
	ProvinceCase string
}

type Cryptocurrency struct {
	ID    int    `json:"id"`
	Coin  string `json:"coin"`
	Harga string `json:"harga"`
}

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AppController struct{}

func (c *AppController) GetDetailCoinController(ctx echo.Context) error {
	id := ctx.Param("id")
	cryptocurrency := Cryptocurrency{Coin: id, Harga: " "}
	return ctx.JSON(http.StatusOK, cryptocurrency)
}

func (c *AppController) LoginController(ctx echo.Context) error {
	var loginRequest LoginRequest
	if err := ctx.Bind(&loginRequest); err != nil {
		errResponse := ErrorResponse{Message: "Requestnya salah,mohon ulangi"}
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	return ctx.JSON(http.StatusOK, loginRequest)
}

func GetUserController(c echo.Context) error {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, users)
}

func (c *AppController) GetCoinController(ctx echo.Context) error {
	var peopleData []Cryptocurrency

	var people Cryptocurrency
	people.Coin = "Bitcoin"
	people.Harga = "$27.192"
	people.ID = 1
	peopleData = append(peopleData, people)

	people.Coin = "Ethereum"
	people.Harga = "$1.903"
	people.ID = 2
	peopleData = append(peopleData, people)

	people.Coin = "Tether"
	people.Harga = "$1"
	people.ID = 3
	peopleData = append(peopleData, people)

	people.Coin = "BNB"
	people.Harga = "$307"
	people.ID = 4
	peopleData = append(peopleData, people)

	people.Coin = "USDC"
	people.Harga = "$0.99"
	people.ID = 5
	peopleData = append(peopleData, people)

	people.Coin = "XRP"
	people.Harga = "$0.5205"
	people.ID = 6
	peopleData = append(peopleData, people)

	people.Coin = "Cardano"
	people.Harga = "$0.3764"
	people.ID = 7
	peopleData = append(peopleData, people)

	people.Coin = "DogeCoin"
	people.Harga = "$0.07264"
	people.ID = 8
	peopleData = append(peopleData, people)

	people.Coin = "Solana"
	people.Harga = "$21.16"
	people.ID = 9
	peopleData = append(peopleData, people)

	people.Coin = "Polygon"
	people.Harga = "$0.9039"
	people.ID = 10
	peopleData = append(peopleData, people)

	var response BaseResponse
	response.Data = peopleData
	response.Message = "Berhasil"

	return ctx.JSON(http.StatusOK, response)
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	connectDatabase()
	e := echo.New()
	controller := &AppController{}

	// Routing
	e.GET("/users", GetUserController) //ini untuk database salahs atu beritanya tapi salah nama lal hehe
	e.GET("/coin", controller.GetCoinController)
	e.GET("/coin/:id", controller.GetDetailCoinController)
	e.POST("/login", controller.LoginController)

	// Start server
	e.Start(":8000")
}

func connectDatabase() {
	dsn := "root:itech227@tcp(127.0.0.1:3306)/peringkat_kasus?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
	fmt.Println("Database")
}
