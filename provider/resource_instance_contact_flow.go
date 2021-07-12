package provider

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInstanceContactFlow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInstanceContactFlowCreate,
		ReadContext:   resourceInstanceContactFlowRead,
		UpdateContext: resourceInstanceContactFlowUpdate,
		DeleteContext: resourceInstanceContactFlowDelete,

		Schema: map[string]*schema.Schema{
			"name":        {Type: schema.TypeString, Required: true},
			"type":        {Type: schema.TypeString, Required: true},
			"description": {Type: schema.TypeString, Required: true},
			"content":     {Type: schema.TypeString, Required: true},
			"instance_id": {Type: schema.TypeString, Required: true},
			"arn":         {Type: schema.TypeString, Computed: true, Optional: true},
		},
	}
}

func resourceInstanceContactFlowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	params := &connect.CreateContactFlowInput{
		InstanceId:  aws.String(d.Get("instance_id").(string)),
		Name:        aws.String(d.Get("name").(string)),
		Type:        aws.String(d.Get("type").(string)),
		Description: aws.String(d.Get("description").(string)),
		Content:     aws.String(d.Get("content").(string)),
	}

	resp, err := connectSvc.CreateContactFlow(params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aws.StringValue(resp.ContactFlowId))
	d.Set("arn", aws.StringValue(resp.ContactFlowArn))

	return diags
}

func resourceInstanceContactFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// connectSvc := meta.(Client).ConnectClient

	// instanceID := d.Get("instance_id").(string)

	// params := &connect.DescribeInstanceInput{
	// 	InstanceId: aws.String(instanceID),
	// }
	// resp, err := connectSvc.DescribeInstance(params)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	// d.SetId(instanceID)
	// d.Set("arn", aws.StringValue(resp.Instance.Arn))
	// d.Set("instance_alias", aws.StringValue(resp.Instance.InstanceAlias))
	// d.Set("identity_management_type", aws.StringValue(resp.Instance.IdentityManagementType))
	// d.Set("inbound_calls_enabled", aws.BoolValue(resp.Instance.InboundCallsEnabled))
	// d.Set("outbound_calls_enabled", aws.BoolValue(resp.Instance.OutboundCallsEnabled))

	return diags
}

func resourceInstanceContactFlowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	params := &connect.UpdateContactFlowContentInput{
		InstanceId:    aws.String(d.Get("instance_id").(string)),
		Content:       aws.String(d.Get("content").(string)),
		ContactFlowId: aws.String(d.Id()),
	}

	_, err := connectSvc.UpdateContactFlowContent(params)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId(aws.StringValue(resp.ContactFlowId))
	// d.Set("arn", aws.StringValue(resp.ContactFlowArn))

	return diags
}

func resourceInstanceContactFlowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	// connectSvc := meta.(Client).ConnectClient

	// params := &connect.DisassociateLexBotInput{
	// 	InstanceId: aws.String(d.Get("instance_id").(string)),
	// 	LexRegion:  aws.String(d.Get("lex_bot_region").(string)),
	// 	BotName:    aws.String(d.Get("lex_bot_name").(string)),
	// }

	// _, err := connectSvc.DisassociateLexBot(params)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	d.SetId("")
	// // d.Set("instance_id", aws.StringValue(resp.Id))
	// // d.Set("arn", aws.StringValue(resp.Arn))

	// // // resourceInstanceContactFlowRead(ctx, d, m)

	return diags
}
