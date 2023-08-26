# 6. Use MariaDB as the RDBMS

Date: 2023-08-23

## Status

Accepted

## Context

The goal of this decision is to select a suitable Relational Database Management System (RDBMS) for the e-commerce backend project developed in Golang. The project aims not only to serve as a functional application but also as a learning opportunity that can eventually scale into a larger personal project. The choice of RDBMS is a critical architectural decision that will influence data storage, retrieval, and overall system performance.

## Decision

After careful consideration of various RDBMS options, we have decided to use MariaDB as the chosen RDBMS for the e-commerce backend project. MariaDB is a well-established open-source RDBMS that is known for its compatibility with MySQL and its performance optimizations. It provides a good balance between features, performance, and ease of use, making it suitable for both learning purposes and potential scalability.

## Consequences

### Advantages

**Compatibility**: MariaDB is a fork of MySQL, which means it shares a similar syntax and design. This compatibility will make it easier to transition between MariaDB and MySQL, should the need arise in the future.

**Community Support**: MariaDB has an active and growing community that contributes to its development, provides support, and creates various resources that can aid in learning and troubleshooting.

**Performance**: MariaDB has shown good performance in various benchmarks and real-world applications. Its query optimizer and storage engines contribute to efficient data retrieval, which is crucial for a growing e-commerce platform.

**Open Source**: MariaDB is open-source software, aligning with the project's open-source nature and providing transparency into its inner workings.

### Challenges and Mitigations

**Learning Curve**: If the development team is not familiar with MariaDB, there might be a learning curve to understand its specific features and optimizations. This can be mitigated by leveraging available learning resources, tutorials, and community support to quickly gain proficiency.

**Migration Concerns**: While MariaDB is MySQL-compatible, there might still be some migration considerations, especially if the project grows significantly. Careful schema design and use of standard SQL practices will help minimize migration challenges.

### Summary

By selecting MariaDB as the RDBMS for the e-commerce backend project, we anticipate benefiting from its compatibility, community support, performance, and open-source nature. While there might be a learning curve and potential migration concerns, these challenges can be addressed through available resources and careful planning. This decision aligns with the project's goal of creating a learning opportunity that can evolve into a scalable personal project.
