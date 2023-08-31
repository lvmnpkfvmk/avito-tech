package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lvmnpkfvmk/avito-tech/internal/model"
	"github.com/lvmnpkfvmk/avito-tech/internal/repository"
)

type CreateSegmentRequest struct {
	Name string `json:"name"`
}
type DeleteSegmentRequest struct {
	Name string `json:"name"`
}

type SegmentHandler struct {
	repo repository.ISegmentRepository
	logger *slog.Logger
}

func NewSegmentHandler(repo repository.ISegmentRepository, logger *slog.Logger) *SegmentHandler {
	return &SegmentHandler{repo, logger}
}

func (sh *SegmentHandler) CreateSegment(c echo.Context) error {
	b := new(CreateSegmentRequest)

	// Binding data
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	fmt.Println(b)

	nb := &model.Segment{Name: b.Name}
	if err := sh.repo.CreateSegment(nb); err != nil {
		data := map[string]interface{}{
			"message": err,
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": b,
	}

	return c.JSON(http.StatusCreated, response)
}

func (sh *SegmentHandler) DeleteSegment(c echo.Context) error {
	b := new(DeleteSegmentRequest)

	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	seg := model.Segment{Name: b.Name}

	err := sh.repo.DeleteSegment(&seg)
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": fmt.Sprintf("Segment %s has been deleted", seg.Name),
	}
	return c.JSON(http.StatusOK, response)
}

func (sh *SegmentHandler) GetAllSegments(c echo.Context) error {

	segments, err := sh.repo.GetAllSegments()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, segments)
}