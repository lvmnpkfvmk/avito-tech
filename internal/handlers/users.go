package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	algo "github.com/lvmnpkfvmk/avito-tech/internal/lib"
	"github.com/lvmnpkfvmk/avito-tech/internal/model"
	"github.com/lvmnpkfvmk/avito-tech/internal/repository"
)
type UpdateUserRequest struct {
	ID uint `json:"id,string"`
	SegmentsToAdd model.Segments `json:"segments_to_add"`
	SegmentsToDelete model.Segments `json:"segments_to_delete"`
}

type GetUserRequest struct {
	ID uint `json:"id,string"`
}

type UserHandler struct {
	repo repository.ISegmentRepository
	logger *slog.Logger
}

func NewUserHandler(repo repository.ISegmentRepository, logger *slog.Logger) *UserHandler {
	return &UserHandler{repo, logger}
}

func (sh *UserHandler) UpdateUser(c echo.Context) error {
	b := new(UpdateUserRequest)

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}
	fmt.Println(b)

	user, err := sh.repo.GetUser(b.ID)

	fmt.Println(user)

	if err != nil && user == nil { // user does not exist
		user = &model.User{}
		user.Segments = &b.SegmentsToAdd
	} else {
		user.ID = b.ID
		x := append(b.SegmentsToAdd, *user.Segments...)
		user.Segments = algo.FilterSegments(&x, &b.SegmentsToDelete)
	}

	if err := sh.repo.UpdateUser(user); err != nil {
		data := map[string]interface{}{
			"message": err,
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": user,
	}

	return c.JSON(http.StatusCreated, response)
}

func (sh *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := sh.repo.GetAllUsers()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, users)
}

func (sh *UserHandler) GetUser(c echo.Context) error {
	b := new(GetUserRequest)

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	user, err := sh.repo.GetUser(b.ID)
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, user)
}