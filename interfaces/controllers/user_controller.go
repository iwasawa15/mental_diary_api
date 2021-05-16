package controllers

import (
	"fmt"
	"mental_diary_api/domain"
	"mental_diary_api/interfaces/database"
	"mental_diary_api/usecase"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	fmt.Println(u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, u)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, user)
}
