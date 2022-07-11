package profile

import (
	"context"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access profiles from the data source.
type Repository interface {
	// Get returns the profile with the specified profile ID.
	Get(ctx context.Context, id string) (entity.Profile, error)
	// Count returns the number of profiles.
	Count(ctx context.Context) (int, error)
	// Query returns the list of profiles with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Profile, error)
	// Create saves a new album in the storage.
	Create(ctx context.Context, profile entity.Profile) error
	// Update updates the album with given ID in the storage.
	Update(ctx context.Context, profile entity.Profile) error
	// Delete removes the album with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists albums in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the profile with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Profile, error) {
	var profile entity.Profile
	err := r.db.With(ctx).Select().Model(id, &profile)
	return profile, err
}

// Create saves a new profile record in the database.
// It returns the ID of the newly inserted profile record.
func (r repository) Create(ctx context.Context, profile entity.Profile) error {
	return r.db.With(ctx).Model(&profile).Insert()
}

// Update saves the changes to an profile in the database.
func (r repository) Update(ctx context.Context, profile entity.Profile) error {
	return r.db.With(ctx).Model(&profile).Update()
}

// Delete deletes an profile with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	profile, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&profile).Delete()
}

// Count returns the number of the profile records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("profile").Row(&count)
	return count, err
}

// Query retrieves the profile records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Profile, error) {
	var profiles []entity.Profile
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&profiles)
	return profiles, err
}
