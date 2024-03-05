# Proxy Server

This section describes techniques to create networking proxies. There are two types of network proxies: forward or reverse (see  [What is a proxy server?](https://www.cloudflare.com/learning/cdn/glossary/reverse-proxy/)).

A forward proxy is a server that sits in front of a group of client machines (see Figure 1).

![Forward proxy](../assets/img/forward_proxy_flow.png)</br>
**Figure 1: Forward Proxy**

A reverse proxy is a server that sits in front of backend machines, for example, webservers, fileservers, etc (see Figure 2).

![Reverse proxy](../assets/img/reverse_proxy_flow.png)</br>
**Figure 2: Reverse Proxy**

## HTTPUtil based Proxy

In this example, we demonstrate an implementation using `httputil` package to build a proxy (Forward or Reverse). The source code for this is [here](../cmd/proxy/httputil/main.go).

An example of the servers deployment is based on [docker-compose](../deployment/docker-compose.yml).  To run this example:

* STEP 1: Ensure docker deamon is running.
* STEP 2: Open a terminal and run the command `./scripts/ops.sh image build` - this builds the proxy server named `httphttputilproxy` and a webserver (running on `http://localhost:8080`).
* STEP 3: Start the docker deployment network by running the command `./scripts/ops.sh network start`.
* STEP 4: Open browser and call the url `http://localhost:3030` which will redirect to `http://localhost:8080`.

## Custom Proxy using Http Client

In this example, we demonstrate an implementation using `http` package, in particular client, to build a proxy (Forward or Reverse). The source code for this is [here](../cmd/custom/main.go).

* STEP 1: Ensure docker deamon is running.
* STEP 2: Open a terminal and run the command `./scripts/ops.sh image build` - this builds the proxy server named `customproxy` and a webserver (running on `http://localhost:8080`).
* STEP 3: Start the docker deployment network by running the command `./scripts/ops.sh network start`.
* STEP 4: Open browser and call the url `http://localhost:3031` which will redirect to `http://localhost:8080`.


