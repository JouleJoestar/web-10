package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	uc      Usecase
	address string
}

func NewServer(ip string, port int, uc Usecase) *Server {
	return &Server{
		uc:      uc,
		address: fmt.Sprintf("%s:%d", ip, port),
	}
}

func (srv *Server) GetUser(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name parameter is required"})
	}

	user, err := srv.uc.FetchUser(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.String(http.StatusOK, "Hello, "+user+"!")
}

func (srv *Server) PostUser(c echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := srv.uc.CreateUser(input.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Запись добавлена!"})
}

func (srv *Server) Run() error {
	e := echo.New()

	e.GET("/api/user", srv.GetUser)
	e.POST("/api/user/create", srv.PostUser)

	return e.Start(srv.address)
}
