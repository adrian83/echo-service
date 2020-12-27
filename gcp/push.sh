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
PROJECT=`gcloud config get-value project`

echo Project: $PROJECT
echo Docker image id: $DOCKER_IMAGE
echo Version: $VERSION


# tag Docker image
docker tag $DOCKER_IMAGE gcr.io/$PROJECT/$APP:$VERSION

# push image into Container Registry
docker push gcr.io/$PROJECT/$APP:$VERSION
