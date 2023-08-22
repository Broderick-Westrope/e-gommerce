# 2. Use VSCode Dev Container

Date: 2023-08-22

## Status

Accepted

## Context

As the sole developer on this project, I recognize the importance of ensuring a consistent development environment while working on different machines and over time. I currently face challenges with maintaining compatibility between my local setup and my development environment. These challenges hinder my productivity and make it difficult to pick up where I left off when switching between machines.

Additionally, I have a strong motivation to make this project accessible for others who may want to learn from it. I envision this project being a valuable resource for those interested in understanding the implementation and principles behind it. To achieve this, it's crucial to provide an easy and consistent way for others to set up and run the project.

## Decision

I have decided to adopt the use of Visual Studio Code (VSCode) Dev Containers for my development workflow. Dev Containers provide a standardized, containerized environment that includes all the necessary tools, dependencies, and extensions for the project. I will define the development environment in a `.devcontainer` folder within the project repository, using a Dockerfile (should one be needed) and a `devcontainer.json` configuration file. This setup will enable me to work in a consistent development environment regardless of the machine I'm using.

## Consequences

With the adoption of VSCode Dev Containers, several consequences arise:

### Advantages

**Consistency**: Using a Dev Container ensures that my development environment remains consistent across different machines, reducing compatibility issues.

**Reproducibility**: The containerized environment guarantees that I can reproduce the exact development environment over time, even if I switch machines.

**Isolation**: The development environment is isolated from the host system, minimizing the risk of conflicts with other software.

**Dependency Management**: Dependencies and tools are specified in the Dev Container configuration, reducing the need to manage them manually on each machine.

**Version Control**: The `.devcontainer` folder, along with the Dockerfile and `devcontainer.json` files, can be version-controlled along with the project code.

**Learning and Onboarding**: By using Dev Containers, the project becomes more accessible for others who want to learn from or contribute to it. Setting up the project becomes easier, enabling newcomers to get started faster.

### Challenges and Mitigations

**Learning Curve**: I might need some time to familiarize myself with using Dev Containers. I can refer to documentation and resources to ensure a smooth transition.

**Containerization Overhead**: Running development environments in containers might introduce some performance overhead. I will continuously monitor and optimize the containerization process for better performance.

**Docker Dependencies**: I need Docker installed on my local machines to use Dev Containers. I will set up Docker according to guidelines to ensure compatibility.

### Summary

By adopting VSCode Dev Containers, I anticipate simplifying my development process, enhancing consistency, and eliminating the challenges associated with maintaining different local development environments. Additionally, I look forward to making this project more accessible to others who can learn from it and contribute to it, thus fostering a collaborative and learning-friendly environment.




