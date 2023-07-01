# Tasks Hub

The Tasks Hub controls the task management process, including task management, approval email communication with approvers, and the management of the approval workflow.

This repository contains the code for the Tasks Hub along with relevant specifications and implementation examples.

## Features

- Task creation: The hub allows users to manage tasks.
- Approval workflow: Upon task creation, email are sent to all approvers with links to agree or reject the task.
- Task reconciliation: If at least one approver rejects the task, it is considered inconsistent and sent to the approval participants. If all approvers approve the task, it is considered approved and sent to all participants in the process.
- API code generation: The Tasks Hub API was generated using the "oapi-codegen" library for seamless integration and code consistency.
- Error handling: The Tasks Hub implements an in-place logging strategy using the "zapctx" library for efficient error tracking and debugging.

## Getting Started

To get started with the Tasks Hub, follow the steps below:

- Clone the repository: `git clone github.com/vbelouso/tasks`
- Run the hub: `make run-server-local`
- Build the hub: `make build/server`
- Access the hub API: `curl http://localhost:8080/v1/tasks?page=1&size=10`

For detailed instructions and examples, please refer to the documentation and code samples provided in this repository.

## Documentation

The documentation for the Tasks Hub can be found in the `docs` directory. It includes information on API endpoints, request/response schemas, and usage examples.

## License

The Tasks Hub is released under the MIT License. Feel free to use, modify, and distribute the code as per the license terms.
