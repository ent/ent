terraform {
  backend "s3" {
    bucket         = "entgo.tfstate"
    region         = "eu-central-1"
    key            = "terraform.tfstate"
    dynamodb_table = "entgo.terraform.lock"
  }
}
