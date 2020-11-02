package ad

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/vault/api"
	"github.com/terraform-providers/terraform-provider-vault/schema"
	"github.com/terraform-providers/terraform-provider-vault/util"
	"github.com/terraform-providers/terraform-provider-vault/vault"
)

var configTestProvider = func() *schema.Provider {
	p := schema.NewProvider(vault.Provider())
	p.RegisterResource("vault_mount", vault.MountResource())
	p.RegisterResource("vault_ad_secret_backend", ConfigResource())
	return p
}()

func TestADSecretBackend(t *testing.T) {
	path := acctest.RandomWithPrefix("tf-test-ad")
	bindDN, bindPass, url := util.GetTestADCreds(t)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"vault": configTestProvider.ResourceProvider(),
		},
		PreCheck:                  func() { util.TestAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		CheckDestroy:              testAccADSecretBackendCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testADSecretBackend_initialConfig(path, bindDN, bindPass, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "path", path),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "description", "test description"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "default_lease_ttl_seconds", "3600"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "max_lease_ttl_seconds", "7200"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "binddn", bindDN),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "bindpass", bindPass),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "url", url),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "insecure_tls", "true"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "userdn", "CN=Users,DC=corp,DC=example,DC=net"),
				),
			},
			{
				Config: testADSecretBackend_updateConfig(path, bindDN, bindPass, url),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "path", path),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "description", "test description"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "default_lease_ttl_seconds", "7200"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "max_lease_ttl_seconds", "14400"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "binddn", bindDN),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "bindpass", bindPass),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "url", url),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "insecure_tls", "false"),
					resource.TestCheckResourceAttr("vault_ad_secret_backend.test", "userdn", "CN=Users,DC=corp,DC=hashicorp,DC=com"),
				),
			},
		},
	})
}

func testAccADSecretBackendCheckDestroy(s *terraform.State) error {
	client := configTestProvider.SchemaProvider().Meta().(*api.Client)

	mounts, err := client.Sys().ListMounts()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "vault_ad_secret_backend" {
			continue
		}
		for path, mount := range mounts {
			path = strings.Trim(path, "/")
			rsPath := strings.Trim(rs.Primary.Attributes["path"], "/")
			if mount.Type == "ad" && path == rsPath {
				return fmt.Errorf("Mount %q still exists", path)
			}
		}
	}
	return nil
}

func testADSecretBackend_initialConfig(path, bindDN, bindPass, url string) string {
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
}`, path, bindDN, bindPass, url)
}

func testADSecretBackend_updateConfig(path, bindDN, bindPass, url string) string {
	return fmt.Sprintf(`
resource "vault_ad_secret_backend" "test" {
	path = "%s"
	description = "test description"
	default_lease_ttl_seconds = "7200"
	max_lease_ttl_seconds = "14400"
	binddn = "%s"
	bindpass = "%s"
	url = "%s"
	insecure_tls = "false"
	userdn = "CN=Users,DC=corp,DC=hashicorp,DC=com"
}`, path, bindDN, bindPass, url)
}
