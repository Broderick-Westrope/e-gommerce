# E-Gommerce

## Description

E-Gommerce is a powerful and comprehensive Go-based web API for building scalable e-commerce applications. The primary objective of this project is to serve as a valuable resource for developers learning Go, providing a well-structured template for building APIs while adhering to best practices and Go idioms. The project aims to strike a balance between being approachable for beginners and offering in-depth insights into creating robust APIs.

If you feel the project is not approachable or could be improved in any way, please [open an issue](https://github.com/Broderick-Westrope/e-gommerce/issues/new/choose). I am always looking for ways to improve the project and make it more accessible to developers of all skill levels.

## Features

- Browse products and retrieve detailed product information.
- User-friendly error handling and messaging.
- Detailed API documentation using Swagger.

Coming Soon:
- Create a user account and perform user authentication.
- Favorite items for future reference.
- Add items to the cart for purchase.

## Target Audience

E-Gommerce is designed for developers who are learning Go and want to explore real-world project structures, idiomatic practices, and essential considerations when building scalable APIs. While not intended as a standalone e-commerce site, it serves as a starting point for those interested in learning how to create feature-rich APIs.

## Technologies Used

- Routing: [Chi](https://go-chi.io/)
- Middleware: [httprate](https://github.com/go-chi/httprate), Logger, Heartbeat, CleanPath, AllowContentType, Recoverer, RedirectSlashes, Limit (See [Chi Middleware](https://go-chi.io/#/pages/middleware))
- Database: [MariaDB](https://mariadb.org/) with [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql).
- Environment Variables: [joho/godotenv](https://github.com/joho/godotenv)
- Swagger Documentation: [swaggo/swag](https://github.com/swaggo/swag) and [swaggo/http-swagger](https://github.com/swaggo/http-swagger)
- [Task Automation](#task-automation): [Taskfile](https://taskfile.dev/)
- Linting: [golangci-lint](https://golangci-lint.run/)
- [Live Reloading](#live-reloading): [Air](https://github.com/cosmtrek/air)

## Infrastructure

- Database: [Azure Database for MariaDB](https://learn.microsoft.com/en-us/azure/mariadb/)

## Installation

To run the project locally, follow these steps:

1. Clone the repository: `git clone https://github.com/YourUsername/e-gommerce.git`
2. Navigate to the project directory: `cd e-gommerce`
3. Install [Go](https://go.dev/doc/install), and [Task](https://taskfile.dev/) (optional, but recommended for [task automation](#task-automation))
4. Build and run the project: `task run` or `go run .`. See [usage](#usage) for more details.

## Usage

E-Gommerce uses [Task](https://taskfile.dev/) for task automation. Alternatively, you can use [the standard Go commands](https://go.dev/doc/tutorial/getting-started) to build and run the project.

### Live Reloading

To enable live reloading, install [Air](https://github.com/cosmtrek/air) and run `air` in the project directory. This will automatically rebuild and restart the project when changes are detected. It is configured to put the build in the `tmp` directory, which is ignored by Git and destroyed when it stops running. Configuration for Air is stored in the [`.air.toml`](./.air.toml) file.

### Task Automation

Available tasks include:
- `task run`: Builds and runs the project.
- `task build`: Creates a build of the project.
- `task lint`: Runs [golangci-lint](https://golangci-lint.run/) for code linting.
- `task test`: Runs tests with coverage.
- `task swag`: Generates [Swagger](https://swagger.io/) documentation using [swag](https://github.com/swaggo/swag).
- `task pcc`: Runs Pre-Commit Checks (PCCs). This performs linting, testing, and Swagger documentation generation in the correct order.

Configuration for Task is stored in the [Taskfile.yml](./Taskfile.yml) file.

### Manual

- Build: `mkdir -p build` then `go build -o ./build/ .`
- Test: `go test ./... -cover`
- Lint: `golangci-lint run`
- Generate [Swagger](https://swagger.io/) Docs: `swag init -o './api' -g './cmd/web/server.go' --parseDependency`, then `swag fmt` to format

## Contributing

E-Gommerce welcomes contributions from the community. You can contribute by:

- Providing feedback and suggestions
- Implementing new features and functionalities
- Enhancing project documentation
- Assisting with testing and quality assurance
- Improving project structure, performance, and security

To contribute:

1. Fork the repository using the fork button at the top of the page or by running `git clone https://github.com/Broderick-Westrope/e-gommerce`.
2. Create a new branch for your feature or improvement: `git checkout -b feature/<your-feature-name>`
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Create a pull request to the main repository.

If you are unsure about any of these please reach out to me at [broderickwestrope@gmail.com](mailto:broderickwestrope@gmail.com).

## Future Plans

- Implement [Vue 3](https://vuejs.org/) frontend for a user-friendly browsing experience.
- Implement authentication and authorization mechanisms.
- Finalize core functionalities for the cart, favorites, and user actions.
- Implement automated deployment strategies.
- Establish a vibrant open-source community around the project.
- Create educational blog posts and documentation to guide developers in using and contributing to the project, and replicating its features.

## Documentation

For detailed information about API endpoints and usage, please refer to [the Swagger Documentation](./api/).

## Contact

For any questions, suggestions, or collaboration opportunities, you can reach out to the project owner, Broderick Westrope, at [broderickwestrope@gmail.com](mailto:broderickwestrope@gmail.com).

## License

E-Gommerce is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html). Please review the license for detailed terms and conditions.

