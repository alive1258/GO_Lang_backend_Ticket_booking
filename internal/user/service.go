package user

import "goticket/internal/user/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(
	req dto.CreateUserRequest,
) (*dto.CreateUserResponse, error) {

	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	response := dto.CreateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}