#!/bin/bash

set -e # Exit script on error

# Minikube profile name
PROFILE_NAME="igi-app-dev"

# Check if the script is running with admin rights
echo ""
if [[ "$OSTYPE" == "msys" ]]; then
    # Windows
    if ! net session &>/dev/null; then
        echo "❗ Please run this script as Administrator"
        exit 1
    fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    if [[ $EUID -ne 0 ]]; then
        echo "❗ Please run this script as root"
        exit 1
    fi
else
    echo "Unsupported OS"
    exit 1
fi
echo "🚀 Initializing the dev cluster..."

# Start Minikube
if ! minikube status -p $PROFILE_NAME &>/dev/null; then
    echo "🌟 Starting Minikube for Development cluster..."
    minikube start --profile=$PROFILE_NAME --cpus=4 --memory=8192 --driver=hyperv --container-runtime=docker --dns-proxy
fi

# Checks
echo "🔧 Running cluster diagnostics..."

echo ""
minikube status -p $PROFILE_NAME
echo "✅  Minikube is running!"

# Enable Docker within Minikube
eval $(minikube -p $PROFILE_NAME docker-env)
echo "✅  Docker deamon is enabled within Minikube!"

# Start docker inside minikube
minikube ssh -p $PROFILE_NAME -- sudo systemctl start docker
# Ensure Docker is running inside Minikube
minikube ssh -p $PROFILE_NAME -- docker info &>/dev/null && echo "✅  Docker engine is running in Minikube." || {
    echo "❌ Docker engine is not running in Minikube. Restarting Docker..."
    minikube ssh -p $PROFILE_NAME -- sudo systemctl restart docker
}
echo "✅  All checks passed!"

echo ""
echo "  👍  '$PROFILE_NAME' is up and ready for deployments."
