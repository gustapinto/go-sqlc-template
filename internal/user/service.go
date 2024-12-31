package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gustapinto/go-sqlc-template/internal/user/repository"

	"golang.org/x/crypto/bcrypt"
)

const (
	UserCreatedLogType = "USER_CREATED"
	UserUpdatedLogType = "USER_UPDATED"

	RepositoryTimeout = 30 * time.Second
)

type Service struct {
	querier repository.Queries
}

func NewService(db *sql.DB) *Service {
	return &Service{
		querier: *repository.New(db),
	}
}

func (s Service) GetByID(id int64) (repository.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RepositoryTimeout)
	defer cancel()

	return s.querier.SelectUserByID(ctx, id)
}

func (s Service) GetByLogin(login string) (repository.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RepositoryTimeout)
	defer cancel()

	return s.querier.SelectUserByLogin(ctx, login)
}

func (s Service) log(userID int64, logType, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), RepositoryTimeout)
	defer cancel()

	logArgs := repository.InsertUserLogParams{
		UserID:    userID,
		Type:      logType,
		Message:   message,
		CreatedAt: time.Now(),
	}
	_, err := s.querier.InsertUserLog(ctx, logArgs)
	return err
}

func (s Service) Create(login, password, email string) (repository.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return repository.User{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), RepositoryTimeout)
	defer cancel()

	args := repository.InsertUserParams{
		Login:     login,
		Password:  string(hashedPassword),
		Email:     email,
		CreatedAt: time.Now(),
	}
	user, err := s.querier.InsertUser(ctx, args)
	if err != nil {
		return repository.User{}, err
	}

	if err := s.log(user.ID, UserCreatedLogType, fmt.Sprintf("User %s created", user.Login)); err != nil {
		return repository.User{}, err
	}

	return user, nil
}

func (s Service) UpdateByID(userID int64, login, email string) (repository.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RepositoryTimeout)
	defer cancel()

	args := repository.UpdateUserByIDParams{
		Login:     login,
		Email:     email,
		UpdatedAt: time.Now(),
		ID:        userID,
	}
	if err := s.querier.UpdateUserByID(ctx, args); err != nil {
		return repository.User{}, err
	}

	return s.GetByID(userID)
}
