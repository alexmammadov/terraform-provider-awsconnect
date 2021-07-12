package provider

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInstanceLexBot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInstanceLexBotCreate,
		ReadContext:   resourceInstanceLexBotRead,
		// UpdateContext: resourceInstanceLexBotUpdate, not supported
		DeleteContext: resourceInstanceLexBotDelete,

		Schema: map[string]*schema.Schema{
			"instance_id":    {Type: schema.TypeString, Required: true, ForceNew: true},
			"lex_bot_region": {Type: schema.TypeString, Required: true, ForceNew: true},
			"lex_bot_name":   {Type: schema.TypeString, Required: true, ForceNew: true},
			"arn":            {Type: schema.TypeString, Computed: true, Optional: true},
		},
	}
}

func resourceInstanceLexBotCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	params := &connect.AssociateLexBotInput{
		InstanceId: aws.String(d.Get("instance_id").(string)),
		LexBot: &connect.LexBot{
			LexRegion: aws.String(d.Get("lex_bot_region").(string)),
			Name:      aws.String(d.Get("lex_bot_name").(string)),
		},
	}

	_, err := connectSvc.AssociateLexBot(params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(d.Get("instance_id").(string))
	// d.Set("instance_id", aws.StringValue())
	// d.Set("arn", aws.StringValue(resp.Arn))

	// // resourceInstanceLexBotRead(ctx, d, m)

	return diags
}

func resourceInstanceLexBotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	instanceID := d.Get("instance_id").(string)

	params := &connect.ListLexBotsInput{
		InstanceId: aws.String(instanceID),
	}

	resp, err := connectSvc.ListLexBots(params)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(resp.LexBots) == 0 {
		return diag.Errorf("lex bot not found for instance: %s", instanceID)
	}

	// d.SetId(aws.StringValue(resp.Id))
	// d.Set("instance_id", aws.StringValue(resp.Id))
	// d.Set("arn", aws.StringValue(resp.Arn))

	// // resourceInstanceLexBotRead(ctx, d, m)

	d.SetId(instanceID)
	d.Set("lex_bot_region", aws.StringValue(resp.LexBots[0].LexRegion))
	d.Set("lex_bot_name", aws.StringValue(resp.LexBots[0].Name))

	return diags
}

// func resourceInstanceLexBotUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics
// 	// connectSvc := meta.(Client).ConnectClient

// 	// instanceID := aws.String(d.Id())

// 	// if d.HasChange("inbound_calls_enabled") {
// 	// 	params := &connect.UpdateInstanceAttributeInput{
// 	// 		InstanceId:    instanceID,
// 	// 		AttributeType: aws.String("INBOUND_CALLS"),
// 	// 		Value:         aws.String(strconv.FormatBool(d.Get("inbound_calls_enabled").(bool))),
// 	// 	}
// 	// 	_, err := connectSvc.UpdateInstanceAttribute(params)
// 	// 	if err != nil {
// 	// 		return diag.FromErr(err)
// 	// 	}
// 	// }
// 	// if d.HasChange("outbound_calls_enabled") {
// 	// 	params := &connect.UpdateInstanceAttributeInput{
// 	// 		InstanceId:    instanceID,
// 	// 		AttributeType: aws.String("OUTBOUND_CALLS"),
// 	// 		Value:         aws.String(strconv.FormatBool(d.Get("outbound_calls_enabled").(bool))),
// 	// 	}
// 	// 	_, err := connectSvc.UpdateInstanceAttribute(params)
// 	// 	if err != nil {
// 	// 		return diag.FromErr(err)
// 	// 	}
// 	// }

// 	return diags

func resourceInstanceLexBotDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	connectSvc := meta.(Client).ConnectClient

	params := &connect.DisassociateLexBotInput{
		InstanceId: aws.String(d.Get("instance_id").(string)),
		LexRegion:  aws.String(d.Get("lex_bot_region").(string)),
		BotName:    aws.String(d.Get("lex_bot_name").(string)),
	}

	_, err := connectSvc.DisassociateLexBot(params)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId(aws.StringValue(resp.Id))
	// d.Set("instance_id", aws.StringValue(resp.Id))
	// d.Set("arn", aws.StringValue(resp.Arn))

	// // resourceInstanceLexBotRead(ctx, d, m)

	return diags
}
