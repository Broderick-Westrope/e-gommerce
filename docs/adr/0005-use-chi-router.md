# 5. Use Chi Router

Date: 2023-08-23

## Status

Accepted

## Context

In the process of developing an e-commerce backend using Golang, the choice of router plays a crucial role in shaping the project's architecture and scalability. This architecture decision record outlines the rationale for adopting the [Chi router](https://go-chi.io/) for the project.

## Decision

We have decided to use [Chi](https://go-chi.io/) as the routing framework for our Golang-based e-commerce backend project. It is a lightweight, fast, and flexible router that aligns well with the goals of our project, which is both a learning opportunity and a potential candidate for future growth into a large-scale application.

One of the significant advantages of Chi is its extensive support for middleware, which can be built using the standard library `net/http` package. This compatibility allows us to leverage existing knowledge and resources while incorporating middleware for functionalities like authentication, logging, and request/response manipulation.

Furthermore, Chi minimizes external dependencies by relying primarily on the standard library. This design choice is aligned with our goal of keeping the project lightweight and reducing the risk of compatibility issues stemming from external libraries.

It also offers several features that enhance our development process:

**Method-Based Routing**: It provides built-in support for method-based routing, allowing us to define different handlers for different HTTP methods (GET, POST, PUT, DELETE, etc.). This enhances code organization and readability, making it clear which functions handle specific HTTP actions.

**Query Parameter Handling**: It offers an easy and intuitive way to access query parameters from incoming requests. This simplifies the process of extracting data from URLs, which is crucial for implementing search, filtering, and pagination features in our e-commerce backend.

**Clean Syntax for Grouping Routes**: It has a clean and straightforward syntax for grouping related routes. This makes it easier to manage and maintain routes that belong to the same logical group, such as authentication, product management, and order processing.

## Consequences

### Advantages

**Performance**: Chi is known for its exceptional performance due to its minimalistic design and efficient routing mechanisms. This will ensure that our e-commerce backend can handle a high volume of requests without introducing significant latency.

**Flexibility**: Chi provides a modular approach to middleware, making it easier to incorporate additional functionality such as authentication, logging, and request/response manipulation as our project evolves.

**Learning Opportunity**: By using Chi, team members, especially those newer to Golang, will have the chance to learn and work with a popular router framework. The concepts and practices learned through this experience can be valuable for their skill growth.

**Scalability**: Chi's design allows for efficient handling of routes and middleware, which is crucial for ensuring the backend can scale gracefully as the e-commerce platform attracts more users.

### Challenges and Mitigations

**Learning Curve**: While Chi is lightweight, it might still have a learning curve for those unfamiliar with routing concepts. To mitigate this, we will provide documentation, tutorials, and possibly pair programming sessions to help team members get up to speed.

**Community and Support**: Chi, although widely used, might have a smaller community compared to some other router options. To address this, we will actively participate in relevant online communities and forums, and contribute back to the open-source project if possible.

### Summary

By adopting Chi for our Golang-based e-commerce backend, we anticipate smoother handling of routes, efficient middleware integration, and overall better performance. The built-in support for method-based routing, easy query parameter handling, and clean syntax for grouping routes will contribute to a more organized and maintainable codebase. The fact that Chi has extensive support for middleware built using the standard library `net/http` package is advantageous, as it allows us to capitalize on existing knowledge. Additionally, the minimal external dependencies of Chi align well with our goal of maintaining a lightweight project. While there might be a learning curve and potential community challenges, the benefits in terms of scalability, flexibility, and learning outweigh the drawbacks. This decision aligns with our goal of creating a learning opportunity that can eventually grow into a large-scale project.
