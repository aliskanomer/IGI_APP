#!/bin/bash

set -e  # Hata alındığında scripti durdur

echo ""
echo ""
echo "🚀 Starting IGI_APP setup..."

# Run system diagnostic
./sys-diag.sh

# Run configuration script
./sys-config.sh

echo "✅ Setup completed successfully!"
