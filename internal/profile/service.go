package profile

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"time"
)

// Service encapsulates usecase logic for profiles.
type Service interface {
	Get(ctx context.Context, id string) (Profile, error)
	Query(ctx context.Context, offset, limit int) ([]Profile, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateProfileRequest) (Profile, error)
	Update(ctx context.Context, id string, input UpdateProfileRequest) (Profile, error)
	Delete(ctx context.Context, id string) (Profile, error)
}

// Album represents the data about an profile.
type Profile struct {
	entity.Profile
}

// CreateAlbumRequest represents an profile creation request.
type CreateProfileRequest struct {
	Bio string `json:"bio"`
}

// Validate validates the CreateAlbumRequest fields.
func (m CreateProfileRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Bio, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateProfileRequest represents an profile update request.
type UpdateProfileRequest struct {
	Bio string `json:"bio"`
}

// Validate validates the CreateAlbumRequest fields.
func (m UpdateProfileRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Bio, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new profile service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the profile with the specified the profile ID.
func (s service) Get(ctx context.Context, id string) (Profile, error) {
	profile, err := s.repo.Get(ctx, id)
	if err != nil {
		return Profile{}, err
	}
	return Profile{profile}, nil
}

// Create creates a new profile.
func (s service) Create(ctx context.Context, req CreateProfileRequest) (Profile, error) {
	if err := req.Validate(); err != nil {
		return Profile{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Profile{
		ID:        id,
		Bio:      req.Bio,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return Profile{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the profile with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateProfileRequest) (Profile, error) {
	if err := req.Validate(); err != nil {
		return Profile{}, err
	}

	profile, err := s.Get(ctx, id)
	if err != nil {
		return profile, err
	}
	profile.Bio = req.Bio
	profile.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, profile.Profile); err != nil {
		return profile, err
	}
	return profile, nil
}

// Delete deletes the profile with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Profile, error) {
	profile, err := s.Get(ctx, id)
	if err != nil {
		return Profile{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Profile{}, err
	}
	return profile, nil
}

// Count returns the number of profiles.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the profiles with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Profile, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Profile{}
	for _, item := range items {
		result = append(result, Profile{item})
	}
	return result, nil
}