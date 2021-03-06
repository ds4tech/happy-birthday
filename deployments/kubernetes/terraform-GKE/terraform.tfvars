#--------------------------------------------
# REQUIRED: Google account credentials
#--------------------------------------------

account_file_path = "service_account.json"
project  = "golang-168812"


#--------------------------------------------------------------------------------------
# Un-comment and provide your preferred values if you wish to overwrite the defaults
#--------------------------------------------------------------------------------------

gcloud-region   = "europe-west4"
gcloud-zone     = "europe-west4-a"
cluster_name    = "k8s-test"
gcp_cluster_count = 2 #2 for consul 3 for vault is required
machine_type= "n1-standard-2"
#machine-type= "g1-small"
# linux_admin_username = "ubuntu"
# linux_admin_password = "vault-agent-test-with-k8s"
