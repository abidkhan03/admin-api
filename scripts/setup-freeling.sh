#!/usr/bin/env bash

# copying the service file
echo "Configuring service..."
SERVICE_FILE=freeling.service
sudo cp $SERVICE_FILE /etc/systemd/system/ || error "Failed to copy service file..."
sudo systemctl daemon-reload || error "Failed to restart daemon..."
sudo systemctl enable $SERVICE_FILE || error "Failed to enable service..."

# starting the service
echo "Starting server..."
sudo service $PROJECT_NAME start || error "Failed to start server..."

echo "Deployment Successful..."
