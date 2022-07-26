package article

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
	"net/http"
)

type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetEmail returns the user email.
	GetEmail() string
	GetUserName() string
	GetBio() string
	GetImage() string
}

type Author struct {
	Username	string	`json:"username"`
	Bio string `json:"bio"`
	Image string `json:"image"`
	Following bool `json:"following"`
}

type Article struct {
	Slug        	string 	`json:"slug"`
	Title       	string 	`json:"title"`
	Description 	string 	`json:"description"`
	Body        	string 	`json:"body"`
	CreatedAt 		string 	`json:"createdAt"`
	UpdatedAt 		string 	`json:"updatedAt"`
	Favorited   	bool 	`json:"favorited"`
	FavoritesCount 	string 	`json:"favoritesCount"`
	Author Author `json:"author"`
	TagList []string `json:"tagList"`
}

type resource struct {
	logger  log.Logger
}

type ArticleRequest struct {
	Title       	string 	`json:"title"`
	Description 	string 	`json:"description"`
	Body        	string 	`json:"body"`
	TagList []string `json:"tagList"`
}

type CreateArticleRequest struct {
	Article       	ArticleRequest 	`json:"article"`
}

type CreateArticleResponse struct {
	Article       	Article 	`json:"article"`
}

type UpdateSetting struct {
	Email 		string `json:"email"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	Bio  		string `json:"bio"`
	Image 		string `json:"image"`
	Token 		string `json:"token"`
}

type UpdateSettingRequest struct {
	User 		UpdateSetting `json:"user"`
}

type UpdateSettingResponse struct {
	User 		UpdateSetting `json:"user"`
	Token 		string `json:token`
}

type Comment struct {
	Id string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Body string `json:"body"`
	Author Author `json:"author"`
}

var articles = []Article {Article{Slug:"Create-a-new-implementation-1", Title:"Create a new implementation", Description:"join the community by creating a new implementation", Body:"Share your knowledge and enpower the community by creating a new implementation", CreatedAt:"2021-11-24T12:11:08.212Z",UpdatedAt:"2021-11-24T12:11:08.212Z",Favorited:false,FavoritesCount:"3065", Author:Author{Username:"rootz",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}, TagList:[]string {"implementations"}},
	Article{Slug:"Explore-implementations-1",Title:"Explore implementations",Description:"discover the implementations created by the RealWorld community",Body:"Over 100 implementations have been created using various languages, libraries, and frameworks.\n\nExplore them on CodebaseShow.",TagList:[]string {"codebaseShow","implementations"},CreatedAt:"2021-11-24T12:11:07.952Z",UpdatedAt:"2021-11-24T12:11:07.952Z",Favorited:false,FavoritesCount:"1787", Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}},
	Article{Slug:"Welcome-to-RealWorld-project-1",Title:"Welcome to RealWorld project",Description:"Exemplary fullstack Medium.com clone powered by React, Angular, Node, Django, and many more",Body:"See how the exact same Medium.com clone (called Conduit) is built using different frontends and backends. Yes, you can mix and match them, because they all adhere to the same API spec",TagList:[]string {"welcome","introduction"},CreatedAt:"2021-11-24T12:11:07.557Z",UpdatedAt:"2021-11-24T12:11:07.557Z",Favorited:false,FavoritesCount:"1262",Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}}}
var comments = []Comment {Comment{Id:"5",CreatedAt:"2021-11-24T12:11:08.480Z",UpdatedAt:"2021-11-24T12:11:08.480Z",Body:"If someone else has started working on an implementation, consider jumping in and helping them! by contacting the author.",Author:Author{Username:"rootz",Bio:"nil",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}},
Comment{Id:"4",CreatedAt:"2021-11-24T12:11:08.340Z",UpdatedAt:"2021-11-24T12:11:08.340Z",Body:"Before starting a new implementation, please check if there is any work in progress for the stack you want to work on.",Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}}}

func RegisterHandlers(rg *routing.RouteGroup, authHandler routing.Handler, logger log.Logger) {
	res := resource{logger}

	rg.Get("/articles", res.geta)
	rg.Get("/articles/feed", res.geta)
	rg.Get(`/articles/<slug>`, res.gets)
	rg.Get(`/articles/<slug>/comments`, res.getc)

	rg.Put("/user", func(c *routing.Context) error{
		var input UpdateSettingRequest
		if err := c.Read(&input); err != nil {
			logger.With(c.Request.Context()).Info(err)
			return errors.BadRequest("")
		}
		
		token, err := generateJWT(entity.User{ID: "100", Email: input.User.Email, Username:input.User.Username, Bio:input.User.Bio, Image:input.User.Image})
		if err != nil {
			return err
		}
		//input.User.Token = token
		//return c.Write(input)
		return c.Write(struct {
			Token string `json:"token"`
		}{token})
	})

	rg.Get("/tags", func(c *routing.Context) error {
		var tags []string
		pages := pagination.NewFromRequest(c.Request, len(tags))
		
		tags = []string {"implementations"}
		pages.Items = tags	
		return c.Write(pages)
	})

	rg.Post("/articles", func(c *routing.Context) error {

		var input CreateArticleRequest

		if err := c.Read(&input); err != nil {
			//log.logger.With(c.Request.Context()).Info(err)
			return errors.BadRequest("")
		}

		tl := input.Article.TagList[:]

		article := Article{
			Slug: input.Article.Title,
			Title: input.Article.Title,
			Description: input.Article.Description,
			Body: input.Article.Body,
			TagList: tl,
			CreatedAt:"2021-11-24T12:11:07.557Z",
			UpdatedAt:"2021-11-24T12:11:07.557Z",
			Favorited:false,
			FavoritesCount:"1262",
			Author:Author{Username:"Gerome", Bio:"null", Image:"https://api.realworld.io/images/demo-avatar.png", Following:false}}

		articles = append(articles, article)

		return c.WriteWithStatus(article, http.StatusCreated)
	})
}

func (r resource) geta(c *routing.Context) error {

		pages := pagination.NewFromRequest(c.Request, len(articles))
		pages.Items = articles

		return c.Write(pages)
}

func (r resource) gets(c *routing.Context) error {

		//return c.Write("update user " + c.Param("slug"))
		article := CreateArticleResponse{Article:articles[0]}

		return c.Write(article)
}

func (r resource) getc(c *routing.Context) error {
	pages := pagination.NewFromRequest(c.Request, len(articles))
	pages.Items = comments
	return c.Write(pages)
}

// generateJWT generates a JWT that encodes an identity.
func generateJWT(identity Identity) (string, error) {

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"email": identity.GetEmail(),
		"username": identity.GetUserName(),
		"bio": identity.GetBio(),
		"image": identity.GetImage(),
		"iat": time.Now().Unix(),
		"exp":  time.Now().Add(time.Duration(72) * time.Hour).Unix(),
	}).SignedString([]byte("LxsKJywDL5O5PvgODZhBH12KE6k2yL8E"))
}