package service

import "student-api/model"

// StudentService defines business operations
type StudentService interface {
	Create(student model.Student) (model.Student, error)
	GetAll() ([]model.Student, error)
	Update(student model.Student) error
	Delete(id int) error
}

type studentService struct {
	repo StudentRepository
}

// StudentRepository is the dependency required by service
// (implemented by repository layer)
type StudentRepository interface {
	Create(model.Student) (model.Student, error)
	GetAll() ([]model.Student, error)
	Update(model.Student) error
	Delete(int) error
}

// NewStudentService creates a new service instance
func NewStudentService(repo StudentRepository) StudentService {
	return &studentService{
		repo: repo,
	}
}

func (s *studentService) Create(student model.Student) (model.Student, error) {
	return s.repo.Create(student)
}

func (s *studentService) GetAll() ([]model.Student, error) {
	return s.repo.GetAll()
}

func (s *studentService) Update(student model.Student) error {
	return s.repo.Update(student)
}

func (s *studentService) Delete(id int) error {
	return s.repo.Delete(id)
}
