# Deployment Instructions

## Assumptions

For the purposes of this project, we make the following assumptions

* you are executing this on a linux machine
* you have a working go install
* you have a working minikube installation with the nginx-ingress controller installed
* you have a working docker installation

## Quickstart

To deploy this project, the simplest method is to execute the following:

```bash
eval $(minikube docker-env)
docker build -t go-hello .
kubectl create -f go-hello-deploy.yaml
```

you may then check to ensure the pods are correctly running by executing `kubectl describe -f go-hello-deploy.yaml`. You should see 4 available replicas listed in the Replicas field at the top.

To access the service, you must first obtain the ingress IP with the following command `kubectl get ingress go-hello-ingress`. Assuming the IP provided is: 192.168.122.138 you may then test the service by executing `curl -v http://192.168.122.138/` multiple times. Please note, it may take several minutes before minikube provides an ingress IP through the CLI tools.

To remove the running containers, this is most easily achieved by executing: `kubectl delete -f go-hello-deploy.yaml`

## go-hello

go-hello is automatically built as part of the `docker build`, but should you wish to do this manually, the simplest solution is to simply run `go build` in the go-hello directory.

Cross compilation may be required if you are not running this on a linux machine.

## Docker container and registry

The simplest way to build the docker container and to push it to a registry minikube can use, is to first inherit the minikube docker information, then to build the container using the provided Dockerfile.

To do this execute the following:

```bash
eval $(minikube docker-env)
docker build -t go-hello .
```

to test it works, you can run the container with the following: `docker run -p 8081:8081 go-hello` and test with `curl -v $(echo ${DOCKER_HOST} | awk -F[/:] '{print $4}'):8081`

## kubernetes

To deploy go-hello on your minikube cluster, you need to first create the necessary resources by executing:

`kubectl create -f go-hello-deploy.yaml`

if is then possible to inspect the running status of these resources several ways. Some of the more detailed views would come from the following:

```bash
kubectl describe deployments # Inspect deployment status
kubectl describe pods        # Inspect running pod status
kubectl describe ingress     # Inspect the ingress status
```

## Additional Information

Authored by John Mylchreest (jmylchreest@technauts.co.uk) on 21/02/2019 for the Equal Expert stage 1 interview.

This package contains responses believed to be appropriate for the minikube requirements, outlined as:

"Mini Kube

Write a simple hello world application in either of these languages: Python, Ruby, Go. Build the application within a Docker container and then load balance the application within a mini kube."

Although out of scope for the request, future improvements may include:

* application instrumentation, and metric collection tooling such as prometheus
* structured logging (I like using JSON based log messages for easy ingest into log collection platforms) output
* more robust testing
* CI/CD Pipeline tooling, such as AWS CodeBuild/CodeDeploy or GoCD/Helm, or Spinnaker etc.