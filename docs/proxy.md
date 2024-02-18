# Proxy Server

This section describes techniques to write networking proxy. There are two kinds of network proxies:

* Forward proxy.
* Reverse proxy.

## HTTPUtil based Proxy

In this example, we demonstrate an implementation using `httputil` package to build a proxy (Forward or Reverse). The source code for this is [here](../cmd/httputil/main.go).

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

## References

* [What is a reverse proxy?](https://www.cloudflare.com/en-gb/learning/cdn/glossary/reverse-proxy/)

