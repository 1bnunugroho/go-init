package article

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	//"time"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
	"net/http"
)

//var articles []Article

var articles = []Article {Article{Slug:"Create-a-new-implementation-1", Title:"Create a new implementation", Description:"join the community by creating a new implementation", Body:"Share your knowledge and enpower the community by creating a new implementation", CreatedAt:"2021-11-24T12:11:08.212Z",UpdatedAt:"2021-11-24T12:11:08.212Z",Favorited:false,FavoritesCount:"3065", Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}, TagList:[]string {"implementations"}},
	Article{Slug:"Explore-implementations-1",Title:"Explore implementations",Description:"discover the implementations created by the RealWorld community",Body:"Over 100 implementations have been created using various languages, libraries, and frameworks.\n\nExplore them on CodebaseShow.",TagList:[]string {"codebaseShow","implementations"},CreatedAt:"2021-11-24T12:11:07.952Z",UpdatedAt:"2021-11-24T12:11:07.952Z",Favorited:false,FavoritesCount:"1787", Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}},
	Article{Slug:"Welcome-to-RealWorld-project-1",Title:"Welcome to RealWorld project",Description:"Exemplary fullstack Medium.com clone powered by React, Angular, Node, Django, and many more",Body:"See how the exact same Medium.com clone (called Conduit) is built using different frontends and backends. Yes, you can mix and match them, because they all adhere to the same API spec",TagList:[]string {"welcome","introduction"},CreatedAt:"2021-11-24T12:11:07.557Z",UpdatedAt:"2021-11-24T12:11:07.557Z",Favorited:false,FavoritesCount:"1262",Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}}}

func RegisterHandlers(rg *routing.RouteGroup, authHandler routing.Handler, logger log.Logger) {
	res := resource{logger}

	rg.Get("/articles", res.geta)
	rg.Get("/articles/feed", res.geta)

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

		article := Article{
			Slug: input.Title,
			Title: input.Title,
			Description: input.Description,
			Body: input.Body,
			TagList:[]string {"welcome","introduction"},
			CreatedAt:"2021-11-24T12:11:07.557Z",
			UpdatedAt:"2021-11-24T12:11:07.557Z",
			Favorited:false,
			FavoritesCount:"1262",
			Author:Author{Username:"Gerome", Bio:"null", Image:"https://api.realworld.io/images/demo-avatar.png", Following:false}}

		articles = append(articles, article)

		return c.WriteWithStatus(article, http.StatusCreated)
	})
}

type resource struct {
	logger  log.Logger
}

type CreateArticleRequest struct {
	Title       	string 	`json:"title"`
	Description 	string 	`json:"description"`
	Body        	string 	`json:"body"`
	TagList []string `json:"tagList"`
}

func (r resource) geta(c *routing.Context) error {
	//func(c *routing.Context) error {
		
		/*var articles []Article

		articles = []Article {Article{Slug:"Create-a-new-implementation-1", Title:"Create a new implementation", Description:"join the community by creating a new implementation", Body:"Share your knowledge and enpower the community by creating a new implementation", CreatedAt:"2021-11-24T12:11:08.212Z",UpdatedAt:"2021-11-24T12:11:08.212Z",Favorited:false,FavoritesCount:"3065", Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}, TagList:[]string {"implementations"}},
			Article{Slug:"Explore-implementations-1",Title:"Explore implementations",Description:"discover the implementations created by the RealWorld community",Body:"Over 100 implementations have been created using various languages, libraries, and frameworks.\n\nExplore them on CodebaseShow.",TagList:[]string {"codebaseShow","implementations"},CreatedAt:"2021-11-24T12:11:07.952Z",UpdatedAt:"2021-11-24T12:11:07.952Z",Favorited:false,FavoritesCount:"1787", Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}},
			Article{Slug:"Welcome-to-RealWorld-project-1",Title:"Welcome to RealWorld project",Description:"Exemplary fullstack Medium.com clone powered by React, Angular, Node, Django, and many more",Body:"See how the exact same Medium.com clone (called Conduit) is built using different frontends and backends. Yes, you can mix and match them, because they all adhere to the same API spec",TagList:[]string {"welcome","introduction"},CreatedAt:"2021-11-24T12:11:07.557Z",UpdatedAt:"2021-11-24T12:11:07.557Z",Favorited:false,FavoritesCount:"1262",Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}}}
		
		articles = append(articles, Article{Slug:"Create-a-new-implementation-1", Title:"Create a new implementation", Description:"join the community by creating a new implementation", Body:"Share your knowledge and enpower the community by creating a new implementation", CreatedAt:"2021-11-24T12:11:08.212Z",UpdatedAt:"2021-11-24T12:11:08.212Z",Favorited:false,FavoritesCount:"3065", Author:Author{Username:"Gerome",Bio:"null",Image:"https://api.realworld.io/images/demo-avatar.png",Following:false}, TagList:[]string {"implementations"}})*/

		pages := pagination.NewFromRequest(c.Request, len(articles))
		pages.Items = articles

		return c.Write(pages)
	//}
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