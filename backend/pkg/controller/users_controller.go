package controller

import (
	"github.com/SzymekN/Car-rental-app/pkg/model"

	"github.com/labstack/echo/v4"
)

// swagger:route POST /api/v1/users/save users_v1 postUserV1
// Save user to  database.
//
//		Consumes:
//	   - application/json
//	 Produces:
//	   - application/json
//
// responses:
//
//	200: userResponse
//	500: errorResponse

type UsersController struct {
	MainController
}

type UsersHandler interface {
	SaveUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetUserById(c echo.Context) error
	GetUsers(c echo.Context) error
}

func (uc *UsersController) SaveUser(c echo.Context) error {
	return GenericPost(c, model.User{})
}

func (uc *UsersController) UpdateUser(c echo.Context) error {
	return GenericUpdate(c, model.User{})
}

func (uc *UsersController) DeleteUser(c echo.Context) error {
	return GenericDelete(c, model.User{})
}

func (uc *UsersController) GetUserById(c echo.Context) error {
	return GenericGetById(c, model.User{})
}

func (uc *UsersController) GetUsers(c echo.Context) error {
	return GenericGetAll(c, []model.User{})
}
