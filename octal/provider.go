package octal

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			//"linkerd2": linkerd2(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"linkerd2": linkerd2(),
		},
	}

	return provider
}
