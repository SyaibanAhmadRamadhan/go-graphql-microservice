package repository

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/infra"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/util"
)

type TwoRepositoryImpl struct {
}

func (u *TwoRepositoryImpl) GetUser(header *Header) (map[string]any, error) {
	route := fmt.Sprintf("%s/auth/user", infra.ServiceTwoUrl)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceTwoKey)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Warn().Msgf(util.LogErrClientDo, err)
		return nil, err
	}
	defer func() {
		if errBody := response.Body.Close(); errBody != nil {
			log.Warn().Msgf(util.LogErrClientDoClose, err)
		}
	}()

	resp := FetchResponse(response)

	return resp, nil
}
