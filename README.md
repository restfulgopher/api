# API

## WHAT IS

The API repository contains an implementation for a simple REST API written in GO, with a single endpoint for validating International Banking Numbers (IBAN), used for international transactions.

The API design, documentation, and stub server were implemented using Swagger Editor and Swagger UI.

Docker and Terraform were used for provisioning and deployment. As NGINX was used for reverse proxying.

All services are available on Digital Ocean for a limited period.

The [OpenIban API](https://openiban.com/) performs the validation of the IBAN provided by the user. This implementation does not guarantee any reliability on the data returned by the external service.

API SPECIFICATION

	http://api.alesr.me/docs

PRODUCTION API (GO)

	curl -k -X GET 'Accept: application/json' 'https://api.alesr.me/v1/iban/valid/DE44500105175407324931'

STUB API

	curl -X GET 'Accept: application/json' 'http://api.alesr.me/v1/iban/valid/DE44500105175407324931'

## HOWTO

### `make editor/run`

Starts a Docker container from Swagger Editor image accessible at `http://127.0.0.1:8080/`, where it is possible to edit the API specification and generate a development (stub) API.

From the Swagger Editor interface, you can import the current API specification located at `/swagger/api/stub/api/swagger.yaml`, add new changes, and generate a new stub API in the `Generate Server` menu. Note that this implementation uses NODEJS as stub server.

After downloading the source code for the stub API, place it at `/swagger/stub/` directory, replacing all existing files. Run `make editor/stop` for stopping and removing the Swagger Editor container.

### `make run`

Build and start services for reverse proxying, documentation, stub and production API.

#### Using the services (local):

PRODUCTION API

	curl -k -X GET 'Accept: application/json' 'https://127.0.0.1/v1/iban/valid/DE44500105175407324931'

STUB API

	curl -X GET 'Accept: application/json' 'http://127.0.0.1/v1/iban/valid/DE44500105175407324931'

API SPECIFICATION

	http://127.0.0.1/docs

To stop all services and clean the environment run `make stop`.

### `make terraform/apply`

Build and push images to DockerHub, and use Terraform for provisioning an Digital Ocean droplet and deploy services.

Note that to use Terraform you must create a Digital Ocean API token and set it on your .bashrc `export DIGITALOCEAN_TOKEN="Your API TOKEN"`

You will also need the API token to retrieve you SSH key ID and past it in the Terraform main file.

```
curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer [API token here]" "https://api.digitalocean.com/v2/account/keys"
```

With the ID key in hand, place it on `main.tf` and run `make terraform/apply`.

```
resource "digitalocean_droplet" "vpn" {
  ssh_keys           = [012345]
  ...
```

To deploy this application under your own domain you should edit or remove the following code block at `/terraform/main.tf`.

The `make terraform/apply` build and push Docker images to Docker Hub before starting the provisioning. You may have to push the images to your own Docker Hub reposistory.

To destroy the droplets, run `make terraform/destroy`.

## AVAILABLE COMMANDS:

Run `make help` from the project root to list all available commands:

```
------------------------------------------------------------------------
RESTFULGOPHER API
------------------------------------------------------------------------
api/run                        start api, stub_api and nginx as reverse proxy
api/stop                       stop and remove services containers
api/test                       run api unit tests
editor/run                     start swagger editor container
editor/stop                    stop and remove swagger editor container
publish                        publish images on docker hub
terraform/apply                create remote vm with terraform and deploy services
terraform/destroy              destroy remote vm with terraform
```
