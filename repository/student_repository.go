package repository

import (
	"database/sql"
	"fmt"

	"student-api/model"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

// CREATE
func (r *StudentRepository) Create(student model.Student) (model.Student, error) {
	query := `
		INSERT INTO students (name, age)
		VALUES ($1, $2)
		RETURNING id
	`

	err := r.db.QueryRow(query, student.Name, student.Age).
		Scan(&student.ID)
	if err != nil {
		return model.Student{}, err
	}

	return student, nil
}

// GET ALL
func (r *StudentRepository) GetAll() ([]model.Student, error) {
	query := `SELECT id, name, age FROM students`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var s model.Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Age); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

// UPDATE
func (r *StudentRepository) Update(student model.Student) error {
	query := `
		UPDATE students
		SET name = $1, age = $2
		WHERE id = $3
	`

	result, err := r.db.Exec(query, student.Name, student.Age, student.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no student found with id %d", student.ID)
	}

	return nil
}

// DELETE
func (r *StudentRepository) Delete(id int) error {
	query := `DELETE FROM students WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no student found with id %d", id)
	}

	return nil
}
