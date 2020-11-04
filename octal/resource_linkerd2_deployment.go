package octal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func linkerd2Deployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: linkerd2DeploymentCreate,
		ReadContext:   linkerd2DeploymentRead,
		UpdateContext: linkerd2DeploymentUpdate,
		DeleteContext: linkerd2DeploymentDelete,
		Schema:        map[string]*schema.Schema{},
	}
}

func linkerd2DeploymentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func linkerd2DeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func linkerd2DeploymentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return linkerd2DeploymentRead(ctx, d, m)
}

func linkerd2DeploymentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
