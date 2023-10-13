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
  project     = "playground-s-11-62fe5fd8" # Your Google Cloud Project ID
  region      = "australia-southeast1"     # The desired region for resources
  zone        = "australia-southeast1-a"   # The specific zone within the region
  credentials = "./keys.json"              # Path to your service account credentials file
}
