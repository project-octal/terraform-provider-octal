package octal

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"time"
)

func resourceOctalCertManager() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOctalCertManagerCreate,
		ReadContext:   resourceOctalCertManagerRead,
		UpdateContext: resourceOctalCertManagerUpdate,
		DeleteContext: resourceOctalCertManagerDelete,
		Schema: map[string]*schema.Schema{
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
		},
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}

func resourceOctalCertManagerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return diag.FromErr(err)
	}

	/*********************
	CERT-MANAGER NAMESPACE
	**********************/
	namespaceObjectMeta := metav1.ObjectMeta{}
	namespaceMeta := d.Get("namespace").([]interface{})[0].(map[string]interface{})
	if v, ok := namespaceMeta["name"]; ok {
		namespaceObjectMeta.Name = v.(string)
	}
	if v, ok := namespaceMeta["annotations"].(map[string]interface{}); ok && len(v) > 0 {
		namespaceObjectMeta.Annotations = expandStringMap(namespaceMeta["annotations"].(map[string]interface{}))
	}
	if v, ok := namespaceMeta["labels"].(map[string]interface{}); ok && len(v) > 0 {
		namespaceObjectMeta.Labels = expandStringMap(namespaceMeta["labels"].(map[string]interface{}))
	}
	namespace := api.Namespace{
		ObjectMeta: namespaceObjectMeta,
	}
	log.Printf("[INFO] Creating new namespace: %#v", namespace)
	out, err := conn.CoreV1().Namespaces().Create(ctx, &namespace, metav1.CreateOptions{})
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO] Submitted new namespace: %#v", out)

	/*********************
	CERT-MANAGER DEPLOYMENT
	**********************/

	// This is where the ID of the deployment is set
	d.SetId(out.Name)
	return diags
}

func resourceOctalCertManagerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

func resourceOctalCertManagerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceOctalCertManagerRead(ctx, d, m)
}

func resourceOctalCertManagerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	conn, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return diag.FromErr(err)
	}

	err = conn.CoreV1().Namespaces().Delete(ctx, d.Id(), metav1.DeleteOptions{})

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Linkerd Yaml output",
			Detail:   err.Error(),
		})
	}

	return diags
}
