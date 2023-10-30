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
  project     = "rapid-chassis-402302"   # Your Google Cloud Project ID. James: rapid-chassis-402302 Jamal: impactful-post-402223 
  region      = "australia-southeast1"   # The desired region for resources
  zone        = "australia-southeast1-a" # The specific zone within the region
  credentials = "./keys.json"            # Path to your service account credentials file
}

# Learn our public IP address
data "http" "icanhazip" {
   url = "http://icanhazip.com"
}

# Create a firewall rule for frontend
resource "google_compute_firewall" "allow_ports_3000s" {
  name    = "allow-ports-3000"
  network = "default" # Replace with your network name if not using the default network

  # Specify the rules for allowing traffic
  allow {
    protocol = "tcp"
    ports    = ["3000"]
  }

  source_ranges = ["${chomp(data.http.icanhazip.body)}"] # Allow traffic from all IP addresses

  depends_on = [data.http.icanhazip]
}

# Create a firewall rule for backend
resource "google_compute_firewall" "allow_ports_8080" {
  name    = "allow-ports-8080"
  network = "default" # Replace with your network name if not using the default network

  # Specify the rules for allowing traffic
  allow {
    protocol = "tcp"
    ports    = ["8080"]
  }

  source_ranges = ["10.152.0.0/20"] # Allow traffic from all IP addresses
}

# Create a firewall rule for database
resource "google_compute_firewall" "allow_ports_5432" {
  name    = "allow-ports-5432"
  network = "default" # Replace with your network name if not using the default network

  # Specify the rules for allowing traffic
  allow {
    protocol = "tcp"
    ports    = ["5432"]
  }

  source_ranges = ["10.152.0.0/20"]

}

output "public_ip" {
  value = "${chomp(data.http.icanhazip.body)}"
}

