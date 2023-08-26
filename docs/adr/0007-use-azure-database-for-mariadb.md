# 7. Use Azure Database for MariaDB

Date: 2023-08-23

## Status

Accepted

## Context

In the process of developing an e-commerce backend using Golang, the choice of a primary database is a pivotal architectural decision. The primary database will not only store critical business data but also play a crucial role in ensuring scalability, performance, and reliability as the project evolves into a large-scale application. As a learning opportunity, it is important to make a decision that allows for hands-on experience with modern cloud technologies while setting a strong foundation for future growth.

## Decision

To address the database selection for the e-commerce backend project, the decision has been made to use Azure Database for MariaDB as the primary database solution. Azure Database for MariaDB is a fully managed database service provided by Microsoft Azure, offering high availability, scalability, and security for MariaDB deployments.

Key reasons for this decision:

**Managed Service**: Azure Database for MariaDB provides a fully managed database service, reducing the operational overhead of database administration, backup, and maintenance. This allows the development team to focus more on application features rather than managing the underlying database infrastructure.

**Scalability**: As the e-commerce backend evolves, the database workload might experience varying levels of demand. Azure Database for MariaDB offers horizontal and vertical scaling options, enabling the application to handle increased traffic and data growth without major architectural changes.

**High Availability and Disaster Recovery**: Azure's infrastructure ensures high availability through features like automatic failover and geo-replication. This helps to minimize downtime in case of failures and provides a robust disaster recovery strategy.

**Security Features**: Azure Database for MariaDB offers advanced security features such as built-in threat detection, data encryption, and compliance certifications, ensuring the protection of sensitive customer data.

## Consequences

### Advantages

**Simplified Management**: By using a managed database service, the development team can offload routine maintenance tasks, allowing more time to be dedicated to developing application features.

**Scalability**: Azure Database for MariaDB's scalability options provide the flexibility to handle growing user loads and larger datasets effectively.

**High Availability**: The automatic failover and replication capabilities of Azure Database for MariaDB enhance the application's availability and reliability.

**Security**: Built-in security features help safeguard customer data and ensure compliance with industry standards.

### Challenges and Mitigations

**Learning Curve**: While Azure Database for MariaDB abstracts much of the database management complexity, there might be a learning curve for the development team to become proficient in using Azure services effectively. This can be mitigated through training, documentation, and hands-on experience.

**Vendor Lock-in**: Choosing a specific cloud provider's managed service can introduce vendor lock-in. To mitigate this, the project can ensure that database-related code is abstracted and encapsulated, making it easier to migrate to other database solutions in the future if needed.

### Summary

By selecting Azure Database for MariaDB as the primary database solution for the e-commerce backend project, the development team aims to leverage the benefits of a managed database service, scalability, high availability, and security. This decision aligns with the project's goals of providing a learning opportunity while setting a solid foundation for future growth. While there might be challenges in terms of learning and potential vendor lock-in, these challenges can be addressed through proper training and architectural considerations.
