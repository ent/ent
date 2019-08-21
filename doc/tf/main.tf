terraform {
  backend "s3" {
    bucket         = "entgo.tfstate"
    region         = "eu-central-1"
    key            = "terraform.tfstate"
    dynamodb_table = "entgo.terraform.lock"
  }

  required_version = "> 0.12"
}

data "terraform_remote_state" "current" {
  backend = "s3"

  config = {
    bucket         = "entgo.tfstate"
    region         = "eu-central-1"
    key            = "terraform.tfstate"
    dynamodb_table = "entgo.terraform.lock"
  }

  defaults = {
    allowed_cidrs = ["0.0.0.0/0"]
  }
}

provider "aws" {
  region  = "eu-central-1"
  version = "~> 2.0"
}

locals {
  name        = "entgo"
  domain_name = "entgo.io"
}
