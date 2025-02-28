package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/okta/terraform-provider-okta/sdk/query"
)

type App interface {
	IsApplicationInstance() bool
}

type ApplicationResource resource

type Application struct {
	Embedded      interface{}               `json:"_embedded,omitempty"`
	Links         interface{}               `json:"_links,omitempty"`
	Accessibility *ApplicationAccessibility `json:"accessibility,omitempty"`
	Created       *time.Time                `json:"created,omitempty"`
	Credentials   *ApplicationCredentials   `json:"credentials,omitempty"`
	Features      []string                  `json:"features,omitempty"`
	Id            string                    `json:"id,omitempty"`
	Label         string                    `json:"label,omitempty"`
	LastUpdated   *time.Time                `json:"lastUpdated,omitempty"`
	Licensing     *ApplicationLicensing     `json:"licensing,omitempty"`
	Name          string                    `json:"name,omitempty"`
	Profile       interface{}               `json:"profile,omitempty"`
	Settings      *ApplicationSettings      `json:"settings,omitempty"`
	SignOnMode    string                    `json:"signOnMode,omitempty"`
	Status        string                    `json:"status,omitempty"`
	Visibility    *ApplicationVisibility    `json:"visibility,omitempty"`
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) IsApplicationInstance() bool {
	return true
}

// Fetches an application from your Okta organization by &#x60;id&#x60;.
func (m *ApplicationResource) GetApplication(ctx context.Context, appId string, appInstance App, qp *query.Params) (App, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	application := appInstance

	resp, err := rq.Do(ctx, req, &application)
	if err != nil {
		return nil, resp, err
	}

	return application, resp, nil
}

// Updates an application in your organization.
func (m *ApplicationResource) UpdateApplication(ctx context.Context, appId string, body App) (App, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	application := body

	resp, err := rq.Do(ctx, req, &application)
	if err != nil {
		return nil, resp, err
	}

	return application, resp, nil
}

// Removes an inactive application.
func (m *ApplicationResource) DeleteApplication(ctx context.Context, appId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates apps added to your organization with pagination. A subset of apps can be returned that match a supported filter expression or query.
func (m *ApplicationResource) ListApplications(ctx context.Context, qp *query.Params) ([]App, *Response, error) {
	url := "/api/v1/apps"
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var application []Application

	resp, err := rq.Do(ctx, req, &application)
	if err != nil {
		return nil, resp, err
	}

	apps := make([]App, len(application))
	for i := range application {
		apps[i] = &application[i]
	}
	return apps, resp, nil
}

// Adds a new application to your Okta organization.
func (m *ApplicationResource) CreateApplication(ctx context.Context, body App, qp *query.Params) (App, *Response, error) {
	url := "/api/v1/apps"
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	application := body

	resp, err := rq.Do(ctx, req, &application)
	if err != nil {
		return nil, resp, err
	}

	return application, resp, nil
}

// Get default Provisioning Connection for application
func (m *ApplicationResource) GetDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) (*ProvisioningConnection, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/connections/default", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var provisioningConnection *ProvisioningConnection

	resp, err := rq.Do(ctx, req, &provisioningConnection)
	if err != nil {
		return nil, resp, err
	}

	return provisioningConnection, resp, nil
}

// Set default Provisioning Connection for application
func (m *ApplicationResource) SetDefaultProvisioningConnectionForApplication(ctx context.Context, appId string, body ProvisioningConnectionRequest, qp *query.Params) (*ProvisioningConnection, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/connections/default", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var provisioningConnection *ProvisioningConnection

	resp, err := rq.Do(ctx, req, &provisioningConnection)
	if err != nil {
		return nil, resp, err
	}

	return provisioningConnection, resp, nil
}

// Activates the default Provisioning Connection for an application.
func (m *ApplicationResource) ActivateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/connections/default/lifecycle/activate", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivates the default Provisioning Connection for an application.
func (m *ApplicationResource) DeactivateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/connections/default/lifecycle/deactivate", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates Certificate Signing Requests for an application
func (m *ApplicationResource) ListCsrsForApplication(ctx context.Context, appId string) ([]*Csr, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var csr []*Csr

	resp, err := rq.Do(ctx, req, &csr)
	if err != nil {
		return nil, resp, err
	}

	return csr, resp, nil
}

