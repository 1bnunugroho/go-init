package article

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	//"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
	//"time"
	"github.com/qiangxue/go-rest-api/pkg/pagination"
)

func RegisterHandlers(rg *routing.RouteGroup, authHandler routing.Handler, logger log.Logger) {
	rg.Get("/articles", func(c *routing.Context) error {
		pages := pagination.NewFromRequest(c.Request, 3)

		albums := Article{Slug:"Create-a-new-implementation-1", Title:"Create a new implementation", Description:"join the community by creating a new implementation", Body:"Share your knowledge and enpower the community by creating a new implementation", CreatedAt:"2021-11-24T12:11:08.212Z",UpdatedAt:"2021-11-24T12:11:08.212Z",Favorited:false,FavoritesCount:"3065"}
		pages.Items = albums

		return c.Write(pages)
	})
}


type Article struct {
	Slug        	string `json:"slug"`
	Title       	string `json:"title"`
	Description 	string `json:"description"`
	Body        	string 				`json:"body"`
	CreatedAt 		string 			`json:"createdAt"`
	UpdatedAt 		string 			`json:"updatedAt"`
	Favorited   	bool 			`json:"favorited"`
	FavoritesCount 	string 				`json:"favoritesCount"`
}