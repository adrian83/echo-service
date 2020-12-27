# Amazon Web Services

#### Prerequisites:
1. AWS Cli
2. Docker
3. jq (optional - if you want to use `push.sh` script)  


#### Create repository in Amazon Elastic Container Registry: 
`aws cloudformation deploy --template-file stack.yml --stack-name ecr-registry-echo`  


#### Clean up:
`aws cloudformation delete-stack --stack-name ecr-registry-echo`  


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

