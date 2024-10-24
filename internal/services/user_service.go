package services

import (
	"OffersApp/internal/auth"
	"OffersApp/internal/entities"
	"OffersApp/internal/repositories"
	"OffersApp/internal/utils"
	"errors"

	"github.com/google/uuid"
)

// UserService defines the methods that any user service should implement.
type UserService interface {
	Register(user entities.User) (uuid.UUID, error)
	GetUserByID(id uuid.UUID) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(id uuid.UUID) error
	GetAllUsers() ([]entities.User, error)
	Authenticate(email string, password string) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(user entities.User) (uuid.UUID, error) {
	if user.Email == "" {
		return uuid.Nil, errors.New("email cannot be empty")
	}

	existingUser, _ := s.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return uuid.Nil, errors.New("email already exists")
	}

	if user.Password == "" {
		return uuid.Nil, errors.New("password cannot be empty")
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	return s.userRepo.Create(user)
}

func (s *userService) GetUserByID(id uuid.UUID) (*entities.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) GetUserByEmail(email string) (*entities.User, error) {
	return s.userRepo.GetByEmail(email)
}

func (s *userService) UpdateUser(user entities.User) error {
	if user.ID == uuid.Nil {
		return errors.New("user ID cannot be empty")
	}
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.userRepo.Delete(id)
}

func (s *userService) GetAllUsers() ([]entities.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) Authenticate(email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if validPassword := utils.CheckPasswordHash(password, user.Password); !validPassword {
		return "", errors.New("invalid email or password")
	}

	// Generate a JWT token
	token, err := auth.GenerateJWT(user.Email)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
