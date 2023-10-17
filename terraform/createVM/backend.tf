# This block defines a Google Compute Engine instance resource with the name "vm2-from-terraform"

resource "google_compute_instance" "backend-vm-from-terraform" {
  name         = "backend-vm-from-terraform"
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
    wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
    mkdir /root/go
    export GOPATH=/root/go
    export GOCACHE=/root/go/cache
    apt-get -y update
    apt-get -y install pip
    apt-get -y install git
    apt-get -y install golang
    apt-get -y install gnupg curl
    curl -fsSL https://pgp.mongodb.com/server-7.0.asc | gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor
    echo "deb [ signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] http://repo.mongodb.org/apt/debian bullseye/mongodb-org/7.0 main" | tee /etc/apt/sources.list.d/mongodb-org-7.0.list
    apt-get -y update
    apt-get install -y mongodb-org
    echo "mongodb-org hold" | dpkg --set-selections
    echo "mongodb-org-database hold" | dpkg --set-selections
    echo "mongodb-org-server hold" | dpkg --set-selections
    echo "mongodb-mongosh hold" | dpkg --set-selections
    echo "mongodb-org-mongos hold" | dpkg --set-selections
    echo "mongodb-org-tools hold" | dpkg --set-selections
    systemctl start mongod
    systemctl daemon-reload
    systemctl enable mongod
    cd /
    git clone https://github.com/yilong100/GoApp.git
    cd /
    cd GoApp/backend/
    go mod init github.com/yilong100/GoApp.git
    go build
    ./GoPractice &
fi
EOF
  # allow_stopping_for_update = true # This line is commented out but can be used to allow VM stopping during updates
}

# Output the assigned IP address
output "backend-ip-address" {
  value = google_compute_instance.backend-vm-from-terraform.network_interface[0].access_config[0].nat_ip
}

# Use local-exec provisioner to save the IP address to a file
resource "null_resource" "save_backend_ip_to_file" {
  triggers = {
    instance_id = google_compute_instance.backend-vm-from-terraform.id
  }

  depends_on = [google_compute_instance.backend-vm-from-terraform]

  provisioner "local-exec" {
    command = <<-EOT
      printf 'const apiUrl = "${google_compute_instance.backend-vm-from-terraform.network_interface[0].access_config[0].nat_ip}"\nexport default apiUrl;' > ../../frontend/react-app/src/backend-ip-address.js
    EOT
  }
}