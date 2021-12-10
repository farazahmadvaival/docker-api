# Docker Access API
API to get docker information from server. This api is 
written using Go language and used Gin 
Framework for API building.
## How to run:
```
# chmod +x docker-api
# ./docker-api "0.0.0.0:8088"
```
If you not provide the port after service name it 
will listen at port 8088.
## Access the api:
```
# curl http://localhost:8088/containers
# curl http://localhost:8088/restart-container/container-id
```