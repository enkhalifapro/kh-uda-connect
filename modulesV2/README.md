# Overview
UdaConnect V2 consists of 3 microservice:
- PersonsService: Responsible for persons management APIs
- LocationsService: Responsible for locations management APIs
- ConnectionService: Responsible for connecting people based on their locations

# Run Instructions
## Apply base infrastructure 
Base Infra consists of the following service
- postgres
- kafka
- zookeeper (required by kafka)
- configmap
- secrets

## Apply PersonsService
- Navigate to modulesV2/persons
- Run ```kubectl apply -f ./k8s```

## Apply LocationsService
- Navigate to modulesV2/locations
- Run ```kubectl apply -f ./k8s```

## Apply ConnectionsService
- Navigate to modulesV2/connections
- Run ```kubectl apply -f ./k8s```