// Generates a new key pair and returns the Certificate Signing Request for it.
func (m *ApplicationResource) GenerateCsrForApplication(ctx context.Context, appId string, body CsrMetadata) (*Csr, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var csr *Csr

	resp, err := rq.Do(ctx, req, &csr)
	if err != nil {
		return nil, resp, err
	}

	return csr, resp, nil
}

func (m *ApplicationResource) RevokeCsrFromApplication(ctx context.Context, appId, csrId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs/%v", appId, csrId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *ApplicationResource) GetCsrForApplication(ctx context.Context, appId, csrId string) (*Csr, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs/%v", appId, csrId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var csr *Csr

	resp, err := rq.Do(ctx, req, &csr)
	if err != nil {
		return nil, resp, err
	}

	return csr, resp, nil
}

func (m *ApplicationResource) PublishCerCert(ctx context.Context, appId, csrId, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs/%v/lifecycle/publish", appId, csrId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/x-x509-ca-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

func (m *ApplicationResource) PublishBinaryCerCert(ctx context.Context, appId, csrId, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs/%v/lifecycle/publish", appId, csrId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.AsBinary().WithAccept("application/json").WithContentType("application/x-x509-ca-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

func (m *ApplicationResource) PublishDerCert(ctx context.Context, appId, csrId, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs/%v/lifecycle/publish", appId, csrId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/pkix-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

func (m *ApplicationResource) PublishBinaryDerCert(ctx context.Context, appId, csrId, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs/%v/lifecycle/publish", appId, csrId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.AsBinary().WithAccept("application/json").WithContentType("application/pkix-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

func (m *ApplicationResource) PublishBinaryPemCert(ctx context.Context, appId, csrId, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/csrs/%v/lifecycle/publish", appId, csrId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.AsBinary().WithAccept("application/json").WithContentType("application/x-pem-file").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Enumerates key credentials for an application
func (m *ApplicationResource) ListApplicationKeys(ctx context.Context, appId string) ([]*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey []*JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Generates a new X.509 certificate for an application key credential
func (m *ApplicationResource) GenerateApplicationKey(ctx context.Context, appId string, qp *query.Params) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys/generate", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Gets a specific application key credential by kid
func (m *ApplicationResource) GetApplicationKey(ctx context.Context, appId, keyId string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys/%v", appId, keyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Clones a X.509 certificate for an application key credential from a source application to target application.
func (m *ApplicationResource) CloneApplicationKey(ctx context.Context, appId, keyId string, qp *query.Params) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys/%v/clone", appId, keyId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := rq.Do(ctx, req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Enumerates the client&#x27;s collection of secrets
func (m *ApplicationResource) ListClientSecretsForApplication(ctx context.Context, appId string) ([]*ClientSecret, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/secrets", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var clientSecret []*ClientSecret

	resp, err := rq.Do(ctx, req, &clientSecret)
	if err != nil {
		return nil, resp, err
	}

	return clientSecret, resp, nil
}

// Adds a new secret to the client&#x27;s collection of secrets.
func (m *ApplicationResource) CreateNewClientSecretForApplication(ctx context.Context, appId string, body ClientSecretMetadata) (*ClientSecret, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/secrets", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var clientSecret *ClientSecret

	resp, err := rq.Do(ctx, req, &clientSecret)
	if err != nil {
		return nil, resp, err
	}

	return clientSecret, resp, nil
}

// Removes a secret from the client&#x27;s collection of secrets.
func (m *ApplicationResource) DeleteClientSecretForApplication(ctx context.Context, appId, secretId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/secrets/%v", appId, secretId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a specific client secret by secretId
func (m *ApplicationResource) GetClientSecretForApplication(ctx context.Context, appId, secretId string) (*ClientSecret, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/secrets/%v", appId, secretId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var clientSecret *ClientSecret

	resp, err := rq.Do(ctx, req, &clientSecret)
	if err != nil {
		return nil, resp, err
	}

	return clientSecret, resp, nil
}

// Activates a specific client secret by secretId
func (m *ApplicationResource) ActivateClientSecretForApplication(ctx context.Context, appId, secretId string) (*ClientSecret, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/secrets/%v/lifecycle/activate", appId, secretId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var clientSecret *ClientSecret

	resp, err := rq.Do(ctx, req, &clientSecret)
	if err != nil {
		return nil, resp, err
	}

	return clientSecret, resp, nil
}

// Deactivates a specific client secret by secretId
func (m *ApplicationResource) DeactivateClientSecretForApplication(ctx context.Context, appId, secretId string) (*ClientSecret, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/secrets/%v/lifecycle/deactivate", appId, secretId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var clientSecret *ClientSecret

	resp, err := rq.Do(ctx, req, &clientSecret)
	if err != nil {
		return nil, resp, err
	}

	return clientSecret, resp, nil
}

// List Features for application
func (m *ApplicationResource) ListFeaturesForApplication(ctx context.Context, appId string) ([]*ApplicationFeature, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/features", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var applicationFeature []*ApplicationFeature

	resp, err := rq.Do(ctx, req, &applicationFeature)
	if err != nil {
		return nil, resp, err
	}

	return applicationFeature, resp, nil
}

// Fetches a Feature object for an application.
func (m *ApplicationResource) GetFeatureForApplication(ctx context.Context, appId, name string) (*ApplicationFeature, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/features/%v", appId, name)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var applicationFeature *ApplicationFeature

	resp, err := rq.Do(ctx, req, &applicationFeature)
	if err != nil {
		return nil, resp, err
	}

	return applicationFeature, resp, nil
}

// Updates a Feature object for an application.
func (m *ApplicationResource) UpdateFeatureForApplication(ctx context.Context, appId, name string, body CapabilitiesObject) (*ApplicationFeature, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/features/%v", appId, name)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var applicationFeature *ApplicationFeature

	resp, err := rq.Do(ctx, req, &applicationFeature)
	if err != nil {
		return nil, resp, err
	}

	return applicationFeature, resp, nil
}

// Lists all scope consent grants for the application
func (m *ApplicationResource) ListScopeConsentGrants(ctx context.Context, appId string, qp *query.Params) ([]*OAuth2ScopeConsentGrant, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/grants", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2ScopeConsentGrant []*OAuth2ScopeConsentGrant

	resp, err := rq.Do(ctx, req, &oAuth2ScopeConsentGrant)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2ScopeConsentGrant, resp, nil
}

// Grants consent for the application to request an OAuth 2.0 Okta scope
func (m *ApplicationResource) GrantConsentToScope(ctx context.Context, appId string, body OAuth2ScopeConsentGrant) (*OAuth2ScopeConsentGrant, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/grants", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2ScopeConsentGrant *OAuth2ScopeConsentGrant

	resp, err := rq.Do(ctx, req, &oAuth2ScopeConsentGrant)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2ScopeConsentGrant, resp, nil
}

// Revokes permission for the application to request the given scope
func (m *ApplicationResource) RevokeScopeConsentGrant(ctx context.Context, appId, grantId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/grants/%v", appId, grantId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetches a single scope consent grant for the application
func (m *ApplicationResource) GetScopeConsentGrant(ctx context.Context, appId, grantId string, qp *query.Params) (*OAuth2ScopeConsentGrant, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/grants/%v", appId, grantId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2ScopeConsentGrant *OAuth2ScopeConsentGrant

	resp, err := rq.Do(ctx, req, &oAuth2ScopeConsentGrant)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2ScopeConsentGrant, resp, nil
}

// Enumerates group assignments for an application.
func (m *ApplicationResource) ListApplicationGroupAssignments(ctx context.Context, appId string, qp *query.Params) ([]*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var applicationGroupAssignment []*ApplicationGroupAssignment

	resp, err := rq.Do(ctx, req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}

	return applicationGroupAssignment, resp, nil
}

// Removes a group assignment from an application.
func (m *ApplicationResource) DeleteApplicationGroupAssignment(ctx context.Context, appId, groupId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetches an application group assignment
func (m *ApplicationResource) GetApplicationGroupAssignment(ctx context.Context, appId, groupId string, qp *query.Params) (*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var applicationGroupAssignment *ApplicationGroupAssignment

	resp, err := rq.Do(ctx, req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}

	return applicationGroupAssignment, resp, nil
}

// Assigns a group to an application
func (m *ApplicationResource) CreateApplicationGroupAssignment(ctx context.Context, appId, groupId string, body ApplicationGroupAssignment) (*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var applicationGroupAssignment *ApplicationGroupAssignment

	resp, err := rq.Do(ctx, req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}

	return applicationGroupAssignment, resp, nil
}

// Activates an inactive application.
func (m *ApplicationResource) ActivateApplication(ctx context.Context, appId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/lifecycle/activate", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivates an active application.
func (m *ApplicationResource) DeactivateApplication(ctx context.Context, appId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/lifecycle/deactivate", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Update the logo for an application.
func (m *ApplicationResource) UploadApplicationLogo(ctx context.Context, appId, file string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/logo", appId)

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType(writer.FormDataContentType()).NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Assign an application to a specific policy. This unassigns the application from its currently assigned policy.
func (m *ApplicationResource) UpdateApplicationPolicy(ctx context.Context, appId, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/policies/%v", appId, policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Previews SAML metadata based on a specific key credential for an application
func (m *ApplicationResource) PreviewSAMLAppMetadata(ctx context.Context, appId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/sso/saml/metadata", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/xml").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Revokes all tokens for the specified application
func (m *ApplicationResource) RevokeOAuth2TokensForApplication(ctx context.Context, appId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/tokens", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Lists all tokens for the application
func (m *ApplicationResource) ListOAuth2TokensForApplication(ctx context.Context, appId string, qp *query.Params) ([]*OAuth2Token, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/tokens", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Token []*OAuth2Token

	resp, err := rq.Do(ctx, req, &oAuth2Token)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Token, resp, nil
}

// Revokes the specified token for the specified application
func (m *ApplicationResource) RevokeOAuth2TokenForApplication(ctx context.Context, appId, tokenId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/tokens/%v", appId, tokenId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a token for the specified application
func (m *ApplicationResource) GetOAuth2TokenForApplication(ctx context.Context, appId, tokenId string, qp *query.Params) (*OAuth2Token, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/tokens/%v", appId, tokenId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Token *OAuth2Token

	resp, err := rq.Do(ctx, req, &oAuth2Token)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Token, resp, nil
}

// Enumerates all assigned [application users](#application-user-model) for an application.
func (m *ApplicationResource) ListApplicationUsers(ctx context.Context, appId string, qp *query.Params) ([]*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users", appId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var appUser []*AppUser

	resp, err := rq.Do(ctx, req, &appUser)
	if err != nil {
		return nil, resp, err
	}

	return appUser, resp, nil
}

// Assigns an user to an application with [credentials](#application-user-credentials-object) and an app-specific [profile](#application-user-profile-object). Profile mappings defined for the application are first applied before applying any profile properties specified in the request.
func (m *ApplicationResource) AssignUserToApplication(ctx context.Context, appId string, body AppUser) (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users", appId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var appUser *AppUser

	resp, err := rq.Do(ctx, req, &appUser)
	if err != nil {
		return nil, resp, err
	}

	return appUser, resp, nil
}

// Removes an assignment for a user from an application.
func (m *ApplicationResource) DeleteApplicationUser(ctx context.Context, appId, userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users/%v", appId, userId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetches a specific user assignment for application by &#x60;id&#x60;.
func (m *ApplicationResource) GetApplicationUser(ctx context.Context, appId, userId string, qp *query.Params) (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users/%v", appId, userId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var appUser *AppUser

	resp, err := rq.Do(ctx, req, &appUser)
	if err != nil {
		return nil, resp, err
	}

	return appUser, resp, nil
}

// Updates a user&#x27;s profile for an application
func (m *ApplicationResource) UpdateApplicationUser(ctx context.Context, appId, userId string, body AppUser) (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users/%v", appId, userId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var appUser *AppUser

	resp, err := rq.Do(ctx, req, &appUser)
	if err != nil {
		return nil, resp, err
	}

	return appUser, resp, nil
}
