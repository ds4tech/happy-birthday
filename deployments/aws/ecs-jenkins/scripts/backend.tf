terraform {
  backend "s3" {
    bucket = "terraform-state-mat"
    key    = "terraform.tfstate"
    region = "eu-west-1"
  }
}
