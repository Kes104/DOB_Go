package service

import (
	"context"

	"user-api/internal/models"
	"user-api/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

/* ---------------- CREATE ---------------- */

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob pgtype.Date,
) (*models.User, error) {

	u, err := s.repo.CreateUser(ctx, name, dob)
	if err != nil {
		return nil, err
	}

	age := CalculateAge(u.Dob.Time)

	return &models.User{
		ID:   u.ID, // int32
		Name: u.Name,
		Dob:  u.Dob.Time.Format("2006-01-02"),
		Age:  age,
	}, nil
}

/* ---------------- GET BY ID ---------------- */

func (s *UserService) GetUser(
	ctx context.Context,
	id int32,
) (*models.User, error) {

	u, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	age := CalculateAge(u.Dob.Time)

	return &models.User{
		ID:   u.ID, // int32
		Name: u.Name,
		Dob:  u.Dob.Time.Format("2006-01-02"),
		Age:  age,
	}, nil
}

/* ---------------- LIST ---------------- */

func (s *UserService) ListUsers(
	ctx context.Context,
) ([]models.User, error) {

	dbUsers, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0, len(dbUsers))

	for _, u := range dbUsers {
		age := CalculateAge(u.Dob.Time)

		users = append(users, models.User{
			ID:   u.ID, // int32
			Name: u.Name,
			Dob:  u.Dob.Time.Format("2006-01-02"),
			Age:  age,
		})
	}

	return users, nil
}

/* ---------------- UPDATE ---------------- */

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob pgtype.Date,
) (*models.User, error) {

	u, err := s.repo.UpdateUser(ctx, id, name, dob)
	if err != nil {
		return nil, err
	}

	age := CalculateAge(u.Dob.Time)

	return &models.User{
		ID:   u.ID, // int32
		Name: u.Name,
		Dob:  u.Dob.Time.Format("2006-01-02"),
		Age:  age,
	}, nil
}

/* ---------------- DELETE ---------------- */

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return s.repo.DeleteUser(ctx, id)
}



