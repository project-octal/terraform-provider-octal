package octal

import (
	"context"
	"fmt"
	"github.com/project-octal/linkderd-cli/cmd"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func linkerd2() *schema.Resource {
	return &schema.Resource{
		CreateContext: linkerd2Create,
		ReadContext:   linkerd2Read,
		UpdateContext: linkerd2Update,
		DeleteContext: linkerd2Delete,
		Schema: linkderd2Schema(),
	}
}

func linkerd2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	yml_output, options, err := cmd.InstallLinkerd2()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Linkerd Yaml output",
			Detail:   err.Error(),
		})
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Got Linkerd Yaml output",
			Detail:   yml_output,
		})
		fmt.Println(yml_output)
		translateInstallOptions(options, d)
	}

	return diags
}

func linkerd2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

func linkerd2Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return linkerd2Read(ctx, d, m)
}

func linkerd2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	err := cmd.UninstallLinkerd2()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Linkerd Yaml output",
			Detail:   err.Error(),
		})
	}

	return diags
}
