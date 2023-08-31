package handlers

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lvmnpkfvmk/avito-tech/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDummyDB struct {
	mock.Mock
}

func (m *mockDummyDB) CreateSegment(segment *model.Segment) error {
	args := m.Called(segment)
	return args.Error(0)
}

func (m *mockDummyDB) DeleteSegment(segment *model.Segment) error {
	args := m.Called(segment)
	return args.Error(0)
}

var (
    // mockDB = map[string]*User{
    //     "jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
    // }
    userJSON = `{"name":"AVITO_VOICE_MESSAGES"}`
)

func TestCreateSegment(t *testing.T) {
    // Setup
    e := echo.New()
    req := httptest.NewRequest(http.MethodPost, "/segment", strings.NewReader(userJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

	mock := new(mockDummyDB)
	mock.On("CreateSegment", &model.Segment{Name: "AVITO_VOICE_MESSAGES"}).Return(nil)

	nr := &SegmentHandler{mock, &slog.Logger{}}
	r := "{\"data\":{\"name\":\"AVITO_VOICE_MESSAGES\"}}\n"

    // Assertions
    if assert.NoError(t, nr.CreateSegment(c)) {
        assert.Equal(t, http.StatusCreated, rec.Code)
        assert.Equal(t, string(r), rec.Body.String())
    }
}

func TestDeleteSegment(t *testing.T) {
    // Setup
    e := echo.New()
    req := httptest.NewRequest(http.MethodPost, "/segment", strings.NewReader(userJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

	mock := new(mockDummyDB)
	mock.On("DeleteSegment", &model.Segment{Name: "AVITO_VOICE_MESSAGES"}).Return(nil)

	nr := &SegmentHandler{mock, &slog.Logger{}}
	r := "{\"message\":\"Segment AVITO_VOICE_MESSAGES has been deleted\"}\n"

    // Assertions
    if assert.NoError(t, nr.DeleteSegment(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Equal(t, string(r), rec.Body.String())
    }
}
// func TestGetUser(t *testing.T) {
//     // Setup
//     e := echo.New()
//     req := httptest.NewRequest(http.MethodGet, "/", nil)
//     rec := httptest.NewRecorder()
//     c := e.NewContext(req, rec)
//     c.SetPath("/users/:email")
//     c.SetParamNames("email")
//     c.SetParamValues("jon@labstack.com")
//     h := &handler{mockDB}

//     // Assertions
//     if assert.NoError(t, h.getUser(c)) {
//         assert.Equal(t, http.StatusOK, rec.Code)
//         assert.Equal(t, userJSON, rec.Body.String())
//     }
// }