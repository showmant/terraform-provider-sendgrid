terraform {
  required_providers {
    sendgrid = {
      source = "registry.terraform.io/showmant/sendgrid"
    }
  }
}

provider "sendgrid" {
}
