package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/ed/gaintime/dto"
	"github.com/ed/gaintime/entity"
	"github.com/ed/gaintime/helper"
	"github.com/ed/gaintime/service"
	"github.com/gin-gonic/gin"
)

type TodoController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type todoController struct {
	todoService service.TodoService
	jwtService  service.JWTService
}

func NewTodoController(todoServ service.TodoService, jwtServ service.JWTService) TodoController {
	return &todoController{
		todoService: todoServ,
		jwtService:  jwtServ,
	}
}

func (c *todoController) All(context *gin.Context) {
	var todos []entity.Todo = c.todoService.All()
	res := helper.BuildResponse(true, "OK!", todos)
	context.JSON(http.StatusOK, res)
}

func (c *todoController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param is was found", err.Error(), helper.EmtyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var todo entity.Todo = c.todoService.FindByID(id)
	if (todo == entity.Todo{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmtyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK!", todo)
		context.JSON(http.StatusOK, res)
	}
}

func (c *todoController) Insert(context *gin.Context) {
	var todoCreateDTO dto.TodoCreateDTO
	errDTO := context.ShouldBind(&todoCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmtyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		ConvertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			todoCreateDTO.UserID = ConvertedUserID
		}
		result := c.todoService.Insert(todoCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *todoController) Update(context *gin.Context) {
	var todoUpdateDTO dto.TodoUpdateDTO
	errDTO := context.ShouldBind(&todoUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmtyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	//驗證是否帶有token,有則回傳id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claim := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claim["user_id"])

	//驗證ID並允許更新資料
	if c.todoService.IsAllowedToEdit(userID, todoUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			todoUpdateDTO.UserID = id
		}
		result := c.todoService.Update(todoUpdateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You don't have premission", "You are not the owner", helper.EmtyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *todoController) Delete(context *gin.Context) {
	var todo entity.Todo
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmtyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	todo.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.todoService.IsAllowedToEdit(userID, todo.ID) {
		c.todoService.Delete(todo)
		res := helper.BuildResponse(true, "Deleted!", helper.EmtyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmtyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *todoController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
