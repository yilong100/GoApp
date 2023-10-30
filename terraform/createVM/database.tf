resource "google_sql_database_instance" "postgres_goApp" {
    name             = "go-app-db"
    database_version = "POSTGRES_13"
    region           = "australia-southeast1"
    settings {
        tier = "db-f1-micro"

        # ip_configuration {
        #     authorized_networks {
        #         name            = "my-subnet"
        #         value           = "10.152.0.0/20"
        #     }
        # }
    }

    deletion_protection = false

}

resource "google_sql_user" "user" {

    depends_on = [google_sql_database_instance.postgres_goApp]

    name     = "goApp"
    instance = google_sql_database_instance.postgres_goApp.name
    password = "goApp1234"
}

# Output the assigned IP address
output "database-ip-address" {
  value = google_sql_database_instance.postgres_goApp.ip_address[0].ip_address
}

resource "null_resource" "save_database_ip_to_file" {
    triggers = {
        instance_id = google_sql_database_instance.postgres_goApp.id
    }

    depends_on = [google_sql_database_instance.postgres_goApp]

    provisioner "local-exec" {
        command = "echo ${google_sql_database_instance.postgres_goApp.ip_address[0].ip_address} > ../../backend/db/database-ip-address.txt"
    }

}