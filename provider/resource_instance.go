package provider

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInstanceCreate,
		ReadContext:   resourceInstanceRead,
		UpdateContext: resourceInstanceUpdate,
		DeleteContext: resourceInstanceDelete,

		Schema: map[string]*schema.Schema{
			"instance_alias":           {Type: schema.TypeString, Required: true},
			"identity_management_type": {Type: schema.TypeString, Required: true},
			"inbound_calls_enabled":    {Type: schema.TypeBool, Required: true},
			"outbound_calls_enabled":   {Type: schema.TypeBool, Required: true},
			"instance_id":              {Type: schema.TypeString, Computed: true, Optional: true},
			"arn":                      {Type: schema.TypeString, Computed: true, Optional: true},
		},
	}
}

func resourceInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	params := &connect.CreateInstanceInput{
		InstanceAlias:          aws.String(d.Get("instance_alias").(string)),
		IdentityManagementType: aws.String(d.Get("identity_management_type").(string)),
		InboundCallsEnabled:    aws.Bool(d.Get("inbound_calls_enabled").(bool)),
		OutboundCallsEnabled:   aws.Bool(d.Get("outbound_calls_enabled").(bool)),
	}

	resp, err := connectSvc.CreateInstance(params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aws.StringValue(resp.Id))
	d.Set("instance_id", aws.StringValue(resp.Id))
	d.Set("arn", aws.StringValue(resp.Arn))

	// resourceInstanceRead(ctx, d, m)

	return diags
}

func resourceInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	instanceID := aws.String(d.Id())

	if d.HasChange("inbound_calls_enabled") {
		params := &connect.UpdateInstanceAttributeInput{
			InstanceId:    instanceID,
			AttributeType: aws.String("INBOUND_CALLS"),
			Value:         aws.String(strconv.FormatBool(d.Get("inbound_calls_enabled").(bool))),
		}
		_, err := connectSvc.UpdateInstanceAttribute(params)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("outbound_calls_enabled") {
		params := &connect.UpdateInstanceAttributeInput{
			InstanceId:    instanceID,
			AttributeType: aws.String("OUTBOUND_CALLS"),
			Value:         aws.String(strconv.FormatBool(d.Get("outbound_calls_enabled").(bool))),
		}
		_, err := connectSvc.UpdateInstanceAttribute(params)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	params := &connect.DeleteInstanceInput{
		InstanceId: aws.String(d.Id()),
	}

	_, err := connectSvc.DeleteInstance(params)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
