# echo-service
Simple application used for creating sample Docker images. Go to `aws` directory, to check, how to push newly created image to ECR registry.

#### Prerequisites:
1. Docker

#### Build image:
`docker build -t echo:<version> .`

##### Example
`docker build -t echo:1.0.7 .`