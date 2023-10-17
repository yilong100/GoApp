# This section specifies the required cloud provider and its source and version
# Terraform will ensure that the specified cloud provider is available for use

terraform {
  required_providers {
    google = {
      source  = "hashicorp/google" # Source for the Google Cloud provider
      version = "5.1.0"            # Version of the Google Cloud provider
    }
  }
}

# change everytime
# This section configures the Google Cloud provider for your infrastructure
provider "google" {
  project     = "impactful-post-402223"    # Your Google Cloud Project ID
  region      = "australia-southeast1"     # The desired region for resources
  zone        = "australia-southeast1-a"   # The specific zone within the region
  credentials = "./keys.json"              # Path to your service account credentials file
}

# Create a firewall rule
resource "google_compute_firewall" "allow_ports_3000" {
  name    = "allow-ports-3000"
  network = "default"  # Replace with your network name if not using the default network

  # Specify the rules for allowing traffic
  allow {
    protocol = "tcp"
    ports    = ["3000"]
  }

  source_ranges = ["0.0.0.0/0"]  # Allow traffic from all IP addresses
}

resource "null_resource" "push_to_gits" {
  triggers = {
    instance_id = google_compute_instance.frontend-vm-from-terraform.id
  }

  depends_on = [null_resource.save_frontend_ip_to_file]

  provisioner "local-exec" {
    command = <<-EOT
      cd D:\Documents\FDM\Up Skilling\GoApp | git add . | git commit -m "hi" | git push
    EOT
  }
}