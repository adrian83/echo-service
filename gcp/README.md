# Google Cloud Platform

#### Prerequisites:
1. GCP Cli (gcloud) (run `gcloud init`)
2. Docker


#### Create project: 
`gcloud projects create <project-name>`
##### Example:
`gcloud projects create docker-registry-demo-27122020`  


#### Set newly created project as active: 
`gcloud config set project <project-name>`
##### Example:
`gcloud config set project docker-registry-demo-27122020`  


#### Enable billing:
[Read this](https://support.google.com/googleapi/answer/6158867?hl=en)


#### Enable service:
`gcloud services enable containerregistry.googleapis.com`


#### Clean up - remove project:
`gcloud projects delete <project-name>`
##### Example:
`gcloud projects delete docker-registry-demo-27122020`  


#### Build new Docker image:
`docker build -t echo:<version> .`  


##### Example:
`docker build -t echo:1.0.8 .`  


#### List Docker images:
`docker images`  


#### Push new Docker image:
`./push.sh <docker_image-id> <version>`  


##### Example:
`./push.sh 1e442ffbf3ef 1.0.8`
