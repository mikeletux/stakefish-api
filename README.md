# Solution proposed to Stakefish challenge
The original exercise description is [HERE](EXERCISE.md)

## 1. Overview
This document aims to give an understanding about how the exercise has been solved. It also tries to enfatice what's been and what has not been done and what would need to be improved.  
As an overview, the following technologies have been used to build this project from the ground up:
  - `Go` as programming language for building the Rest API.
  - `Postgresql` as the storage backend for the Rest API.
  - `Docker` for application ontainerization.
  - `Docker compose` for setting up a quick dev environment.
  - `Minikube` for local `k8s` development.
  - `Helm` for `k8s` application packaging.
  - `GitHub Actions` for the CI pipeline.

## 2. Rest API
The rest API has been developed using Go as programming language, which is very conveniente for these kind of backends. Apart from the standard library, the following third party libraries have been used:
  - `gorilla/mux`: simple HTTP request router and dispatcher.
  - `go-pg`: Go ORM for Postgres database backends.

The following endpoints have been implemented according specification:
  - `/`
  - `/health`
  - `/v1/tools/lookup`
  - `/v1/tools/validate`
  - `/v1/history`
### 2.1 Notes regarding endpoints
Regarding `/` endpoint, it returns the software version deployed, among other things. It gets this info from env var `STAKEFISH_API_VERSION`, which is injected at `Docker` image creation. It also returns if the app is running in `k8s`. To find that out it checks env var `KUBERNETES_SERVICE_HOST`.  
  
The `/health` endpoint just return if the connection between the application and the database backend is healthy.

The following have NOT been implemented due to lack of time:
  - Swagger retrieval: I would have implemented another endpoint for the user to get this definition so that users could autogenerate client side code. I.e: `/swagger.yaml` and `/swagger.json`.  
  - `/metrics` endpoint: I would have added another component that could return metrics when queried. Based on [Instumenting a Go application](https://prometheus.io/docs/guides/go-application/).
  - Improvement of swagger definition: for instance, adding `/health` endpoint to it.  

### 2.2 Software architecture
The Rest API software has three major packages:
  - `Controller`: This package handles HTTP incoming requests. It defines the endpoints and how to handle the incoming requests.  
  This package also contains some unit tests for `/` and `/v1/tools/validate` endpoints. Testing would have needed to be more throrough, but at least it works to ilustrate the `GitHub Actions` testing step.  
  The main component, `Manager` needs three structs that implement three different interfaces. This is done to avoid having tighly coupled components and allow unit testing. The structs that implement those interfaces are passed in the `Manager` instantiator following the dependency injection paradigm. These interfaces are:
    - Database connector: Allows the `Manager` to interact with a database backend. Useful for saving/quering information.
    - Network Infrastructure: Allows the `Manager` to interact with the network, either the actual Internet or a mocked component.
    - Logger: useful for loggin incomming request when things happen within the `Manager` component.   

  - `Infrastructure`: This package contains the interfaces as well as the implementation that allows the access to the database backend as well as the internet. It also includes the mock objects used for unit testing.

  - `Debug`: This package contains the interface for logging as well as a simple struct that implements it.

Apart from these three packages, there is the additional one `models` that keeps the structs with the needed tags to return the right json objects to the users upon request and also the tags for modeling the tables for the postgres ORM used.

Also the main function have the logic to gracefully close the HTTP server and connection to the database upon user request, either doing `CTRL+C` when running standalonse or shutting down the container gracefully when running on top of docker/k8s.


### 2.3 Database backend
The database backend chosen for this project is postgresql. It uses the default `postgres` database and creates two tables on it. Tables are:
  - `query`: each request from users is logged here. The IPv4 resolved are stored in the `address` table.
  - `address`: all IPv4 from a user request are stored here. Each entry have a foreign key referencing one query.

The relation between `query` and `address` is 1-n, having one `query` many `address`.
The definition file for this can be found [HERE](db/create_tables.sql).

## 3. Docker related work
The application can be contenerized using Docker. To do so please refer to the [Dockerfile](Dockerfile). To comment it quickly, I've used the Go official image to build the binary and then used a distroless imaged based on debian for the final image. This improves storage use as well as image security.  
It is important to note that at build time the image building process needs the argument `stakefish_api_version` so it creates the env var `STAKEFISH_API_VERSION` inside the container. This will be the one read to show the application version upon user request.
  
The project also comes with a [docker-compose.yml](docker-compose.yaml) file so a development environment can be quicked off very quickly. As reference, the image building argument for this environming will be set to `testing`.

## 4. Helm related work
To quickly deploy the application on top of a `k8s` cluster, a `Helm` chart has been developed. To install the application using `Helm`, please add the right `Helm` reposotory and install it from there.  
**⚠️Before using it⚠️**, some `k8s` secrets need to be put in place. Please refer to the example [secrets.yaml](helm/stakefish-chart/example/secrets.yaml). If these secrets are not set before installing the chart, the pods won't come up. This behaviour could be improved in the future by auto-generating secrets if they are not present.  
This chart uses also as dependency the postgres chart created by [Bitnami](https://bitnami.com/stack/postgresql/helm).  
To install the chart, follow the steps below:
Add my `Helm` repo to your `Helm` installation:
```
$ helm repo add stakefish https://mikeletux.github.io/helm-chart/
```
Update your `Helm` repos:
```
$ helm repo update
```
Proceed to install the chart:
```
$ helm search repo stakefish
$ helm install stakefish stakefish/stakefish-chart
```

## 5. GitHub Actions related work
This project implements a CI pipeline built upon `GitHub Actions`. The pipeline is triggered every single time some code is pused to any branch, no matter which one. The only different behavior occurs when in a pull request. In this event, neither the `Docker` container built is pushed to github container registry nor the `Helm` chart is published.  
The pipeline has three steps (second and third stages depend on their previous steps respectively):
  - `lint-test-build`: Lints, runs unit tests and builds the Rest API from `Go` project.
  - `build-push-docker`: Builds and pushes the `Docker` image to GH image repository.
  - `package-helm-chart`: Packages and publishes the `Helm` chart.
  
Steps 2 and 3 creates and publishes artifacts.

Just to add more information about the container image building process, it uses as tags the branch name and the commit sha (short version) in which the pipeline was triggered. This is used as `app_version` when packaging the `Helm` chart so the chart is packaged with the same docker image tag as the docker image published to the repository.  

## 6. Additional work
I created a `Helm` repository using `Github Pages` on `https://mikeletux.github.io/helm-chart/`. In order to do that I needed to use the `GitHub Action` [stefanprodan/helm-gh-page](https://github.com/stefanprodan/helm-gh-pages) and followed his [tutorial](https://helm.sh/docs/howto/chart_releaser_action/). In order for this repo to perform changes on that one, I created one granular token allowing to modify `https://github.com/mikeletux/helm-chart` repo and used it from this repo pipeline as a `Github Secret`.   
  
The project also comes with a `k8s` [YAML](stakefish-k8s.yaml) file that could be used to deploy the needed resources on top of a `k8s` cluster without `Helm`.

/Miguel Sama 2023