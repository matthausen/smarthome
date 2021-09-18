## Minikube

Minikube can be used for deployment. However it does not support nginx ingress to clusterIP services, as it runs in a VM.
As a workaround, services can be exposed as NodePort or LoadBalancer.
Alternatively, you can run with this configuration in a production environment with an actual cluster.

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

Port forward to a service
- `kubectl port-forward openweather-service-7f7c54dbf8-29jbm 8080:8080`

### Types of services
There are 4 configuration we can use to expose a service:
- ClusterIP: (default value if not specified). The service is available only inside the cluster
- NodePort:  Exposes the Service on the same port of each selected Node in the cluster using NAT
- LoadBalancer - Creates an external load balancer in the current cloud (if supported) and assigns a fixed, external IP to the Service. Superset of NodePort.
- Ingress: exposes HTTP and HTTPS routes from outside the cluster to services within the cluster

### Minikube cheatsheet:
Get the external IP of the cluster
- `minikube ip` Get the IP where the cluster is running

Expose service externally
- `minikube service --url <service-name>`

Loadbalancer services must be exopsed via the `tunnel` command.