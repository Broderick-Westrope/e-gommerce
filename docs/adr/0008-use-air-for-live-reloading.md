# 8. Use `air` for Live Reloading

Date: 2023-08-29

## Status

Accepted

## Context

Our Golang API project is being developed and maintained, and we are looking for ways to streamline our development process. One of the pain points identified is the need to manually stop, rebuild, and restart the application whenever code changes are made. This workflow slows down development and increases the risk of human error.

## Decision

We will incorporate the use of [air](https://github.com/cosmtrek/air), a live reloading tool for Golang applications, into our development workflow. [air](https://github.com/cosmtrek/air) monitors the project directory for changes and automatically rebuilds and restarts the application whenever changes are detected.

## Consequences

### Advantages

**Faster development cycle**: Developers no longer need to manually rebuild and restart the application after each code change, resulting in increased productivity.

**Immediate feedback**: Code changes are reflected in the running application without manual intervention, allowing developers to see the effects of their changes in real time.

### Challenges and Mitigations

**Potential resource consumption**: Running [air](https://github.com/cosmtrek/air) may consume additional system resources due to continuous monitoring and rebuilding. However, this is manageable through system resource monitoring and allocation.

**Integration with existing build processes**: We need to ensure that the use of [air](https://github.com/cosmtrek/air) does not conflict with our existing build and deployment pipelines. We'll need to carefully integrate [air](https://github.com/cosmtrek/air) into our development process.

### Summary

By adopting [air](https://github.com/cosmtrek/air) for live reloading in our Golang API project, we expect to significantly reduce development cycle times and improve the overall development experience. While there might be challenges in terms of resource usage and integration, these can be mitigated with proper monitoring and integration efforts.
