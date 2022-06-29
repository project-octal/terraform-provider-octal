package octal

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func certManagerSchema() map[string]*schema.Schema {
	thisSchema := map[string]*schema.Schema{
		"namespace": octalNamespaceSpecSchema("cert-manager"),

		"controller": octalDeploySpecSchema(
			"cert-manager-cainjector",
			"v1.8.1",
			"jetstack/cert-manager-controller"),

		"cainjector": octalDeploySpecSchema(
			"cert-manager-cainjector",
			"v1.8.1",
			"jetstack/cert-manager-cainjector"),

		"webhook": octalDeploySpecSchema(
			"cert-manager-cainjector",
			"v1.8.1",
			"jetstack/cert-manager-webhook"),

		"custom_resources": octalCustomResourceSpecSchema("cert-manager"),
	}
	return thisSchema
}
