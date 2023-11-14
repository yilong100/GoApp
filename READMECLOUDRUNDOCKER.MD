To upload docker image onto artifact registry on gcp

1. Create a repository on artifact registry

2. docker tag 1dda5758810846f4450127510001cfc7dfa88594941e6f8f5b9be0bd41f9c546 australia-southeast2-docker.pkg.dev/rapid-chassis-402302/quickstart-docker-repo/docker-backend-image -> (docker tag [image_id] [zone_of_repository]/[project_id]/[repo_name]/[image_name])

3. docker push australia-southeast2-docker.pkg.dev/rapid-chassis-402302/quickstart-docker-repo/docker-backend-image    (on google cloud cli)

(You may need to make sure docker is configured to use gcloud as a credockdential helper by typing in gcloud auth configure-docker)
(If py command works but not python, type python=py $\*)

Then u can deploy your image on Cloud Run on GCP
