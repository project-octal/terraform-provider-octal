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
		Schema:        map[string]*schema.Schema{},
	}
}

func linkerd2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	yml_output, err := cmd.InstallLinkerd2()
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	} else {
		fmt.Println(yml_output)
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

	return diags
}
