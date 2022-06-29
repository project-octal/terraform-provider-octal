package octal

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func buildSchema(componentName, componentType string, componentSchema *schema.Schema) *schema.Schema {
	schemaSpec := map[string]*schema.Schema{
		"uid": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("The unique in time and space value for this %s. More info: http://kubernetes.io/docs/user-guide/identifiers#uids", componentType),
			Computed:    true,
		},
		"resource_version": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("An opaque value that represents the internal version of this %s that can be used by clients to determine when %s has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency", componentType, componentName),
			Computed:    true,
		},
		"apiVersion": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("The apiVersion used by this %s", componentType),
			Computed:    true,
		},
		"kind": {
			Type:        schema.TypeString,
			Description: fmt.Sprintf("The kind of object"),
			Computed:    true,
		},
		"metadata": {
			Type:        schema.TypeMap,
			Description: fmt.Sprintf("Metadata for %s", componentName),
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "A name that will be given to the deployment",
						Default:     componentName,
					},
					"labels": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "Additional labels to add to the deployment",
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"annotations": {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "Additional annotations to add to the deployment",
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
	}

	if componentSchema != nil {
		schemaSpec["spec"] = componentSchema
	}

	return &schema.Schema{
		Type:        schema.TypeMap,
		Description: fmt.Sprintf("Deployment configuration for %s", componentName),
		Elem: &schema.Resource{
			Schema: schemaSpec,
		},
	}
}
