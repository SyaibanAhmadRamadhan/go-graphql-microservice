package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/infra"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/query"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/repository"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/util"
)

func main() {
	infra.InitEnv()
	infra.LogInit()

	repo := repository.NewAllServiceRepoImpl()
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: query.GetRootFields(repo),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal().Msgf("failed new scheman graph | err %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema,
	})

	http.Handle("/gql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != infra.Key {
			httpResp := map[string]any{
				"status":  403,
				"message": "forbidden",
				"errors":  "invalid api key page",
				"data":    nil,
			}

			w.WriteHeader(403)
			if errEncode := json.NewEncoder(w).Encode(httpResp); errEncode != nil {
				log.Err(errEncode).Msgf(util.LogErrEncode, httpResp, errEncode)
			}
			return
		}

		ctxHeaders := context.WithValue(r.Context(), "headers", r.Header)
		r = r.WithContext(ctxHeaders)

		w.Header().Set("Authorization", r.Header.Get("Authorization"))

		httpHandler.ServeHTTP(w, r)
	}))

	http.ListenAndServe(":7007", nil)
}
