package okta

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-okta/sdk"
	"github.com/okta/terraform-provider-okta/sdk/query"
)

func dataSourceAuthServer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAuthServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"audiences": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"kid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials_last_rotated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials_next_rotation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials_rotation_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issuer_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAuthServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	servers, _, err := getOktaClientFromMetadata(m).AuthorizationServer.ListAuthorizationServers(ctx, &query.Params{Q: name, Limit: defaultPaginationLimit})
	if err != nil {
		return diag.Errorf("failed to find auth server '%s': %v", name, err)
	}
	var authServer *sdk.AuthorizationServer
	for i := range servers {
		if servers[i].Name == name {
			authServer = servers[i]
		}
	}
	if authServer == nil {
		return diag.Errorf("authorization server with name '%s' does not exist", name)
	}
	d.SetId(authServer.Id)
	_ = d.Set("name", authServer.Name)
	_ = d.Set("description", authServer.Description)
	_ = d.Set("audiences", convertStringSliceToSet(authServer.Audiences))
	if authServer.Credentials != nil && authServer.Credentials.Signing != nil {
		_ = d.Set("kid", authServer.Credentials.Signing.Kid)
		_ = d.Set("credentials_rotation_mode", authServer.Credentials.Signing.RotationMode)
		if authServer.Credentials.Signing.NextRotation != nil {
			_ = d.Set("credentials_next_rotation", authServer.Credentials.Signing.NextRotation.String())
		}
		if authServer.Credentials.Signing.LastRotated != nil {
			_ = d.Set("credentials_last_rotated", authServer.Credentials.Signing.LastRotated.String())
		}
	}
	_ = d.Set("status", authServer.Status)
	_ = d.Set("issuer", authServer.Issuer)
	// Do not sync these unless the issuer mode is specified since it is an EA feature
	if authServer.IssuerMode != "" {
		_ = d.Set("issuer_mode", authServer.IssuerMode)
	}
	return nil
}
