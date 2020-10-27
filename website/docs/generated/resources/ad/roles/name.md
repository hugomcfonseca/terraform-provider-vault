---
layout: "vault"
page_title: "Vault: <TODO>"
sidebar_current: "<TODO>"
description: |-
  <TODO>
---

# <TODO>

<TODO>

## Example Usage

<TODO - this and HCL example below>
```hcl
resource "vault_jwt_auth_backend" "example" {
    description  = "Demonstration of the Terraform JWT auth backend"
    path = "jwt"
    oidc_discovery_url = "https://myco.auth0.com/"
    bound_issuer = "https://myco.auth0.com/"
}
```

## Argument Reference

The following arguments are supported:
* `path` - (Required) Path to where the back-end is mounted within Vault.
* `name` - (Required) Name of the role
* `service_account_name` - (Optional) The username/logon name for the service account with which this role will be associated.
* `ttl` - (Optional) In seconds, the default password time-to-live.
