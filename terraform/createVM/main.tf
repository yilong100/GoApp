# This block defines a Google Compute Engine instance resource with the name "vm2-from-terraform"

resource "google_compute_instance" "vm-from-terraform" {
  name         = "vm2-from-terraform"
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
  }

  metadata_startup_script = "echo hi > /test.txt" # Specifies a startup script for the VM
  # allow_stopping_for_update = true # This line is commented out but can be used to allow VM stopping during updates
}
