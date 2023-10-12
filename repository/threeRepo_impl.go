package repository

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/infra"
	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/util"
)

type ThreeRepositoryImpl struct {
}

func (p *ThreeRepositoryImpl) GetPayment(header *Header, order, cursor string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/payment?order=%s&cursor=%s", infra.ServiceThreeUrl, order, cursor)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetSpendingType(header *Header, order, cursor string, configValue string) (map[string]any, error) {
	var route string
	if configValue == "" {
		route = fmt.Sprintf("%s/finance/spending-type?order=%s&cursor=%s", infra.ServiceThreeUrl, order, cursor)
	} else {
		route = fmt.Sprintf("%s/finance/spending-type/%s?order=%s&cursor=%s", infra.ServiceThreeUrl, configValue, order, cursor)
	}

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetDetailSpendingType(header *Header, id string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/spending-type/detail/%s", infra.ServiceThreeUrl, id)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetIncomeType(header *Header, order, cursor string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/income-type?order=%s&cursor=%s", infra.ServiceThreeUrl, order, cursor)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetDetailIncomeType(header *Header, id string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/income-type/detail/%s", infra.ServiceThreeUrl, id)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetBalance(header *Header) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/balance", infra.ServiceThreeUrl)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetIncomeHistory(header *Header, param *QueryParam) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/income-history?order=%s&cursor=%s&start=%s&end=%s&type=%s", infra.ServiceThreeUrl, param.Order, param.Cursor, param.StartTime, param.EndTime, param.Type)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetDetailIncomeHistory(header *Header, id string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/income-history/%s", infra.ServiceThreeUrl, id)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetSpendingHistory(header *Header, param *QueryParam) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/spending-history?order=%s&cursor=%s&start=%s&end=%s&type=%s", infra.ServiceThreeUrl, param.Order, param.Cursor, param.StartTime, param.EndTime, param.Type)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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

func (p *ThreeRepositoryImpl) GetDetailSpendingHistory(header *Header, id string) (map[string]any, error) {
	route := fmt.Sprintf("%s/finance/spending-history/%s", infra.ServiceThreeUrl, id)

	request, err := http.NewRequest("GET", route, nil)
	if err != nil {
		log.Warn().Msgf(util.LogErrHttpNewRequest, err)
		return nil, err
	}

	request.Header.Set("X-Api-Key", infra.ServiceThreeKey)

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
