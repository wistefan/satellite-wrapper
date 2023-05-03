package main

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"wistefan/satellite-wrapper/generated"

	"encoding/base64"

	"github.com/fiware/dsba-pdp/ishare"
	"github.com/fiware/dsba-pdp/logging"
	"github.com/fiware/dsba-pdp/model"
	b58 "github.com/mr-tron/base58/base58"
)

type IssuerService struct {
	tokenHandler    *ishare.TokenHandler
	satelliteConfig SatelliteConfig
}

func getARObject(s SatelliteConfig) model.AuthorizationRegistry {

	return model.AuthorizationRegistry{Id: s.Id, Host: s.Host, TokenPath: s.TokenPath}
}

// GetIssuer - Returns a trusted issuer identified by its decentralised identifier (DID).
func (s *IssuerService) GetIssuer(ctx context.Context, did string) (res generated.ImplResponse, err error) {
	logging.Log().Infof("Get issuer %s", did)
	pubKey, err := getPublicKeyFromDID(did)
	logging.Log().Infof("Got public key: %s, Err: %v", pubKey, err)

	/**arConfig := getARObject(s.satelliteConfig)
	accessToken, httpErr := s.tokenHandler.GetTokenFromAR(&arConfig)

	if httpErr != (model.HttpError{}) {
		return res, httpErr.RootError
	}
	iShareClient := IShareClient{s.satelliteConfig}
	party, err := iShareClient.GetParty(accessToken, did)*/

	// TODO - update GetIssuer with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Issuer{}) or use other options such as http.Ok ...
	//return Response(200, Issuer{}), nil

	//TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(400, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(404, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(500, ProblemDetails{}), nil

	return generated.Response(http.StatusNotImplemented, nil), errors.New("GetIssuer method not implemented")
}

func getPublicKeyFromDID(did string) (publicKey string, err error) {
	didSlice := strings.Split(did, ":")
	keyString := didSlice[2]
	decodedArray, err := b58.Decode(keyString)
	if err != nil {
		return publicKey, err
	}

	return base64.StdEncoding.EncodeToString(decodedArray), nil

}

// GetIssuers - Returns a list of trusted issuers.
func (s *IssuerService) GetIssuers(ctx context.Context, pageSize float32, pageAfter string) (generated.ImplResponse, error) {
	// TODO - update GetIssuers with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, IssuersResponse{}) or use other options such as http.Ok ...
	//return Response(200, IssuersResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(400, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(500, ProblemDetails{}), nil

	return generated.Response(http.StatusNotImplemented, nil), errors.New("GetIssuers method not implemented")
}
