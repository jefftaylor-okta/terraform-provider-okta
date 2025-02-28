package okta

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceOktaProfileMapping_crud(t *testing.T) {
	resourceName := fmt.Sprintf("%s.test", profileMapping)
	mgr := newFixtureManager(profileMapping, t.Name())
	config := mgr.GetFixtures("basic.tf", t)
	updatedConfig := mgr.GetFixtures("updated.tf", t)
	preventDelete := mgr.GetFixtures("prevent_delete.tf", t)

	oktaResourceTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		CheckDestroy:      checkResourceDestroy(profileMapping, doesOktaProfileExist),
		Steps: []resource.TestStep{
			{
				Config: preventDelete,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "delete_when_absent", "false"),
				),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "delete_when_absent", "true"),
				),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "delete_when_absent", "true"),
				),
			},
		},
	})
}

func doesOktaProfileExist(profileID string) (bool, error) {
	client := sdkSupplementClientForTest()
	_, response, err := client.GetEmailTemplate(context.Background(), profileID)
	return doesResourceExist(response, err)
}
