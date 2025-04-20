# Mando

### Command
```
kubectl cluster-info
kubectl config current-context # Should output 'minikube'
kubectl get nodes # Should show one node named 'minikube' with status 'Ready'
kubectl logs <pod-name>
```

Docker Environment (Optional but useful for local images): To build Docker images directly inside Minikube's Docker daemon (useful for using local images without pushing to a registry), use:

```
# Connect terminal's Docker client to Minikube's Docker daemon
eval $(minikube -p minikube docker-env)

# Build your image here...
# docker build -t my-local-image:latest .

# Disconnect when done
eval $(minikube -p minikube docker-env -u)
```

Apply configuration
```
# Apply all resources defined in your application stack file
kubectl apply -f mando-stack-kube.yaml

# Apply the Argo CD Application definition (to the correct namespace)
kubectl apply -f argocd-app-mando.yaml -n argocd

# Forward local port 8082 to the Argo CD server's port 443 in namespace argocd
kubectl port-forward svc/argocd-server -n argocd 8082:443
```