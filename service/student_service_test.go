package service

import (
	"testing"

	"student-api/model"
)

type mockRepo struct {
	CreateFn func(model.Student) (model.Student, error)
}

func (m *mockRepo) Create(s model.Student) (model.Student, error) {
	return m.CreateFn(s)
}
func (m *mockRepo) GetAll() ([]model.Student, error) { return nil, nil }
func (m *mockRepo) Update(model.Student) error       { return nil }
func (m *mockRepo) Delete(int) error                 { return nil }

func TestCreateStudent(t *testing.T) {
	repo := &mockRepo{
		CreateFn: func(s model.Student) (model.Student, error) {
			s.ID = 1
			return s, nil
		},
	}

	service := NewStudentService(repo)

	student := model.Student{Name: "Ajith", Age: 29}
	result, err := service.Create(student)

	if err != nil {
		t.Fatal(err)
	}

	if result.ID != 1 {
		t.Fatalf("expected id 1, got %d", result.ID)
	}
}
