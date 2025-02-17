#!/bin/bash

set -e # Stop the script if any command fails

# Function to install packages
install_package() {
    local package_name=$1
    local install_command=$2
    local version_command=$3
    local version_label=$4

    if command -v "$package_name" &>/dev/null; then
        local version=$($version_command)
        echo "✓ $package_name detected! Version: $version"
    else
        echo ""
        echo "📦 Installing $package_name..."
        eval "$install_command"
    fi
}

# OS detection
OS=$(uname -s)

# Install packages for macOS with Homebrew
if [[ "$OS" == "Darwin" ]]; then
    echo ""
    echo "🖥️ Darwin detected! Checking dependency tree.."

    install_package "minikube" "brew install minikube" "minikube version --output=yaml" "Minikube"
    install_package "kubectl" "brew install kubectl" "kubectl version --client --output=json | awk -F '\"gitVersion\":\"' '{print $2}' | awk -F '\"' '{print $1}'" "kubectl"
    install_package "helm" "brew install helm" "helm version --short" "Helm"
    install_package "docker" "brew install --cask docker" "docker --version | awk '{print \$3}' | sed 's/,//'" "Docker"

# Install packages for Windows or Unix with git-bash
elif [[ "$OS" == "MINGW"* || "$OS" == "CYGWIN"* ]]; then
    echo ""
    echo "🖥️ WIN detected! Checking dependency tree.."

    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
    export PATH="$INSTALL_DIR:$PATH"
    echo 'export PATH="$HOME/.local/bin:$PATH"' >>~/.bashrc

    install_package "minikube" "curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-windows-amd64.exe && mv minikube-windows-amd64.exe \"$INSTALL_DIR/minikube\" && chmod +x \"$INSTALL_DIR/minikube\"" "minikube version --short" "Minikube"
    install_package "kubectl" "curl -LO \"https://dl.k8s.io/release/\$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/windows/amd64/kubectl.exe\" && mv kubectl.exe \"$INSTALL_DIR/kubectl\" && chmod +x \"$INSTALL_DIR/kubectl\"" "kubectl version --client --output=json | sed -n 's/.*\"gitVersion\":\"\\([^\\\"]*\\)\".*/\\1/p'" "kubectl"
    install_package "helm" "curl -LO https://get.helm.sh/helm-v3.8.0-windows-amd64.zip && unzip helm-v3.8.0-windows-amd64.zip && mv windows-amd64/helm \"$INSTALL_DIR/helm\" && chmod +x \"$INSTALL_DIR/helm\" && rm -rf windows-amd64 helm-v3.8.0-windows-amd64.zip" "helm version --short" "Helm"
    install_package "docker" "curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh && rm get-docker.sh" "docker --version | awk '{print \$3}' | sed 's/,//'" "Docker"

else
    echo "Unsupported operating system!"
    exit 1
fi

echo ""
echo "👍 Dependency injection completed!"
