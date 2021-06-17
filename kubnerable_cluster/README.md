# Kubnerable cluster
Kubnerable cluster is a collection of different scenarios designed to be insecure and vulnerable. The objective is to have an easy-to-setup insecure local environment to perform security researching, powered by Kind (*Kubernetes-in-Docker*).

In each scenario, different vulnerabilities will be set up in place.

## Requirements
* kind
* Docker
* kubectl

## How to set up
* Run `kind/bootstrap_cluster.sh <cluster_name>` to create the cluster
* To delete the cluster, run `kind delete cluster --name <cluster_name>`

## Scenarios
- [x] IP Info Webapp
- [x] Cluster Administrator Shell
