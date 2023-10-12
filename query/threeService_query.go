package query

import (
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/args"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/repository"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/types"
)

func GetPayment(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.PaymentType, "PaymentResponse"),
		Args: args.InfiniteScroll,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetPayment(&repository.Header{
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

func GetSpendingType(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingTypeType, "SpendingTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":   args.OrderInfiniteScroll,
			"cursor":  args.CursorInfiniteScroll,
			"periode": args.PeriodeBool,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			var configValue string
			if p.Args["periode"].(bool) {
				p.Args["config_name"] = "MONTHLY_PERIOD"
				resolveProfileConfig, _ := GetProfileConfig(&repository.OneRepositoryImpl{}).Resolve(p)

				parse := resolveProfileConfig.(map[string]any)
				data, ok := parse["data"].(map[string]any)
				if !ok {
					return resolveProfileConfig, nil
				}

				configValue = data["config_value"].(string)
			}

			resp, err := finance.GetSpendingType(&repository.Header{
				YourCustomHeader: header.Get("YourCustomHeader"),
			}, p.Args["order"].(string), p.Args["cursor"].(string), configValue)

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

func GetDetailSpendingType(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingTypeList, "SpendingTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailSpendingType(&repository.Header{
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

func GetIncomeType(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeTypeType, "IncomeTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetIncomeType(&repository.Header{
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

func GetDetailIncomeType(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeListType, "IncomeTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailIncomeType(&repository.Header{
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

func GetBalance(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.BalanceType, "BalanceResponse"),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetBalance(&repository.Header{
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

func GetIncomeHistory(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeHistoryTypeList, "IncomeHistoryTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
			"start":  args.StartTimeParam,
			"end":    args.EndTimeParam,
			"type": &graphql.ArgumentConfig{
				Type:         args.TypeParam,
				DefaultValue: "",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetIncomeHistory(&repository.Header{
				YourCustomHeader: header.Get("YourCustomHeader"),
			}, &repository.QueryParam{
				Order:     p.Args["order"].(string),
				Cursor:    p.Args["cursor"].(string),
				StartTime: p.Args["start"].(string),
				EndTime:   p.Args["end"].(string),
				Type:      p.Args["type"].(string),
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

func GetDetailIncomeHistory(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.IncomeHistoryType, "IncomeHistoryTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailIncomeHistory(&repository.Header{
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

func GetSpendingHistory(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingHistoryTypeList, "SpendingHistoryTypeListResponse"),
		Args: graphql.FieldConfigArgument{
			"order":  args.OrderInfiniteScroll,
			"cursor": args.CursorInfiniteScroll,
			"start":  args.StartTimeParam,
			"end":    args.EndTimeParam,
			"type": &graphql.ArgumentConfig{
				Type:         args.TypeParam,
				DefaultValue: "",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetSpendingHistory(&repository.Header{
				YourCustomHeader: header.Get("YourCustomHeader"),
			}, &repository.QueryParam{
				Order:     p.Args["order"].(string),
				Cursor:    p.Args["cursor"].(string),
				StartTime: p.Args["start"].(string),
				EndTime:   p.Args["end"].(string),
				Type:      p.Args["type"].(string),
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

func GetDetailSpendingHistory(finance *repository.ThreeRepositoryImpl) *graphql.Field {
	return &graphql.Field{
		Type: types.Response(types.SpendingHistoryType, "SpendingHistoryTypeResponse"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			header := p.Context.Value("headers").(http.Header)

			resp, err := finance.GetDetailSpendingHistory(&repository.Header{
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
