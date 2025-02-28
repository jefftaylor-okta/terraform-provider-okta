package okta

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceOktaDomain(t *testing.T) {
	mgr := newFixtureManager(domain, t.Name())
	config := mgr.GetFixtures("basic.tf", t)
	resourceName := fmt.Sprintf("%s.test", domain)
	domainName := fmt.Sprintf("testacc-%d.example.com", mgr.Seed)

	oktaResourceTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		CheckDestroy:      checkResourceDestroy(domain, domainExists),
		Steps: []resource.TestStep{
			{
				ExpectNonEmptyPlan: true,
				Config:             config,
				Check: resource.ComposeTestCheckFunc(
					ensureResourceExists(resourceName, domainExists),
					resource.TestCheckResourceAttr(resourceName, "name", domainName),
					resource.TestCheckResourceAttr(resourceName, "dns_records.#", "2"),
				),
			},
		},
	})
}

func domainExists(id string) (bool, error) {
	client := sdkV2ClientForTest()
	domain, resp, err := client.Domain.GetDomain(context.Background(), id)
	if err := suppressErrorOn404(resp, err); err != nil {
		return false, err
	}
	return domain != nil, nil
}
