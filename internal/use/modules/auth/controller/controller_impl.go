package controller

import (
	"context"
	"fmt"

	"dogovorsBackEnd/internal/use/modules/auth/dto"

	"dogovorsBackEnd/internal/use/models"
	"dogovorsBackEnd/internal/use/modules/auth/entities"
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

	if err := c.validator.SignUpRequest(ctx, request); err != nil {
		return dto.SignUpResponseDTO{}, err
	}

	s, err := hash.Hash(request.Password)
	if err != nil {
		return dto.SignUpResponseDTO{}, err
	}

	u := entities.User{
		Username:    request.Username,
		Password:    s,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Permissions: request.Permissions,
	}

	id, err := c.repository.CreateUser(ctx, u)
	if err != nil {
		return dto.SignUpResponseDTO{}, err
	}

	u.UserID = id
	return dto.SignUpResponseDTO{
		UserID:      u.UserID,
		Username:    u.Username,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Permissions: u.Permissions,
	}, err
}

func (c *controller) Login(ctx context.Context, request dto.LoginRequestDTO) (dto.LoginResponseDTO, error) {
	if err := c.validator.LoginRequest(ctx, request); err != nil {
		return dto.LoginResponseDTO{}, err
	}

	user, err := c.repository.GetUserByUsername(ctx, request.Username)
	if err != nil {
		return dto.LoginResponseDTO{}, ErrInvalidUsername
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
			UserID:      user.UserID,
			Username:    user.Username,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Permissions: user.Permissions,
		},
		Token: session.Token,
	}, nil
}

func (c *controller) PatchUser(ctx context.Context, user models.User, request dto.UserUpdateRequestDTO) (dto.UserUpdateResponseDTO, error) {
	if err := c.validator.UserUpdateRequest(ctx, request); err != nil {
		return dto.UserUpdateResponseDTO{}, err
	}

	u := entities.User{
		UserID: request.UserID,
	}

	if request.UserID != user.UserID {
		if err := user.AssertPermission(models.PermissionAdmin); err != nil {
			return dto.UserUpdateResponseDTO{}, fmt.Errorf("%w: %w", err, ErrForbidden)
		}
	}

	if request.Password != "" {
		if user.UserID != request.UserID {
			return dto.UserUpdateResponseDTO{}, fmt.Errorf("cannot change password: %w", ErrForbidden)
		}

		hashedPassword, err := hash.Hash(request.Password)
		if err != nil {
			return dto.UserUpdateResponseDTO{}, err
		}

		u.Password = hashedPassword
	}

	if request.FirstName != "" {
		u.FirstName = request.FirstName
	}

	if request.LastName != "" {
		u.LastName = request.LastName
	}

	if request.Permissions != nil {
		u.Permissions = request.Permissions
	}

	err := c.repository.UpdateUser(ctx, u)
	if err != nil {
		return dto.UserUpdateResponseDTO{}, err
	}

	return dto.UserUpdateResponseDTO{
		UserID:      u.UserID,
		Username:    u.Username,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Permissions: u.Permissions,
	}, nil
}

func (c *controller) DeleteUserByID(ctx context.Context, user models.User, userID int) error {
	if err := user.AssertPermission(models.PermissionAdmin); err != nil {
		return fmt.Errorf("%w: %w", err, ErrForbidden)
	}

	if user.UserID == userID {
		return fmt.Errorf("cannot delete yourself: %w", ErrForbidden)
	}

	return c.repository.DeleteUserByID(ctx, userID)
}
