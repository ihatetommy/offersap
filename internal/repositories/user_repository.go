package repositories

import (
	"OffersApp/internal/entities"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user entities.User) (uuid.UUID, error)
	GetAllUsers() ([]entities.User, error)
	GetByID(id uuid.UUID) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user entities.User) error
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entities.User) (uuid.UUID, error) {
	query := `INSERT INTO users (email, password, created_at, updated_at) 
						VALUES ($1, $2, NOW(), NOW()) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(query, user.Email, user.Password).Scan(&id)
	if err != nil {
			return uuid.Nil, err
	}
	return id, nil
}

func (r *userRepository) GetAllUsers() ([]entities.User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
			return nil, err
	}
	defer rows.Close()
	users := []entities.User{}
	for rows.Next() {
			var user entities.User
			err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
			if err != nil {
					return nil, err
			}
			users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetByID(id uuid.UUID) (*entities.User, error) {
	query := `SELECT id, email, password, created_at, updated_at 
						FROM users WHERE id = $1`
	user := &entities.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
			if err == sql.ErrNoRows {
					return nil, errors.New("user not found")
			}
			return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(email string) (*entities.User, error) {
	query := `SELECT id, email, password, created_at, updated_at 
						FROM users WHERE email = $1`
	user := &entities.User{}
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
			if err == sql.ErrNoRows {
					return nil, errors.New("user not found")
			}
			return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user entities.User) error {
	query := `UPDATE users SET email = $1, password = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.db.Exec(query, user.Email, user.Password, user.ID)
	return err
}

func (r *userRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
