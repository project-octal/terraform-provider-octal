terraform {
  required_providers {
    octal = {
      source = "project-octal/octal"
      version = "0.0.2"
    }
  }
}

provider "octal" {
  # Configuration options
}

#data "linkerd2" "all" {}
