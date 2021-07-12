terraform {
  required_providers {
    awsconnect = {
      version = "~> 0.1"
      source  = "alexmammadov/aws/awsconnect"
    }
  }
}

provider "awsconnect" {
}

# data "awsconnect_instance" "first" {
#   instance_id = "ebdfba29-ca78-41ea-b980-b328ab640a71"
# }

# output "first_order" {
#   value = data.awsconnect_instance.first
# }

resource "awsconnect_instance" "second" {
  instance_alias           = "sample-instance-2"
  identity_management_type = "CONNECT_MANAGED"
  inbound_calls_enabled    = true
  outbound_calls_enabled   = true
}

output "second_order" {
  value = awsconnect_instance.second
}

resource "awsconnect_instance_lex_bot" "second" {
  instance_id    = awsconnect_instance.second.instance_id
  lex_bot_region = "us-east-1"
  lex_bot_name   = "LexBotForConnect"
}

# resource "awsconnect_instance_contact_flow" "second" {
#   instance_id = awsconnect_instance.second.instance_id
#   name        = "LexBotConnect"
#   type        = "CONTACT_FLOW"
#   description = "Contact flow description"
#   content     = file("contact_flow.json")
# }