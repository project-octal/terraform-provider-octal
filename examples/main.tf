terraform {
  required_providers {
    octal = {
      source = "hashicorp.com/edu/octal"
    }
  }
}


provider "octal" {}

data "linkerd2" "all" {}







#module "psl" {
#  source = "./coffee"
#
#  coffee_name = "Packer Spiced Latte"
#}
#
#output "psl" {
#  value = module.psl.coffee
#}