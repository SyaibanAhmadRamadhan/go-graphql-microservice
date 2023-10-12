package repository

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/SyaibanAhmadRamadhan/go-graphql-microservice/util"
)

type Header struct {
	YourCustomHeader string
}

type QueryParam struct {
	Order     string
	Cursor    string
	StartTime string
	EndTime   string
	Type      string
}
type AllServiceRepoImpl struct {
	Account *OneRepositoryImpl
	User    *TwoRepositoryImpl
	Finance *ThreeRepositoryImpl
}

func NewAllServiceRepoImpl() *AllServiceRepoImpl {
	return &AllServiceRepoImpl{
		Account: &OneRepositoryImpl{},
		User:    &TwoRepositoryImpl{},
		Finance: &ThreeRepositoryImpl{},
	}
}

var ResponseGateway = map[string]any{
	"data":    nil,
	"errors":  "bad gateway",
	"message": "bad gateway",
	"status":  502,
}

func FetchResponse(r *http.Response) map[string]any {
	var resMap map[string]any
	err := json.NewDecoder(r.Body).Decode(&resMap)
	if err != nil {
		log.Warn().Msgf(util.LogErrDecode, r.Body, err)
		return ResponseGateway
	}

	return resMap
}
