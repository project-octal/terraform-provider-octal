package octal

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func octalNamespaceSpecSchema(componentName string) *schema.Schema {
	return buildSchema(componentName, "namespace", nil)
}
