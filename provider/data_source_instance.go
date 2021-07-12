package provider

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInstanceRead,
		Schema: map[string]*schema.Schema{
			"instance_id":              {Type: schema.TypeString, Required: true},
			"instance_alias":           {Type: schema.TypeString, Computed: true},
			"identity_management_type": {Type: schema.TypeString, Computed: true},
			"inbound_calls_enabled":    {Type: schema.TypeBool, Computed: true},
			"outbound_calls_enabled":   {Type: schema.TypeBool, Computed: true},
			"arn":                      {Type: schema.TypeString, Computed: true},
		},
	}
}

func dataSourceInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	instanceID := d.Get("instance_id").(string)

	params := &connect.DescribeInstanceInput{
		InstanceId: aws.String(instanceID),
	}
	resp, err := connectSvc.DescribeInstance(params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(instanceID)
	d.Set("arn", aws.StringValue(resp.Instance.Arn))
	d.Set("instance_alias", aws.StringValue(resp.Instance.InstanceAlias))
	d.Set("identity_management_type", aws.StringValue(resp.Instance.IdentityManagementType))
	d.Set("inbound_calls_enabled", aws.BoolValue(resp.Instance.InboundCallsEnabled))
	d.Set("outbound_calls_enabled", aws.BoolValue(resp.Instance.OutboundCallsEnabled))

	return diags
}
