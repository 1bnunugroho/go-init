package auth

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	//"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, input CreateUserRequest) (string, error)
	Get(ctx context.Context, id string) (User, error)
	Query(ctx context.Context, offset, limit int) ([]User, error)
	Count(ctx context.Context) (int, error)
}

// User represents the data about an user.
type User struct {
	entity.User
}

// CreateUserRequest represents an user creation request.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Validate validates the CreateUserRequest fields.
func (m CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Username, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Password, validation.Required, validation.Length(0, 128)),
	)
}

// Validate validates the CreateUserRequest fields.
func (m CreateUserRequest) ValidateUsers() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Username, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Password, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetEmail returns the user email.
	GetEmail() string
	GetUserName() string
	GetBio() string
	GetImage() string
}

type service struct {
	signingKey      string
	tokenExpiration int
	repo            Repository
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration int, repo Repository, logger log.Logger) Service {
	return service{signingKey, tokenExpiration, repo, logger}
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, username, password string) (string, error) {
	if identity := s.authenticate(ctx, username, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

// Register register a user.
// Otherwise, an error is returned.
func (s service) Register(ctx context.Context, req CreateUserRequest) (string, error) {
	//func (s service) Register(ctx context.Context, username, password, email string) (string, error) {
	logger := s.logger.With(ctx, "req", req)
	if err := req.ValidateUsers(); err != nil {
		return "", err
	}

	logger.Infof("Validate successful")
	logger.Infof(req.Password)
	var hashedPassword, err = hashPassword(req.Password)
	logger.Infof(hashedPassword)

	if err != nil {
		logger.Infof("Error hashing password")
	}

	if err != nil {
		//return id, nil
		logger.Infof("Error create")
		return "", err
	}

	if identity := s.authenticate(ctx, req.Email, req.Password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

// Count returns the number of users.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the albums with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]User, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []User{}
	for _, item := range items {
		result = append(result, User{item})
	}
	return result, nil
}

// Get returns the user with the specified the user ID.
func (s service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}

// Hash password
func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// Check if two passwords match using Bcrypt's CompareHashAndPassword
// which return nil on success and an error on failure.
func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

// authenticate authenticates a user using username and password.
// If username and password are correct, an identity is returned. Otherwise, nil is returned.
func (s service) authenticate(ctx context.Context, email, password string) Identity {
	logger := s.logger.With(ctx, "user", email)

	// TODO: the following authentication logic is only for demo purpose
	if email == "demo@local.host" && password == "pass" {
		logger.Infof("authentication successful")
		//return entity.User{ID: "100", Name: "demo"}
		return entity.User{ID: "100", Email: "demo@local.host", Username: "rootz", Bio: "null", Image: "null"}
	}

	user, err := s.repo.GetUserName(ctx, email)
	if err != nil {
		logger.Infof("authentication failed ga dapet username", err)
		return nil
	}

	if doPasswordsMatch(user.Password, password) {
		logger.Infof(`dpt user user.Password , password`)
		logger.Infof("dpt user")
		return entity.User{ID: user.ID, Email: user.Email, Username: user.Username, Bio: user.Bio, Image: user.Image}
	}

	// if username == user.Username {
	// 	if comparePasswords(user.Password, getPwd(password)) {
	// 		logger.Infof("authentication successful")
	// 		return entity.User{ID: user.ID, Email: user.Email}
	// 	}
	// }

	logger.Infof("authentication failed")
	return nil
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       identity.GetID(),
		"email":    identity.GetEmail(),
		"username": identity.GetUserName(),
		"bio":      identity.GetBio(),
		"image":    identity.GetImage(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
