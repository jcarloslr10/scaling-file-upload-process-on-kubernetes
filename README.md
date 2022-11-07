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

- Go 1.19 (https://go.dev/doc/install)
- Docker 20.10 (https://docs.docker.com/get-docker/)
- Minikube 1.27.0 (https://minikube.sigs.k8s.io/docs/start/)

## Install

Run the following command located in the `file-api` folder:

```sh
go install -v .
```

## Build

Run all the following commands located in the `file-api` folder.

### Docker environment

To use local Docker images in Minikube run:

```sh
minikube docker-env
minikube docker-env | Invoke-Expression
```

To build local Docker image run:

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
kubectl apply -f file-api.yml
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