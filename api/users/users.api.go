package api_users

import (
	"demoapp/db"
	dbmodel "demoapp/db/models"
	"demoapp/helpers"
	apiModel "demoapp/models/api"
	"demoapp/schemas"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUsers(req *gin.Context) {
	defer func() {
		err := recover()

		if err != nil {
			req.JSON(http.StatusBadRequest, helpers.InternalServerError(helpers.SError{
				Error:   fmt.Errorf("%v", err),
				Message: "",
			}))
		}
	}()

	var users []dbmodel.User

	result := db.GetDb().Find(&users)

	if result.Error != nil {
		req.JSON(http.StatusBadRequest, helpers.InternalServerError(helpers.SError{
			Error:   result.Error,
			Message: "",
		}))
	}

	req.JSON(http.StatusOK, schemas.SSuccessResponse[any]{
		Success: true,
		Message: "",
		Data:    users,
	})
}

func getUser(req *gin.Context) {
	var params SGetUserParams
	var user dbmodel.User

	if err := req.ShouldBindUri(&params); err != nil {
		result := db.GetDb().First(&user, params.Id)

		if result.Error != nil {
			req.JSON(http.StatusBadRequest, helpers.InternalServerError(helpers.SError{
				Error:   result.Error,
				Message: "",
			}))
		}

		req.JSON(http.StatusOK, schemas.SSuccessResponse[any]{
			Success: true,
			Message: "",
			Data:    user,
		})
	}
}

func createUser(req *gin.Context) {
	var user SCreateUser

	if err := req.ShouldBindJSON(&user); err != nil {
		req.JSON(http.StatusBadRequest, helpers.BadRequest(helpers.SError{
			Error:   err,
			Message: "Invalid Payload",
		}))
	} else {
		result := db.GetDb().Create(&dbmodel.User{
			Name:  user.Name,
			Email: user.Email,
		})

		if result.Error != nil {
			req.JSON(http.StatusBadRequest, helpers.InternalServerError(helpers.SError{
				Error:   result.Error,
				Message: "Failed to create user",
			}))
		}

		if result.RowsAffected > 0 {
			req.JSON(http.StatusOK, schemas.SSuccessResponse[SCreateUser]{
				Success: true,
				Data:    user,
			})
		}
	}
}

func createTask(req *gin.Context) {
	defer func() {
		err := recover()

		if err != nil {
			log.Println(err)
			req.JSON(http.StatusInternalServerError, helpers.InternalServerError(helpers.SError{
				Error:   fmt.Errorf("%v", err),
				Message: "",
			}))
			return
		}
	}()

	var task SCreateTaskBody
	var params SCreateTaskParams
	var user dbmodel.User

	errParsingParams := req.ShouldBindUri(&params)

	if errParsingParams != nil {
		panic(errParsingParams)
	}

	if err := req.ShouldBindJSON(&task); err != nil {
		req.JSON(http.StatusBadRequest, helpers.BadRequest(helpers.SError{
			Error:   err,
			Message: "Invalid payload",
		}))
		return
	}

	result := db.GetDb().Preload("User").Find(&user, params.UserId)

	if result.RowsAffected < 1 {
		req.JSON(http.StatusBadRequest, helpers.BadRequest(helpers.SError{
			Error:   result.Error,
			Message: "User Not Found",
		}))
		return
	}

	createTaskResult := db.GetDb().Create(&dbmodel.Task{
		Name:        task.Name,
		Description: task.Description,
		UserId:      params.UserId,
	})

	if createTaskResult.Error != nil {
		panic(createTaskResult.Error)
	}

	if createTaskResult.RowsAffected > 0 {
		req.JSON(http.StatusAccepted, schemas.SSuccessResponse[any]{
			Success: true,
			Data:    nil,
			Message: "Task has been created",
		})
		return
	}
}

func getUserTasks(req *gin.Context) {
	var params SCreateTaskParams
	var tasks []dbmodel.Task

	req.BindUri(&params)

	result := db.GetDb().Where(&dbmodel.Task{UserId: params.UserId}).Find(&tasks)

	if result.Error != nil {
		panic(fmt.Sprintf("Error - %v", result.Error))
	}

	req.JSON(http.StatusOK, schemas.SSuccessResponse[[]dbmodel.Task]{
		Success: true,
		Data:    tasks,
	})
}

func Routes() apiModel.SRouteGroupDef {
	return apiModel.SRouteGroupDef{
		Path: "users",
		Routes: []apiModel.SRoute{
			{
				Route:      "/:userId/tasks",
				Method:     "GET",
				Controller: getUserTasks,
			},
			{
				Route:      "/",
				Method:     "GET",
				Controller: getAllUsers,
			},
			{
				Route:      "/create",
				Method:     "POST",
				Controller: createUser,
			},
			{
				Route:      "/create-task/:userId",
				Method:     "POST",
				Controller: createTask,
			},
			{
				Route:      "/:userId",
				Method:     "GET",
				Controller: getUser,
			},
		},
	}
}
