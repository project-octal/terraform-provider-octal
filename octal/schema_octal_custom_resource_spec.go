package octal

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func octalCustomResourceSpecSchema(componentName string) *schema.Schema {
	componentSpec := &schema.Schema{
		Type:        schema.TypeMap,
		Description: fmt.Sprintf("Custom Resource Definition configuration for %s", componentName),
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
	return buildSchema(componentName, "deployment", componentSpec)
}
