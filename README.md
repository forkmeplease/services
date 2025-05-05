# M3O

A Go microservices platform

## Overview

M3O provides the fundamental building blocks for any products, apps or service. It can be used in isolation 
or combined with other tools to create powerful distributed systems. M3O services are intended to be consumed 
by each other using RPC and externally through the Micro API.

## Usage

M3O depends on [Micro](https://github.com/micro/micro)

### Run the Micro API

Install and run the API

```
go get github.com/micro/micro/v5/cmd/micro-api@latest
```

Run the api
```
micro-api
```

### Run a Service

Run a service from source

```
micro run github.com/micro/m3o/helloworld
```

### Call a Service

Call it through the API

```
curl "http://localhost:8080/helloworld/Call?name=Alice"
```
