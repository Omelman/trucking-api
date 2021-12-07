package service

import (
	"context"
	"errors"

	"github.com/Omelman/trucking-api/src/models"
)

// CreateUser creates new user.
func (s *Service) CreateUser(
	ctx context.Context,
	user *models.User,
) error {
	userCandidate, err := s.profileRepo.GetProfilesByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if userCandidate.ID != 0 {
		return errors.New("user already exists")
	}

	pwd, err := s.EncryptPassword(user.Password)
	user.Password = pwd

	user, err = s.profileRepo.AddNewProfile(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
