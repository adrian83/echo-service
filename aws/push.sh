DOCKER_IMAGE=$1
VERSION=$2

if [[ -z "$DOCKER_IMAGE" ]] || [[ -z "$VERSION" ]]
then
    echo ""
    echo "Error! Missing parameter(s)"
    echo "Usage ./push.sh <docker image id> <version>"
    echo ""
    exit 1
fi

APP='echo'
ACCOUNT_ID=`aws sts get-caller-identity | jq '.Account' | cut -d "\"" -f 2`
REGION=`aws configure get region`

echo Account id: $ACCOUNT_ID
echo Region: $REGION
echo Docker image id: $DOCKER_IMAGE
echo Version: $VERSION

# authenticate Docker to an Amazon ECR registry 
aws ecr get-login-password --region $REGION | docker login --username AWS --password-stdin $ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com

# tag Docker image
docker tag $DOCKER_IMAGE $ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com/$APP

# push image into ECR registry
docker push $ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com/$APP

# change tag to version number
MANIFEST=$(aws ecr batch-get-image --repository-name $APP --image-ids imageTag=latest --query 'images[].imageManifest' --output text)
aws ecr put-image --repository-name $APP --image-tag $VERSION --image-manifest "$MANIFEST"