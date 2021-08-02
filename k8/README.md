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