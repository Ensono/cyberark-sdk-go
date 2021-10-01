package cyberark_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/amido/cyberark-sdk-go/cyberark/config"
	"github.com/amido/cyberark-sdk-go/cyberark/models"
	"github.com/amido/cyberark-sdk-go/services/authenticationApi"
)

func TestMain(m *testing.M) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cfg := config.NewConfig().WithUsername("Administrator").WithPassword("CmHpWaMCHRrj9!c").WithEndpoint("https://pam-int-master.nonprod.idam.cloud.iairgroup.com")
	svc := authenticationApi.New(cfg)

	_, _, err := svc.DoAuth(models.AuthReq{Username: }
		Body: models.ServerSettings{
			FederationInfo: &models.FederationInfo{
				BaseUrl:       pf.String("https://localhost:9031"),
				Saml2EntityId: pf.String("testing"),
			},
			RolesAndProtocols: &models.RolesAndProtocols{
				IdpRole: &models.IdpRole{
					Enable: pf.Bool(true),
					Saml20Profile: &models.SAML20Profile{
						Enable: pf.Bool(true),
					},
				},
				OauthRole: &models.OAuthRole{
					EnableOauth: pf.Bool(true),
				},
				SpRole: &models.SpRole{
					Enable: pf.Bool(true),
					Saml20Profile: &models.SpSAML20Profile{
						Enable: pf.Bool(true),
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to login into CA %s", err)
	}
	os.Exit(m.Run())
}
