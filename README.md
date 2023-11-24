# Todo

- [ ] Validator
- [ ] DB?
- [x] Authorization -> Casbin
- [x] Enum
- [x] Common Functions
- [ ] Storage / S3 Protocols
- [ ] Worker Abstraction (Pool, Internal CRON) -> pond & gocron
- [ ] Local Feature Flag
- [ ] Excel

# Layer Definition

- Repository: A repository responsible for retrieving data.
  - Consists of database, cache, external client, etc.
- Service: A service responsible for doing business logic.
  - Consists of another service and repository.
- Controller: A controller responsible for handling request and response.
  - Consists of services and request handlers.
- Module: A module is a collection of components that are grouped together to provide a specific set of 
functionality within the application.

# Test Command
```bash
go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html
```

# Note

Component Internal -> Install Directory (api v1, api v2)
Rest Module -> Component Module (context based)

# Improvement
- [x] dependency injection: boilerplate -> rename, refactor, improve
- [ ] graphql
- [x] rest: internal component, core, module -> refactor, rename & improve
- [ ] event-driven: test signal changes, unit testing
- [ ] rest token: refine/redesign
- [ ] authn: rest module
- [ ] authz: permission refactor, add rest to search by action/resource, permission list return value rest refactor