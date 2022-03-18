terraform {
  required_providers {
    vultr = {
      source = "vultr/vultr"
      version = "2.9.1"
    }
  }
}

# Configure the Vultr Provider
provider "vultr" {
  api_key = "CHANGE ME!!!!!"
  rate_limit = 700
  retry_limit = 3
}


variable "pvt_key" {}
