To start Postgres instance go to terraform/createVM folder and run:

```
terraform apply -target=google_sql_database_instance.postgres_goApp
```

Once the instance is fulling running go to:

    GCP Console -> SQL -> go-app-db -> Users (on the sidebar) -> 
    Create User: 
    Username:   goApp 
    Password:   goApp1234

Go back to Overview, select 'Open Cloud Shell'

![Alt text](image.png)


Once Connected, enter the password you assigned to the user

![Alt text](image-1.png)

Now to create the database type:

    CREATE DATABASE users

Check that database is created by typing:

    \l

And use arrow keys to scroll down.

Go to Connections:

![Alt text](image-2.png)

GO to Networking, scroll down and ADD A NETWORK with the following settings (or the ip address of the backend):

![Alt text](image-3.png)
