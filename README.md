# Echo-Service
Project presenting examples of pushing Docker images into Docker registers created in GCP and AWS. 
To check, how to push newly created image to:
- Elastic Container Registry (on Amazon Web Services) go to `aws` directory
- Container Registry (on Google Cloud Platform) go to `gcp` directory

#### Prerequisites:
1. Docker

#### Build image:
`docker build -t echo:<version> .`

##### Example
`docker build -t echo:1.0.7 .`
