package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"student-api/model"
)

type mockService struct {
	CreateFn func(model.Student) (model.Student, error)
}

func (m *mockService) Create(s model.Student) (model.Student, error) {
	return m.CreateFn(s)
}
func (m *mockService) GetAll() ([]model.Student, error) { return nil, nil }
func (m *mockService) Update(model.Student) error       { return nil }
func (m *mockService) Delete(int) error                 { return nil }

func TestCreateStudent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := &mockService{
		CreateFn: func(s model.Student) (model.Student, error) {
			s.ID = 1
			return s, nil
		},
	}

	handler := NewStudentHandler(service)

	router := gin.New()
	router.POST("/students", handler.CreateStudent)

	req := httptest.NewRequest(
		"POST",
		"/students",
		strings.NewReader(`{"name":"Ajith","age":29}`),
	)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}
