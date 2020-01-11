resource "aws_s3_bucket" "terraform-state" {
  bucket = "terraform-state-mat"
  acl    = "private"

  tags = {
    Name = "Terraform state"
  }
}

