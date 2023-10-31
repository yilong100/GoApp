1. To start Postgres instance go to terraform/createVM folder and run:

```
terraform apply -target=google_sql_database_instance.postgres_goApp
```

2. Once the instance is fully running go to:

   GCP Console -> SQL -> go-app-db -> Users (on the sidebar) ->
   Change Password for postgres user

3. Go back to Overview, select 'Open Cloud Shell'

![Alt text](image.png)

4. Once Connected, enter the password you assigned to the user

![Alt text](image-1.png)

5. Now to create the database type:

<<<<<<< HEAD
   CREATE DATABASE users;
=======
    CREATE DATABASE users;
>>>>>>> 4c196d3349118033e29e1b3860ffa5bcd06a3799

6. Check that database is created by typing:

   \l

And use arrow keys to scroll down.

7. Go to Connections:

![Alt text](image-2.png)

GO to Networking, scroll down and ADD A NETWORK with the following settings (or the ip address of the backend):

![Alt text](image-3.png)

8. Now you can deploy the rest using:

   terraform apply
