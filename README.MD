# Mando

### Command
```
kubectl cluster-info
kubectl config current-context
kubectl get nodes
kubectl logs <pod-name>
kubectl get deployments
kubectl get pods
kubectl get services
kubectl get pvc
```

Apply configuration in order
```
# Fetch images from docker local when imagePullPolicy: IfNotPresent 
# and Apply all resources defined in your application stack file
eval $(minikube -p minikube docker-env) 
docker build -t mando-app:latest .
eval $(minikube -p minikube docker-env -u)
kubectl apply -f mando-stack-kube.yaml

# Apply the Argo CD Application definition (to the correct namespace)
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl apply -f argocd-app-mando.yaml -n argocd

# Forward local port 8082 to the Argo CD server's port 443 in namespace argocd
kubectl port-forward svc/argocd-server -n argocd 8082:443
argocd admin initial-password -n argocd
```

```
kubectl delete -f mando-stack-kube.yaml
kubectl delete -f argocd-app-mando.yaml
kubectl delete secret argocd-secret -n argocd

kubectl rollout restart deployment argocd-server -n argocd
kubectl rollout restart deployment argocd-repo-server -n argocd
kubectl rollout restart deployment argocd-application-controller -n argocd

# 1. Create the namespace (run this if the namespace doesn't exist)
kubectl create namespace argocd

# 2. Apply the standard installation manifest
kubectl delete -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl delete namespace argocd
```

### Openstack
```
https://docs.openstack.org/devstack/latest/
```