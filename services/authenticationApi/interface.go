package authenticationApi

import (
	"context"
	"net/http"

	"github.com/amido/cyberark-sdk-go/cyberark/models"
)

type AuthAPI interface {
	DoAuth() (input *models.AuthReq, output *models.AuthToken, resp *http.Response, err error)
	DoAuthWithContext(ctx context.Context) (input *models.AuthReq, output *models.AuthToken, resp *http.Response, err error)
}
