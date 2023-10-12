package query

import (
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/args"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/repository"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/types"
)

func GetProfile(profileRepo *repository.OneRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.ProfileType, "ProfileResponse"),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)
			resp, err := profileRepo.GetProfile(&repository.Header{
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

func GetProfileConfig(account *repository.OneRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.ProfileConfigType, "ProfileConfigResponse"),
		Args: graphql.FieldConfigArgument{
			"config_name": &graphql.ArgumentConfig{
				Type:         args.ProfileConfigEnum,
				DefaultValue: "MONTHLY_PERIOD",
				Description:  "Input konfigurasi profil Config, by default MONTHLY_PERIOD",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := account.GetProfileConfig(&repository.Header{
				YourCustomHeader: header.Get("YourCustomHeader"),
			}, p.Args["config_name"].(string))

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

func GetNotification(account *repository.OneRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.NotificationType, "NotificationResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := account.GetNotification(&repository.Header{
				YourCustomHeader: header.Get("YourCustomHeader"),
			}, p.Args["order"].(string), p.Args["cursor"].(string))

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

func GetDetailNotification(account *repository.OneRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.NotificationTypeList, "NotificationListResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := account.GetDetailNotification(&repository.Header{
				YourCustomHeader: header.Get("YourCustomHeader"),
			}, p.Args["id"].(string))

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
