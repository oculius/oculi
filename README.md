# Todo

- [ ] Excel
- [ ] Validator
- [ ] Storage / S3 Protocols
- [ ] Worker Abstraction (Pool, Internal CRON) -> pond & gocron
- [ ] Local Feature Flag
- [ ] DB?
- [x] Authorization -> Casbin
- [x] Enum
- [x] Common Functions

# Layer Definition

- Repository: A repository responsible for retrieving data.
  - Consists of database, cache, external client, etc.
- Service: A service responsible for doing business logic.
  - Consists of another service and repository.
- Controller: A controller responsible for handling request and response.
  - Consists of services and request handlers.
- Component:
- Module: A module is a collection of components that are grouped together to provide a specific set of 
functionality within the application.

# Test Command
```bash
go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html
```