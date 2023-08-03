# Helpdesk Proto

#### Helpdesk is test project of helpdesk page using microservices

## Helpdesk Proto 

- Build with go 1.20.2
- Uses the [Protocol Buffers](https://protobuf.dev/)
- Uses [gRPC](https://grpc.io/)

## Helpdesk has below services:

- Proto
- [service-agents]() - to manage agents
- [service-frontend]() - to serve frontend UI of application and communicate with other services when necessary
- [service-notifications]() - to manage and send notification about events for tickets
- [service-tickets]() - to manage tickets
- [service-users]() - to manage users


## Generate proto files
To generate proto files (pb.go and _grpc.pb.go) run ./generate.sh script. 