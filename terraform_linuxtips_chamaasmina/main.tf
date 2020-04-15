provider "aws" {
  region  = "us-east-1"
  version = "~> 2.0"
}

terraform {
  backend "s3" {
    // trocar o nome do bucket para um criado na sua conta aws
    bucket = "iaasweek-terraform-barbara"
    key    = "terraform-test.tfstate"
    region = "us-east-1"
  }
}