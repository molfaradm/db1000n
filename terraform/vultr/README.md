# Vultr deployment

## Requirements

- Vultr account
- API token (Go to API - Personal access tokens) and generate Personal access token (with write permissions)
- `terraform` (1.0+) installed

## Deploy

To deploy run:

Change API token in the  provider.tf file


```sh
terraform init
terraform plan -var "pvt_key=$HOME/.ssh/id_rsa_key"
terraform apply -var "pvt_key=$HOME/.ssh/id_rsa_key"   -auto-approve
```


## Destroy

To destroy infrastructure use commands:

```sh
terraform destroy -var "pvt_key=$HOME/.ssh/id_rsa_key"
```
