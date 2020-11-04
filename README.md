# Terraform Provider Octal
Terraform provider for deploying and managing Project-Octal on a Kubernetes cluster.

Run the following command to build the provider

```shell
go build -o terraform-provider-octal
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```