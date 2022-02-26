#!/bin/bash

echo "$ minikube start --cpus 4 --memory 4096 --vm"
minikube start --cpus 4 --memory 4096 --vm
echo "$ minikube kubectl get nodes"
minikube kubectl get nodes
echo "$ minikube kubectl get namespaces"
minikube kubectl get namespaces
echo "$ minikube kubectl get pods"
minikube kubectl get pods
echo "$ minikube kubectl create deploy hello-world -- --image='quay.io/connordoyle/cpuset-visualizer'"
minikube kubectl create deploy hello-world -- --image="quay.io/connordoyle/cpuset-visualizer"
minikube kubectl wait -- --for=condition=ready -l app=hello-world
echo "$ minikube kubectl get pods"
minikube kubectl get pods
echo "$ minikube kubectl expose deploy hello-world -- --type=NodePort --port=80"
minikube kubectl expose deploy hello-world -- --type=NodePort --port=80
echo "$ minikube kubectl get service"
minikube kubectl get service
echo "$ minikube service hello-world"
minikube service hello-world