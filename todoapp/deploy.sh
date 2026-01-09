#!/usr/bin/env bash
docker build -t todoapp:latest .
k3d image import todoapp:latest
kubectl apply -f manifests/deployment.yaml
