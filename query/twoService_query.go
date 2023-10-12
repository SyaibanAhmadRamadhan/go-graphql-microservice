package query

import (
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/repository"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/types"
)

func GetUser(userRepo *repository.TwoRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.UserType, "UserResponse"),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := userRepo.GetUser(&repository.Header{
				YourCustomHeader: header.Get("YourCustomHeader"),
			})

			if err != nil {
				httpResp := map[string]any{
					"status":  500,
					"message": "internal server error",
					"errors":  err.Error(),
					"data":    nil,
				}
				return httpResp, nil
			}

			return resp, nil
		},
	}
}
