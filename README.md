# Lockpick

Lockpick is a tool designed to use with concourse to pull [git-crypt](https://github.com/AGWA/git-crypt) key from [Vault](https://www.vaultproject.io/).

It shall be using [App ID backend](https://www.vaultproject.io/docs/auth/app-id.html) to store key and lockpick will aid in retrieving key from Vault.

## Configurations

By default the key is assume to be in `~/.lockpick`. This can be overwritten with `LOCKPICK_CONF`

```yaml
app_id: foobar
user_id: myuserid
key: mykeyid
output_file: ~/mygitcrypt.key
vault_address: "https://127.0.0.1:8200"
```


## Protip

The key need to be base64 before putting into vault.