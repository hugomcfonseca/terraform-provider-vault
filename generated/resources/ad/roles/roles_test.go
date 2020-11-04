package roles

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/vault/api"
	"github.com/terraform-providers/terraform-provider-vault/generated/resources/ad"
	"github.com/terraform-providers/terraform-provider-vault/schema"
	"github.com/terraform-providers/terraform-provider-vault/util"
	"github.com/terraform-providers/terraform-provider-vault/vault"
)

var roleTestProvider = func() *schema.Provider {
	p := schema.NewProvider(vault.Provider())
	p.RegisterResource("vault_mount", vault.MountResource())
	p.RegisterResource("vault_ad_secret_backend", ad.ConfigResource())
	p.RegisterResource("vault_ad_secret_role", RoleResource())

	return p
}()

//func TestAccADSecretBackendRole_import(t *testing.T) {
//	path := acctest.RandomWithPrefix("tf-test-ad")
//	bindDN, bindPass, url := util.GetTestADCreds(t)
//	role := "bob"
//	serviceAccountName := "Bob"
//	ttl := 60
//
//	resource.Test(t, resource.TestCase{
//		Providers: map[string]terraform.ResourceProvider{
//			"vault": roleTestProvider.ResourceProvider(),
//		},
//		PreCheck:                  func() { util.TestAccPreCheck(t) },
//		PreventPostDestroyRefresh: true,
//		CheckDestroy:              testAccADSecretBackendRoleCheckDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testADSecretBackendRoleConfig(path, bindDN, bindPass, url, role, serviceAccountName, ttl),
//				Check: resource.ComposeTestCheckFunc(
//					resource.TestCheckResourceAttr("vault_ad_secret_backend_role.test", "role", role),
//					resource.TestCheckResourceAttr("vault_ad_secret_backend_role.test", "service_account_name", serviceAccountName),
//					resource.TestCheckResourceAttr("vault_ad_secret_backend_role.test", "ttl", fmt.Sprintf("%d", ttl)),
//				),
//			},
//			{
//				ResourceName:      "vault_database_secret_backend_role.test",
//				ImportState:       true,
//				ImportStateVerify: true,
//			},
//		},
//	})
//}

func TestAccADSecretBackendRole_basic(t *testing.T) {
	path := acctest.RandomWithPrefix("tf-test-ad")
	bindDN, bindPass, url := util.GetTestADCreds(t)

    t.Log(testADSecretBackendRoleConfig(path, bindDN, bindPass, url, "bob", "Bob", 60))
	t.Log(testADSecretBackendRoleUpdatedConfig("bob", "Bob", 120))

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"vault": roleTestProvider.ResourceProvider(),
		},
		PreCheck:                  func() { util.TestAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		CheckDestroy:              testAccADSecretBackendRoleCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testADSecretBackendRoleConfig(path, bindDN, bindPass, url, "bob", "Bob", 60),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vault_ad_secret_role.test", "role", "bob"),
					resource.TestCheckResourceAttr("vault_ad_secret_role.test", "service_account_name", "Bob"),
					resource.TestCheckResourceAttr("vault_ad_secret_role.test", "ttl", "60"),
				),
			},
			{
				Config: testADSecretBackendRoleUpdatedConfig("bob", "Bob", 120),
				//Check: resource.ComposeTestCheckFunc(
				//	resource.TestCheckResourceAttr("vault_ad_secret_role.test", "role", "bob"),
				//	resource.TestCheckResourceAttr("vault_ad_secret_role.test", "service_account_name", "Bob"),
				//	resource.TestCheckResourceAttr("vault_ad_secret_role.test", "ttl", "120"),
				//),
			},
		},
	})
}

func testAccADSecretBackendRoleCheckDestroy(s *terraform.State) error {
	client := roleTestProvider.SchemaProvider().Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vault_ad_secret_role" {
			continue
		}
		secret, err := client.Logical().Read(rs.Primary.ID)
		if err != nil {
			return err
		}
		if secret != nil {
			return fmt.Errorf("role %q still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testADSecretBackendRoleConfig(path, bindDN, bindPass, url, role, serviceAccountName string, ttl int) string {
	return fmt.Sprintf(`
resource "vault_ad_secret_backend" "test" {
	path = "%s"
	description = "test description"
	default_lease_ttl_seconds = "3600"
	max_lease_ttl_seconds = "7200"
	binddn = "%s"
	bindpass = "%s"
	url = "%s"
	insecure_tls = "true"
	userdn = "CN=Users,DC=corp,DC=example,DC=net"
}

resource "vault_ad_secret_role" "test" {
    path = "${vault_ad_secret_backend.test.path}"
    role = "%s"
    service_account_name = "%s"
    ttl = %d
}
`, path, bindDN, bindPass, url, role, serviceAccountName, ttl)
}

func testADSecretBackendRoleUpdatedConfig(role, serviceAccountName string, ttl int) string {
	return fmt.Sprintf(`
resource "vault_ad_secret_role" "test" {
    path = "${vault_ad_secret_backend.test.path}"
    role = "%s"
    service_account_name = "%s"
    ttl = %d
}
`, role, serviceAccountName, ttl)
}


