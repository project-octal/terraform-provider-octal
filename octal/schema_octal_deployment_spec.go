package octal

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func octalDeploySpecSchema(componentName string, imageTag string, imageName string) *schema.Schema {
	componentSpec := &schema.Schema{
		Type:        schema.TypeMap,
		Description: fmt.Sprintf("Deployment configuration for %s", componentName),
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{

				"image_tag": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The image tag used by the deployment",
					Default:     imageTag,
				},
				"image_name": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The image name used by the deployment",
					Default:     imageName,
				},
				"image_repository": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The image repository to use when pulling images",
				},
				"image_pull_policy": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Default:     "IfNotPresent",
					Description: "Determines when the image should be pulled prior to starting the container. `Always`: Always pull the image. | `IfNotPresent`: Only pull the image if it does not already exist on the node. | `Never`: Never pull the image",
				},
			},
		},
	}
	return buildSchema(componentName, "deployment", componentSpec)
}
