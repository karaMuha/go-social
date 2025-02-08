# go-social
## Purpose
The purpose of this project is to practice hexagonal / clean architecture in the golang programming language, explore pros and cons of these kind of domain-centric architectures and come up with ideas to mitigate the cons.
My motivation was to try out an alternative to classic N-layered architecture which encapsules the core business logic from any dependencies.

## Overview
The application is build as a modular monolith and each module is based on ports and adapters with the domain as the fully independet core accessible via commands and queries. The core application defines ports (interfaces) that the surrounding technical necessities adapt to. Those technical necessities are divided in to driving and driven components. Think of the driving components as 'feeds the application with data' and driven components as 'is fed by the application with data'.
As the domain logic runs completely independent, it is much easier to unit test compared to the N-layered architecture.

![Diagram of ports and adapter architecture](/diagram.png)

## Takeaways
### Pros
- Code is kept very clean with clear seperation of concerns
- Domain logic is kept independent and very easy to unit test
- Easy to adapt to technological changes, e.g. if the endpoint technology is changed from REST to gRPC

### Cons
- A lot of boilerplate code
- Distributed and in parts deeply nested folder structure. In order to touch one feature from end to end the developer has to jump around a lot

## Outlook
The domain-centic architecutre makes an application not just clean but also resilient to technological changes. On the other hand this kind of application design makes it uncomfortable to adapt to changes in terms of business requirenments since the code which covers one feature is distributes throughout different folders.
A promising application design that tackles this challenge is the 'Vertical Slice Architecture' which I am exploring currently with the goal to combine both application designs.

## Running the application
In order to run the application clone the repo and run `make setup` in the root directory of the project. Make sure Docker Desktop is up an running then execute the command `make run`. This will launch a docker compose cluster consisting of the application, a postgres database and a mailhog dev server for sending/receiving mails.
Swagger is not setup yet, but feel free to explore the endpoints in the `module.go` file in each module.
Mailhog UI is accessible on `localhost:8025`.

To run unit tests execute the command `make unit_tests` and for integration tests `make integration_tests`.

## Todo
- provide Swagger documentation
- implement logging (maybe save logs in a NoSQL database)
- implement graceful shutdown
- implement cache for frequently viewed profiles
- implement file server to support images (and videos?) for content