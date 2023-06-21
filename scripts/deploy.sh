#!/usr/bin/env bash

source ./scripts/env.sh

set -o allexport; source .env; set +o allexport

function error() {
	echo ${1}
	exit 1
}

# pulling latest code
echo "Pulling latest code..."
git checkout $DEPLOY_BRANCH
git pull origin $DEPLOY_BRANCH || error "Failed to pull latest code..."

# building fresh binaries
echo "Testing new code..."
if [ ! -d $OUTPUT_DIRECTORY ]; then
	mkdir -p $OUTPUT_DIRECTORY || error "Failed to create build directory..."
fi

sudo /usr/local/go/bin/go build -o $BINARY ./cmd/backend/*.go || error "Failed to compile new code..."

# stop service
echo "Stopping server..."
sudo service $PROJECT_NAME stop # || error "Failed to stop server..."

# removing old build
echo "Removing old build..."
sudo rm -rf $DEPLOY_DIRECTORY || error "Failed to remove last build..."

# creating new build
echo "Creating new build..."
sudo mkdir -p $DEPLOY_DIRECTORY || error "Failed to create target directory..."

echo "Configuring environment..."
sudo cp .env $DEPLOY_DIRECTORY/ || error "Failed to update environment variables..."

echo "Copying binaries..."
sudo cp $BINARY $DEPLOY_TARGET || error "Failed to copy new binaries..."

# copying the service file
echo "Configuring service..."
SERVICE_FILE=$PROJECT_NAME.service
sudo cp $SERVICE_FILE /etc/systemd/system/ || error "Failed to copy service file..."
sudo systemctl daemon-reload || error "Failed to restart daemon..."
sudo systemctl enable $SERVICE_FILE || error "Failed to enable service..."

# starting the service
echo "Starting server..."
sudo service $PROJECT_NAME start || error "Failed to start server..."

echo "Deployment Successful..."
