package controllers

import (
	"context"
	"dogovorsBackEnd/internal/use/dto"
	"dogovorsBackEnd/internal/use/entities"
	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/pkg/hash"
	"dogovorsBackEnd/pkg/randomizer"
)

func (c *controller) Auth(ctx context.Context, token string) (models.User, error) {
	user, err := c.repository.GetUserByToken(ctx, token)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		UserID:      user.UserID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Permissions: user.Permissions,
	}, nil
}

func (c *controller) SignUp(ctx context.Context, user models.User, request dto.SignUpRequestDTO) (dto.SignUpResponseDTO, error) {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return dto.SignUpResponseDTO{}, err
	}

	if err := c.validator.SignUpDTO(ctx, request); err != nil {
		return dto.SignUpResponseDTO{}, err
	}

	s, err := hash.Hash(request.Password)
	if err != nil {
		return dto.SignUpResponseDTO{}, err
	}

	u := entities.User{
		Username:  request.Username,
		Password:  s,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	id, err := c.repository.CreateUser(ctx, u)
	if err != nil {
		return dto.SignUpResponseDTO{}, err
	}

	u.UserID = id
	return dto.SignUpResponseDTO{
		UserID:    u.UserID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}, err
}

func (c *controller) Login(ctx context.Context, request dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	if err := c.validator.LoginDTO(ctx, request); err != nil {
		return dto.LoginResponseDTO{}, err
	}

	user, err := c.repository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	if !hash.CompareHashAndPassword(user.Password, request.Password) {
		return dto.LoginResponseDTO{}, ErrInvalidPassword
	}

	token, err := randomizer.RandomString(50)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	session := entities.Session{
		UserID: user.UserID,
		Token:  token,
	}

	err = c.repository.CreateSession(ctx, session)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	return dto.LoginResponseDTO{
		User: dto.LoginUserResponseDTO{
			UserID:    user.UserID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Token: session.Token,
	}, nil
}
