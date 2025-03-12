# k8s-clientgo-demo

A simple example that shows how to use client-go to interact with a Kubernetes cluster by creating a Deployment and a Service.

## Usage

1. Run the Go program to create the pod:
   ```sh
   go run main.go

2. Verify that the pod is created:
   ```sh
   kubeectl get deployment
   kubectl get pods
   kubectl get service
