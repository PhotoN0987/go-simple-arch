package repository

import (
	"database/sql"
	"fmt"
	"go-simple-arch/pkg/database"
	"go-simple-arch/pkg/model"
	"log"
)

// UserRepository repository
type UserRepository interface {
	GetAll() ([]model.UserTable, error)
	GetByID(id int) (model.UserTable, error)
	Create(user model.User) error
	Update(user model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository is init for UserController
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{
		db: db.Connection,
	}
}

// GetAll Get all usersdata
func (r *userRepository) GetAll() ([]model.UserTable, error) {
	users := []model.UserTable{}

	query := `
		SELECT * FROM users
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user model.UserTable
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, err
}

// GetByID Get single usersdata
func (r *userRepository) GetByID(id int) (model.UserTable, error) {
	user := model.UserTable{}

	query := `
		SELECT * FROM users 
		WHERE id=?
	`
	row := r.db.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			return user, err
		}
	}

	return user, err
}

// Create Create user
func (r *userRepository) Create(user model.User) error {
	query := `
		INSERT INTO 
		users(name, email) 
		VALUES(?, ?)
	`
	stmtInsert, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(user.Name, user.Email)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(lastInsertID)

	return err
}

// Update Update user
func (r *userRepository) Update(user model.User) error {
	query := `
		UPDATE users 
		SET name=?, email=? 
		WHERE id=?
	`
	stmtUpdate, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()

	result, err := stmtUpdate.Exec(user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(rowsAffect)

	return err
}

// Delete Delete userdata
func (r *userRepository) Delete(id int) error {
	query := `
		DELETE 
		FROM users 
		WHERE id=?
	`
	stmtDelete, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(id)
	if err != nil {
		return err
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(rowsAffect)

	return err
}
