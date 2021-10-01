package authenticationApi

import (
	"context"
	"net/http"

	"github.com/amido/cyberark-sdk-go/cyberark"
	"github.com/amido/cyberark-sdk-go/cyberark/client"
	"github.com/amido/cyberark-sdk-go/cyberark/client/metadata"
	"github.com/amido/cyberark-sdk-go/cyberark/config"
	"github.com/amido/cyberark-sdk-go/cyberark/models"
	"github.com/amido/cyberark-sdk-go/cyberark/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "AuthenticationApi"
)

type AuthenticationApiService struct {
	*client.CaClient
}

// New creates a new instance of the AuthenticationApiService client.
func New(cfg *config.Config) *AuthenticationApiService {

	return &AuthenticationApiService{CaClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  cyberark.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthenticationApi operation
func (c *AuthenticationApiService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetAuthenticationApiSettings - Get the Authentication API settings.
//RequestType: GET
//Input:
func (s *AuthenticationApiService) DoAuth() (input *models.AuthReq, output *models.AuthToken, resp *http.Response, err error) {
	return s.GetAuthenticationApiSettingsWithContext(context.Background())
}

//GetAuthenticationApiSettingsWithContext - Get the Authentication API settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *AuthenticationApiService) GetAuthenticationApiSettingsWithContext(ctx context.Context) (input *models.AuthReq, output *models.AuthToken, resp *http.Response, err error) {
	path := "/PasswordVault/v10/logon/cyberark"
	op := &request.Operation{
		Name:       "DoAuthentication",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	// output = string(*models.AuthToken)
	output = &models.AuthToken{}
	req := s.newRequest(op, input, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return input, output, req.HTTPResponse, nil
	}
	return input, nil, req.HTTPResponse, req.Error
}
