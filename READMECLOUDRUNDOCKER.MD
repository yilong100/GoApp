To upload docker image onto container registry on gcp

1. docker tag 8e5511afb5c86988c12dfb0d0f84dda30de0212fb2d43049e6573698e91ccdda gcr.io/rapid-chassis-402302/docker-frontend-image -> (docker tag [image_id] gcr.io/[project_id]/[image_name])

2. docker push gcr.io/rapid-chassis-402302/docker-frontend-image (on google cloud cli)

(You may need to make sure docker is configured to use gcloud as a credential helper by typing in gcloud auth configure-docker)
(If py command works but not python, type python=py $\*)

Then u can deploy your image on Cloud Run on GCP
