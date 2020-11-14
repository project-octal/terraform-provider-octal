terraform {
  required_providers {
    octal = {
      source = "project-octal/octal"
      version = "0.0.4"
    }
  }
}

provider "octal" {
  # Configuration options
}

data "octal_linkerd2" "all" {}
