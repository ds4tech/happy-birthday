terraform {
  backend "gcs" {
    bucket  = "golang-168812.appspot.com"
    prefix  = "terraform/state"
    credentials = "service_account.json"
  }
}