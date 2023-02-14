# Solution proposed to Stakefish challenge
The original exercise description is [HERE](EXERCISE.md)

## Overview
This document aims to give an understanding about how the exercise has been solved. It also tries to enfatice what's been and what has not been done and what would need to be improved.  
As an overview, the following technologies have been used to build this project from the ground up:
  - `Go` as programming language for building the Rest API.
  - `Postgresql` as the storage backend for the Rest API.
  - `Docker` for application ontainerization.
  - `Docker compose` for setting up a quick dev environment.
  - `Minikube` for local `k8s` development.
  - `Helm` for `k8s` application packaging.
  - `GitHub Actions` for the CI pipeline.

## Rest API
The rest API has been developed using Go as programming language, which is very conveniente for these kind of backends. Apart from the standard library, the following third party libraries have been used:
  - `gorilla/mux`: simple HTTP request router and dispatcher.
  - `go-pg`: Go ORM for Postgres database backends.

The following endpoints have been implemented according specification:
  - `/`
  - `/health`
  - `/v1/tools/lookup`
  - `/v1/tools/validate`
  - `/v1/history`
### Notes regarding endpoints
Regarding `/`, it returns the software version deployed. It gets this info from env var `STAKEFISH_API_VERSION`, which is injected at Docker image creation. Also to find out if it's running in a k8s cluster, it check env var `KUBERNETES_SERVICE_HOST`.  
  
The `/health` just return if the connection between the application and the database backend is healthy.

The following have NOT been implemented due to lack of time:
  - Swagger retrieval: I would have implemented another endpoint for the user to get this definition so that it could autogenerate client side code. I.e: `/swagger.yaml` and `/swagger.json`.  
  - `/metrics` endpoint: I would have added another component that could return metrics when queried. Based on [Instumenting a Go application](https://prometheus.io/docs/guides/go-application/).
  - Improvement of swagger definition: for instance, adding `/health` endpoint to it.  

### Software architecture
The software has three major components:
  - `Controller`: This is the part that handles HTTP incoming requests. It defines the endpoints and how to handle the incoming requ

### Database backend

### Improvements that I'd done if I had more time
  - Add dependabot to GitHub

## Docker related work

## Helm related work

## GitHub Actions related work