#!/bin/bash

set -e # Hata alındığında scripti durdur

echo ""
echo "🚀 Running system diagnostics..."
sleep 2

# Determine package manager
if command -v apt &>/dev/null; then
    PACKAGE_MANAGER="apt"
    UPDATE_CMD="sudo apt update -y"
    INSTALL_CMD="sudo apt install -y"
elif command -v yum &>/dev/null; then
    PACKAGE_MANAGER="yum"
    UPDATE_CMD="sudo yum update -y"
    INSTALL_CMD="sudo yum install -y"
elif command -v brew &>/dev/null; then
    PACKAGE_MANAGER="brew"
    UPDATE_CMD="brew update"
    INSTALL_CMD="brew install"
elif command -v choco &>/dev/null; then
    PACKAGE_MANAGER="choco"
    UPDATE_CMD="choco upgrade chocolatey -y"
    INSTALL_CMD="choco install -y"
elif command -v snap &>/dev/null; then
    PACKAGE_MANAGER="snap"
    UPDATE_CMD=""
    INSTALL_CMD="sudo snap install"
else
    echo "❌ No supported package manager found. Install a package manager (brew, apt, choco, yum, snap) manually and rerun the script."
    exit 1
fi

echo "📦 Detected package manager: $PACKAGE_MANAGER"

# Step 2: Required Dependencies (K8s & Docker Infra)
declare -A dependencies
dependencies=(
    ["Docker"]="docker"
    ["Minikube"]="minikube"
    ["kubectl"]="kubectl"
    ["Helm"]="helm"
)

MISSING=()
for dep in "${!dependencies[@]}"; do
    if ! command -v ${dependencies[$dep]} &>/dev/null; then
        MISSING+=("$dep")
        echo "🔗 Detected missing dependency: $dep"
    fi
done

echo "✅ Diagnostic complete!"

# Step 4: Install missing dependencies

if [ ${#MISSING[@]} -eq 0 ]; then
    echo ""
    echo "🎉 All required dependencies are already installed!"
    echo "✨ System is ready to run the infrastructure."
    exit 0
fi



# Step 3: If missing dependencies found, ask for approval

if [ ${#MISSING[@]} -gt 0 ]; then
    echo ""
    echo "  ⚠️  The following dependencies are required to run the infrastructure:"
    echo "-----------------------------------"
    printf "| %-20s | %-10s |\n" "Dependency" "Status"
    echo "-----------------------------------"
    for dep in "${MISSING[@]}"; do
        printf "| %-20s | ❌ Missing |\n" "$dep"
    done
    echo "-----------------------------------"
    echo ""
    read -p "       ❓ Do you approve installing them? (y/n): " APPROVE
    if [[ "$APPROVE" != "y" ]]; then
        echo ""
        echo "❌ Terminating script..."
        exit 1
    fi
fi




# Step 4: Install missing dependencies

echo "🏃‍♂️ Installing missing dependencies..."
sleep 2

for dep in "${MISSING[@]}"; do
    case "$dep" in
    "Docker")
        echo "🔧 Installing Docker..."
        if [[ "$PACKAGE_MANAGER" == "apt" ]]; then
            curl -fsSL https://get.docker.com | sh
        elif [[ "$PACKAGE_MANAGER" == "brew" ]]; then
            brew install --cask docker
        else
            eval "$INSTALL_CMD docker"
        fi
        ;;
    "Minikube")
        echo "🔧 Installing Minikube..."
        if [[ "$PACKAGE_MANAGER" == "brew" ]]; then
            brew install minikube
        elif [[ "$PACKAGE_MANAGER" == "choco" ]]; then
            choco install minikube -y
        else
            curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
            sudo install minikube-linux-amd64 /usr/local/bin/minikube
            rm minikube-linux-amd64
        fi
        ;;
    "kubectl")
        echo "🔧 Installing kubectl..."
        if [[ "$PACKAGE_MANAGER" == "snap" ]]; then
            sudo snap install kubectl --classic
        elif [[ "$PACKAGE_MANAGER" == "brew" ]]; then
            brew install kubectl
        else
            eval "$INSTALL_CMD kubectl"
        fi
        ;;
    "Helm")
        echo "🔧 Installing Helm..."
        if [[ "$PACKAGE_MANAGER" == "brew" ]]; then
            brew install helm
        elif [[ "$PACKAGE_MANAGER" == "choco" ]]; then
            choco install kubernetes-helm -y
        else
            curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
        fi
        ;;
    esac
    echo "✅ $dep installed successfully."
done

echo "🎉 All dependencies installed successfully!"
