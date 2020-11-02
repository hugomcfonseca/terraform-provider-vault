package generated

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	adCreds "github.com/terraform-providers/terraform-provider-vault/generated/datasources/ad/creds"
	"github.com/terraform-providers/terraform-provider-vault/generated/datasources/transform/decode"
	"github.com/terraform-providers/terraform-provider-vault/generated/datasources/transform/encode"
	"github.com/terraform-providers/terraform-provider-vault/generated/resources/ad"
	adRoles "github.com/terraform-providers/terraform-provider-vault/generated/resources/ad/roles"
	"github.com/terraform-providers/terraform-provider-vault/generated/resources/transform/alphabet"
	"github.com/terraform-providers/terraform-provider-vault/generated/resources/transform/role"
	"github.com/terraform-providers/terraform-provider-vault/generated/resources/transform/template"
	"github.com/terraform-providers/terraform-provider-vault/generated/resources/transform/transformation"
)

// Please alphabetize.
var DataSourceRegistry = map[string]*schema.Resource{
	"vault_ad_secret_role_creds": adCreds.CredsDataSource(),
	"vault_transform_encode":     encode.RoleNameDataSource(),
	"vault_transform_decode":     decode.RoleNameDataSource(),
}

// Please alphabetize.
var ResourceRegistry = map[string]*schema.Resource{
	"vault_ad_secret_backend":        ad.ConfigResource(),
	"vault_ad_secret_role":           adRoles.RoleResource(),
	"vault_transform_alphabet":       alphabet.NameResource(),
	"vault_transform_role":           role.NameResource(),
	"vault_transform_template":       template.NameResource(),
	"vault_transform_transformation": transformation.NameResource(),
}
