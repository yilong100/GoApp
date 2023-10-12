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

  # allow_stopping_for_update = true # This line is commented out but can be used to allow VM stopping during updates
}
