# This block defines a Google Compute Engine instance resource with the name "vm2-from-terraform"

resource "google_compute_instance" "frontend-vm-from-terraform" {
  name         = "frontend-vm-from-terraform"
  machine_type = "e2-medium"
  zone         = "australia-southeast1-a"

  # Block that defines the boot disk (OS) for the VM
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11" # Specifies the boot image
      labels = {
        my_label = "value" # Labels applied to the boot disk
      }
    }
  }

  # Block that defines the network interface (how network is configurated) for the VM
  network_interface {
    network = "default" # Specifies the network to which the VM is connected.

    // gives vm instance a public IP
    access_config {

    }
  }

  //Specifies a startup script for vm
  metadata_startup_script = <<-EOF
#!/bin/bash
if [ ! -f /var/run/my_script_ran_before ]; then
    # Mark that the script has run before
    sudo touch /var/run/my_script_ran_before

    # Install the Ops Agent
    curl -sSO https://dl.google.com/cloudagents/add-google-cloud-ops-agent-repo.sh
    sudo bash add-google-cloud-ops-agent-repo.sh --also-install
    sudo apt-get update
    sudo apt-get -y install ops-agent

    # Execute the desired script
    cd /
    apt-get -y update
    apt-get -y install pip
    apt-get -y install git
    apt-get -y install npm
    cd /
    git clone https://github.com/yilong100/GoApp.git
    cd /
    cd GoApp/frontend/react-app/
    npm install
    npm start
fi
EOF

  # Make the second instance depend on the first one
  depends_on = [google_compute_instance.backend-vm-from-terraform]

}

# Create a firewall rule
resource "google_compute_firewall" "allow_ports_8080" {
  name    = "allow-ports-8080"
  network = "default" # Replace with your network name if not using the default network

  # Specify the rules for allowing traffic
  allow {
    protocol = "tcp"
    ports    = ["8080"]
  }

  source_ranges = ["0.0.0.0/0"] # Allow traffic from all IP addresses
}

# Output the assigned IP address
output "frontend-ip-address" {
  value = google_compute_instance.frontend-vm-from-terraform.network_interface[0].access_config[0].nat_ip
}
