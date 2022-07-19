package auth

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(rg *routing.RouteGroup, service Service, logger log.Logger) {
	rg.Post("/login", login(service, logger))
	rg.Post("/users/login", login(service, logger))
	rg.Post("/register", register(service, logger))
	rg.Get("/users", query(service, logger))
}

// login returns a handler that handles user login request.
func login(service Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {

		type userreq struct {
			//Username string `json:"username"`
			Email string `json:"email"`
			Password string `json:"password"`
		}

		var req struct {
			User userreq `json:"user"`
		}

		if err := c.Read(&req); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		token, err := service.Login(c.Request.Context(), req.User.Email, req.User.Password)
		if err != nil {
			return err
		}
		return c.Write(struct {
			Token string `json:"token"`
		}{token})
	}
}

// register returns a handler that handles user register request.
func register(service Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {

		var input CreateUserRequest

		if err := c.Read(&input); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		id, err := service.Register(c.Request.Context(), input)
		if err != nil {
			return err
		}
		return c.Write(struct {
			ID string `json:"id"`
		}{id})
	}
}

func query(service Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		
		count, err := service.Count(c.Request.Context())
		if err != nil {
			return err
		}
		pages := pagination.NewFromRequest(c.Request, count)
		users, err := service.Query(c.Request.Context(), pages.Offset(), pages.Limit())
		if err != nil {
			return err
		}
		pages.Items = users
		return c.Write(pages)
	}
}