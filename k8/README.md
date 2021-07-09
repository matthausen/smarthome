## Architecture
The smarthome is made of services running as indipendent pods.
Minikube is is the Kubernetes cluster that takes care of everything.

Each application has its own deployment and service and runs an image hosted in the Docker Hub container registry.

### kubectl cheat sheet
https://kubernetes.io/docs/reference/kubectl/cheatsheet/

Check all pods running in the cluster:
- `kubectl get pods`

Create a deployment
- `kubectl apply -f <deployment.yml`

Describe the deployment
- `kubectl describe <pod name>`

List all deployments:
- `kubectl get deployments`

Delete a specific deployment
- `kubectl delete deployment <deployment-name>`

Force delete a pod after deleting its deployment:
- `kubectl delete pods <pod> --grace-period=0 --force`


### Redis
The redis database exposes port 6379.
However to connect to it from an external app, you can run:
- `minikube service redis-master --url`
to find out which port you can use for connection.