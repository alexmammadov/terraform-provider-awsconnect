package provider

import (
	"context"

	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"awsconnect_instance":              resourceInstance(),
			"awsconnect_instance_lex_bot":      resourceInstanceLexBot(),
			"awsconnect_instance_contact_flow": resourceInstanceContactFlow(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"awsconnect_instance": dataSourceInstance(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// Client -
type Client struct {
	ConnectClient *connect.Connect
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c := Client{
		ConnectClient: connectService(),
	}

	return c, diags
}
