#!/bin/bash

# Check correct usage of the script
if [ "$#" -eq 0 ]; then
   echo "Usage:  ./bootstrap_cluster.sh <cluster_name>"
   echo "   eg:  ./bootstrap_cluster.sh kubnerable"
   exit
fi

# Check if Kind is available locally
kind version > /dev/null 2>&1
if [ $? -eq 0 ];
then
    echo "✓ Kind is available."
else
    echo "✕ Kind is not available."
    exit
fi

# Check if Docker is available locally
docker version --format '{{.Server.Version}}' > /dev/null 2>&1
if [ $? -eq 0 ];
then
    echo "✓ Docker is available."
else
    echo "✕ Docker is not available."
    exit
fi

# Setup the cluster with the given CLUSTER_NAME
CLUSTER_NAME=$1
kind create cluster --name "${CLUSTER_NAME}" --config kind.yaml

# Check if Kubectl is available locally
kubectl version --short > /dev/null 2>&1
if [ $? -eq 0 ];
then
    echo "✓ kubectl is available and has access to the cluster."
else
    echo "✕ kubectl is not available."
    exit
fi

# Deploy the scenarios
kubectl apply -f ../scenarios/root_admin_webshell/deployment.yaml
kubectl apply -f ../scenarios/ipinfo_webapp/deployment.yaml
