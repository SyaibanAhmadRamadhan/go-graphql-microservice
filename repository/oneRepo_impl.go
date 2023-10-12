package repository

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/infra"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/util"
)

type OneRepositoryImpl struct {
}

func (p *OneRepositoryImpl) GetProfile(header *Header) (map[string]any, error) {
	route := fmt.Sprintf("%s/account/profile", infra.ServiceOneUrl)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceOneKey)

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

func (p *OneRepositoryImpl) GetProfileConfig(header *Header, configName string) (map[string]any, error) {
	route := fmt.Sprintf("%s/account/profile-config/%s", infra.ServiceOneUrl, configName)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceOneKey)

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
func (p *OneRepositoryImpl) GetNotification(header *Header, order, cursor string) (map[string]any, error) {
	route := fmt.Sprintf("%s/account/notif?order=%s&cursor=%s", infra.ServiceOneUrl, order, cursor)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceOneKey)

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

func (p *OneRepositoryImpl) GetDetailNotification(header *Header, id string) (map[string]any, error) {
	route := fmt.Sprintf("%s/account/notif/%s", infra.ServiceOneUrl, id)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceOneKey)

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
