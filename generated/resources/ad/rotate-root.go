package ad

// DO NOT EDIT
// This code is generated.

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/vault/api"
)

func RotateRootResource() *schema.Resource {
	fields := map[string]*schema.Schema{
		"path": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: `The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".`,
			StateFunc: func(v interface{}) string {
				return strings.Trim(v.(string), "/")
			},
		},
	}
	return &schema.Resource{
		Update: updateRotateRootResource,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: fields,
	}
}
func updateRotateRootResource(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	path := fmt.Sprintf("%s/rotate-root", d.Get("path").(string))

	log.Printf("[DEBUG] Rotating root credentials at %q", path)

	data := make(map[string]interface{})
	_, err := client.Logical().Write(path, data)
	if err != nil {
		return fmt.Errorf("error rotating root credentials %q: %s", path, err)
	}
	log.Printf("[DEBUG] Rotated root credentials %q", path)

	return nil
}
