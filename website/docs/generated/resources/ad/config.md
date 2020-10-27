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
* `anonymous_group_search` - (Optional) Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test).
* `binddn` - (Optional) LDAP DN for searching for the user DN (optional)
* `bindpass` - (Optional) LDAP password for searching for the user DN (optional)
* `case_sensitive_names` - (Optional) If true, case sensitivity will be used when comparing usernames and groups for matching policies.
* `certificate` - (Optional) CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded (optional)
* `client_tls_cert` - (Optional) Client certificate to provide to the LDAP server, must be x509 PEM encoded (optional)
* `client_tls_key` - (Optional) Client certificate key to provide to the LDAP server, must be x509 PEM encoded (optional)
* `deny_null_bind` - (Optional) Denies an unauthenticated LDAP bind request if the user's password is empty; defaults to true
* `discoverdn` - (Optional) Use anonymous bind to discover the bind DN of a user (optional)
* `formatter` - (Optional) Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix".
* `groupattr` - (Optional) LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate user group membership. Examples: "cn" or "memberOf", etc. Default: cn
* `groupdn` - (Optional) LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org)
* `groupfilter` - (Optional) Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username Example: (&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}})) Default: (|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))
* `insecure_tls` - (Optional) Skip LDAP server SSL Certificate verification - VERY insecure (optional)
* `last_rotation_tolerance` - (Optional) The number of seconds after a Vault rotation where, if Active Directory shows a later rotation, it should be considered out-of-band.
* `length` - (Optional) The desired length of passwords that Vault generates.
* `max_ttl` - (Optional) In seconds, the maximum password time-to-live.
* `request_timeout` - (Optional) Timeout, in seconds, for the connection when making requests against the server before returning back an error.
* `starttls` - (Optional) Issue a StartTLS command after establishing unencrypted connection (optional)
* `tls_max_version` - (Optional) Maximum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'
* `tls_min_version` - (Optional) Minimum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'
* `ttl` - (Optional) In seconds, the default password time-to-live.
* `upndomain` - (Optional) Enables userPrincipalDomain login with [username]@UPNDomain (optional)
* `url` - (Optional) LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order.
* `use_pre111_group_cn_behavior` - (Optional) In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations.
* `use_token_groups` - (Optional) If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones.
* `userattr` - (Optional) Attribute used for users (default: cn)
* `userdn` - (Optional) LDAP domain to use for users (eg: ou=People,dc=example,dc=org)
