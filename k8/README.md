## Architecture
The smarthome runs with a kubernetes cluster managed by minikube.
Each service is contained in a docker image hosted on DockerHub and pulled by kubernetes into individual deployments.

Secrets are managed in a Secret config manifest file

### kubectl cheat sheet
https://kubernetes.io/docs/reference/kubectl/cheatsheet/

Check all pods running in the cluster:
- `kubectl get pods`

Create a deployment
- `kubectl apply -f <deployment.yml`

Describe the deployment
- `kubectl describe <deployment>`

Describe a service:
- `kubectl describe service <service-name>`

Find external IP of a load balanced service:
- `kubectl get services <my-service>`

List all deployments:
- `kubectl get deployments`

List all services:
- `kubectl get services`

Delete a specific deployment
- `kubectl delete deployment <deployment-name>`

Force delete a pod after deleting its deployment:
- `kubectl delete pods <pod> --grace-period=0 --force`

### Port forward to service
To access the backend service from the react app running locally, you can portforward to one of its pods:
- `kubectl port-forward openweather-service-7f7c54dbf8-29jbm 8080:8080`

### Exposing a service:
There are 4 configuration we can use to expose a service:
- ClusterIP: (default value if not specified). The service is available only inside the cluster
- LoadBalancer: expose application on public network
- NodeIP: exposes the service on each nodeIP at a static port (NodePort)
- Ingress: exposes HTTP and HTTPS routes from outside the cluster to services within the cluster



This command "tunnel" is useful only for minikube because it does not support service:LoadBalancer
- `minikube service redis-master --url`