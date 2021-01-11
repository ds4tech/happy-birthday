provider "google" {
  credentials = file(var.account_file_path)
  project     = var.project
  region      = var.gcloud-region
}

resource "google_container_cluster" "primary" {
  name               = var.cluster_name
  location           = var.gcloud-zone
  initial_node_count = var.gcp_cluster_count

  master_auth {
    username = var.linux_admin_username
    password = var.linux_admin_password

    client_certificate_config {
      issue_client_certificate = false
    }
  }
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "my-node-pool"
  location   = var.gcloud-zone
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = var.machine_type
    tags = ["dev", "work"]

    labels = {
      this-is-for = "dev-cluster"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]

  }
}

#--------------------------------------------
# Outputs
#--------------------------------------------
output "gcp_cluster_endpoint" {
  value = google_container_cluster.primary.endpoint
}

output "gcp_cluster_name" {
  value = google_container_cluster.primary.name
}

output "gcp_ssh_command" {
  value = "ssh ${var.linux_admin_username}@${google_container_cluster.primary.endpoint}"
}

output "gcp_zone" {
  value = google_container_cluster.primary.location
}

output "gcp_project" {
  value = var.project
}
