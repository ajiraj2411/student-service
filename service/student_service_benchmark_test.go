package service

import (
	"testing"

	"student-api/model"
)

type benchRepo struct{}

func (b *benchRepo) Create(s model.Student) (model.Student, error) {
	s.ID = 1
	return s, nil
}
func (b *benchRepo) GetAll() ([]model.Student, error) { return nil, nil }
func (b *benchRepo) Update(model.Student) error       { return nil }
func (b *benchRepo) Delete(int) error                 { return nil }

func BenchmarkStudentService_Create(b *testing.B) {
	repo := &benchRepo{}
	service := NewStudentService(repo)

	student := model.Student{Name: "Ajith", Age: 29}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.Create(student)
	}
}
