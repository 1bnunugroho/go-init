package auth

import (
	"fmt"
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/dgrijalva/jwt-go"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"time"
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
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validates the CreateUserRequest fields.
func (m CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Username, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.Password, validation.Required, validation.Length(0, 128)),
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
	repo   			Repository
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
	logger := s.logger.With(ctx, "req", req)
	if err := req.Validate(); err != nil {
		return "", err
	}
	
	logger.Infof("Validate successful")

	id := entity.GenerateID()
	logger.Infof(id)

	now := time.Now()
	logger.Infof("time")
/*
	spass := getPwd(req.Password)
	logger.Infof("spass")

	Password := hashAndSalt(spass)
	logger.Infof("hashp")
*/

	err := s.repo.Create(ctx, entity.User{
		ID:        	id,
		Email:      req.Email,
		Username:	req.Username,
		//Password:	hashAndSalt(getPwd(req.Password)),
		Password:	req.Password,
		//Password:	Password,
		CreatedAt: 	now,
		UpdatedAt: 	now,
	})

	if err != nil {
		//return id, nil
		logger.Infof("Error create")
		return "", err	
	}

	//if err != nil {
		logger.Infof("harusnya sukses")
		return id, nil
	//}
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

func getPwd(pwd string) ([]byte) {    // Prompt the user to enter a password
    //fmt.Println("Enter a password")    // We will use this to store the users input
    //var pwd string    // Read the users input
    _, err := fmt.Scan(&pwd)
    if err != nil {
        //log.Println(err)
        //logger.Infof("authentication successful")
        //return err
    }    // Return the users input as a byte slice which will save us
    // from having to do this conversion later on
    return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {
    
    // Use GenerateFromPassword to hash & salt pwd
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        //log.Println(err)
        //return err
    }    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)    
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        //logger.Infof("authentication successful")
        //return err
        return false
    }
    
    return true
}

// authenticate authenticates a user using username and password.
// If username and password are correct, an identity is returned. Otherwise, nil is returned.
func (s service) authenticate(ctx context.Context, email, password string) Identity {
	logger := s.logger.With(ctx, "user", email)

	// TODO: the following authentication logic is only for demo purpose
	if email == "demo@local.host" && password == "pass" {
		logger.Infof("authentication successful")
		//return entity.User{ID: "100", Name: "demo"}
		return entity.User{ID: "100", Email: "demo@local.host", Username:"rootz", Bio:"null", Image:"null"}
	}

	user, err := s.repo.GetUserName(ctx, email)
	if err != nil {
		logger.Infof("authentication failed ga dapet username", err)
		return nil
	}

	if password == user.Password {
		logger.Infof("dpt user")
		return entity.User{ID: user.ID, Email: user.Email, Username:user.Username, Bio:user.Bio, Image:user.Image}
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
		"id":   identity.GetID(),
		"email": identity.GetEmail(),
		"username": identity.GetUserName(),
		"bio": identity.GetBio(),
		"image": identity.GetImage(),
		"iat": time.Now().Unix(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
