package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheTenzou/gis-diplom/user-service/apperrors"
	"github.com/TheTenzou/gis-diplom/user-service/mocks"
	"github.com/TheTenzou/gis-diplom/user-service/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUser(test *testing.T) {
	gin.SetMode(gin.TestMode)

	test.Run("Success", func(test *testing.T) {
		userID, _ := primitive.ObjectIDFromHex("607db993fad7324170a4debc")

		mockUserResponse := model.User{
			ID:    userID,
			Login: "alice",
			Name:  "Alice",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("FindByID", mock.AnythingOfType("*context.emptyCtx"), userID).Return(mockUserResponse, nil)

		recorder := httptest.NewRecorder()

		router := gin.Default()

		InitUserHandler(router, mockUserService)

		request, err := http.NewRequest(http.MethodGet, "/api/users/v1/user/"+userID.Hex(), nil)

		assert.NoError(test, err)

		router.ServeHTTP(recorder, request)

		responseBody, err := json.Marshal(mockUserResponse)

		assert.NoError(test, err)

		assert.Equal(test, 200, recorder.Code)
		assert.Equal(test, responseBody, recorder.Body.Bytes())
		mockUserService.AssertExpectations(test)
	})

	test.Run("NotFound", func(test *testing.T) {

		userID, _ := primitive.ObjectIDFromHex("607db993fad7324170a4debc")

		mockErr := apperrors.NewNotFound("user", userID.Hex())
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("FindByID", mock.Anything, userID).Return(model.User{}, mockErr)

		recorder := httptest.NewRecorder()

		router := gin.Default()

		InitUserHandler(router, mockUserService)

		request, err := http.NewRequest(http.MethodGet, "/api/users/v1/user/"+userID.Hex(), nil)
		assert.NoError(test, err)

		router.ServeHTTP(recorder, request)

		respBody, err := json.Marshal(mockErr)

		assert.NoError(test, err)

		assert.Equal(test, mockErr.StatusCode, recorder.Code)
		assert.Equal(test, respBody, recorder.Body.Bytes())
		mockUserService.AssertExpectations(test)
	})
}

func TestCreateUser(test *testing.T) {
	gin.SetMode(gin.TestMode)

	test.Run("Login and Password Required", func(test *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On(
			"Create",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("*models.User"),
		).Return(model.User{}, nil)

		recorder := httptest.NewRecorder()

		router := gin.Default()

		InitUserHandler(router, mockUserService)

		requestBody, err := json.Marshal(gin.H{
			"login": "",
		})

		assert.NoError(test, err)

		request, err := http.NewRequest(
			http.MethodPut,
			"/api/users/v1/user",
			bytes.NewBuffer(requestBody),
		)
		assert.NoError(test, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, request)

		assert.Equal(test, 400, recorder.Code)
		mockUserService.AssertNotCalled(test, "Create")
	})

	test.Run("Password to short", func(test *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On(
			"Create",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("*models.User"),
		).Return(model.User{}, nil)

		recorder := httptest.NewRecorder()

		router := gin.Default()

		InitUserHandler(router, mockUserService)

		requestBody, err := json.Marshal(gin.H{
			"login":    "alice",
			"password": "pass",
		})

		assert.NoError(test, err)

		request, err := http.NewRequest(
			http.MethodPut,
			"/api/users/v1/user",
			bytes.NewBuffer(requestBody),
		)
		assert.NoError(test, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, request)

		assert.Equal(test, 400, recorder.Code)
		mockUserService.AssertNotCalled(test, "Create")
	})

	test.Run("Password to long", func(test *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On(
			"Create",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("*models.User"),
		).Return(model.User{}, nil)

		recorder := httptest.NewRecorder()

		router := gin.Default()

		InitUserHandler(router, mockUserService)

		requestBody, err := json.Marshal(gin.H{
			"login":    "alice",
			"password": "passsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss",
		})

		assert.NoError(test, err)

		request, err := http.NewRequest(
			http.MethodPut,
			"/api/users/v1/user",
			bytes.NewBuffer(requestBody),
		)
		assert.NoError(test, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, request)

		assert.Equal(test, 400, recorder.Code)
		mockUserService.AssertNotCalled(test, "Create")
	})

	test.Run("Error returnd from userService", func(test *testing.T) {
		user := model.User{
			Login:    "alice1",
			Password: "alicePassword",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On(
			"Create",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("model.User"),
		).Return(user, apperrors.NewConflict("User Already Exits", user.Login))

		recorder := httptest.NewRecorder()

		router := gin.Default()

		InitUserHandler(router, mockUserService)

		requestBody, err := json.Marshal(gin.H{
			"login":    user.Login,
			"password": user.Password,
		})

		assert.NoError(test, err)

		request, err := http.NewRequest(
			http.MethodPut,
			"/api/users/v1/user",
			bytes.NewBuffer(requestBody),
		)
		assert.NoError(test, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, request)

		assert.Equal(test, 409, recorder.Code)
		mockUserService.AssertExpectations(test)
	})

	test.Run("Successful user creation", func(test *testing.T) {
		user := model.User{
			Login:    "alice1",
			Password: "alicePassword",
		}

		mockUserService := new(mocks.MockUserService)

		mockUserService.On(
			"Create",
			mock.AnythingOfType("*context.emptyCtx"),
			user,
		).Return(user, nil)

		recorder := httptest.NewRecorder()

		router := gin.Default()

		InitUserHandler(router, mockUserService)

		requestBody, err := json.Marshal(gin.H{
			"login":    user.Login,
			"password": user.Password,
		})
		assert.NoError(test, err)

		request, err := http.NewRequest(
			http.MethodPut,
			"/api/users/v1/user",
			bytes.NewBuffer(requestBody),
		)
		assert.NoError(test, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(recorder, request)

		responseBody, err := json.Marshal(user)
		assert.NoError(test, err)

		assert.Equal(test, http.StatusCreated, recorder.Code)
		assert.Equal(test, responseBody, recorder.Body.Bytes())

		mockUserService.AssertExpectations(test)
	})
}
