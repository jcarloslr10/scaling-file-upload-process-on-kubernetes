<h1 align="center">Welcome to Scaling file upload process on Kubernetes 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: ISC" src="https://img.shields.io/badge/License-ISC-yellow.svg" />
  </a>
  <a href="https://twitter.com/jcarloslr10" target="_blank">
    <img alt="Twitter: jcarloslr10" src="https://img.shields.io/twitter/follow/jcarloslr10.svg?style=social" />
  </a>
</p>

## Prerequisites

The application works with the versions listed below. It can work with other even lower versions although it has not been tested.

- [Go 1.19](https://go.dev/doc/install/)
- [Docker 20.10](https://docs.docker.com/get-docker/)
- [Minikube 1.27.0](https://minikube.sigs.k8s.io/docs/start/) with the following enabled addons:
    * dashboard: `minikube dashboard`
    * default-storageclass
    * ingress: `minikube addons enable ingress`
    * metrics-server: `minikube addons enable metrics-server`
    * storage-provisioner
- [K6](https://k6.io/docs/)

## Project components

- REST API developed in Go to upload PDF files is located in `file-api` folder.
    * Index page ([`index.html`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/blob/main/file-api/index.html)) to render a form to upload a PDF file.
    * API endpoints to handle requests ([`main.go`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/blob/main/file-api/main.go)).
    * Docker support using [`Dockerfile`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/blob/main/file-api/Dockerfile) and [`docker-compose.yml`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/blob/main/file-api/docker-compose.yml) files.
- Manifests (`.yml`) to deploy the API in Kubernetes cluster.
    * API manifests are located in the [`file-api-part-i.yml`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/blob/main/k8s/file-api-part-i.yml) file.
    * Auto-scaling manifest is located in the [`file-api-part-ii.yml`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/blob/main/k8s/file-api-part-ii.yml) file.
- Load tests are located in [`load-tests`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/tree/main/load-tests) folder.

## Install

Run the following command located in the `file-api` folder:

```sh
go install -v .
```

## Build

Run all the following commands located in the `file-api` folder.

### Docker environment

To use local Docker images in Minikube run the following statements depending on the OS:

- Windows
  ```sh
  minikube docker-env | Invoke-Expression
  ```

- Linux
  ```sh
  eval $(minikube docker-env)
  ```

Then, to build local Docker image run:

```sh
docker build --tag local/file-api:latest .
```

## Usage

Run all the following commands located in the `file-api` folder.

### Local environment

```sh
go run .
```

- Index endpoint (GET) available on: `http://127.0.0.1:8080`
- Upload endpoint (POST) available on: `http://127.0.0.1:8080/upload`

### Docker environment

```sh
docker compose up
```

App will be accesible from the following endpoints:

- Index endpoint (GET) on: `http://127.0.0.1:8000`
- Upload endpoint (POST) on: `http://127.0.0.1:8000/upload`

### Kubernetes environment

Run the following command located in the `k8s` folder:

```sh
kubectl apply -f file-api-part-i.yml
```

Run the following command to get the ingress IP address:

```sh
kubectl get ingress -n file-api
```

It should output something like:

```
NAME               CLASS    HOSTS         ADDRESS        PORTS   AGE
file-api-ingress   <none>   fileapi.com   172.31.54.63   80      21m
```

Open the `hosts` file (Windows/Linux) and add the following entry at the end of the file depending on the OS:

- Windows (`C:\Windows\System32\drivers\etc\hosts`)
  - `172.31.54.63 fileapi.com`
- Linux (`/etc/hosts`)
  - `172.31.54.63 fileapi.com`

Open the browser on the following URL: `http://fileapi.com`

If you want to take a look at the container filesystem, run the following command to get access to the container shell:

```sh
kubectl exec -it -n=file-api file-api-deployment-<hash> -- /bin/ash
ls -la
```

App will be accesible from the following endpoints:

- Index endpoint (GET) on: `http://fileapi.com`
- Upload endpoint (POST) on: `http://fileapi.com/upload`

## Load tests

The K6 tool helps us simulate a heavy workload scenario to allow HPA (Horizontal Pod Autoscaler) to come into the picture.

Run the following command to start the load test:

```
k6 run .\index.js
```

Then, run the following command to watch the current status of the deployment using HPA:

```
kubectl get hpa -n file-api --watch
```

If you want to simulate a heavier workload, you can go to [`index.js`](https://github.com/jcarloslr10/scaling-file-upload-process-on-kubernetes/blob/main/load-tests/index.js#L9) file located in `load-tests` folder and adjust the parameters of the `fileapi` scenario according to the K6 documentation (https://k6.io/docs/using-k6/scenarios/). Here is the specific code:

```
export const options = {
    discardResponseBodies: true,
    scenarios: {
        fileapi: {
            executor: 'constant-vus',
            vus: 20,
            duration: '1m00s',
        },
    },
};
```

As output of the load test, two reports will be generated:

- Console report.
- HTML report (result.html).

These reports show a set of metrics and counters related to requests sent to the API to upload documents.

## Author

👤 **Juan Carlos López**

* Website: https://jcarloslr10.github.io
* Twitter: [@jcarloslr10](https://twitter.com/jcarloslr10)
* Github: [@jcarloslr10](https://github.com/jcarloslr10)
* LinkedIn: [@juan-carlos-lópez-6154017b](https://linkedin.com/in/juan-carlos-lópez-6154017b)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_