## Project structure

```
library_management/
├── bin/
├── controllers/
│   └── task_controller.go
├── models/
│   └── user.go
│   └── task.go
├── data/
│   └── db_connection.go
│   └── task_service.go
│   └── user_service.go
├── docs/
│   └── documentation.md
├── middleware/
│   └── auth_middleware.go
├── router/
│   └── router.go
├── tmp/
├── go.mod
├── go.sum
├── main.go
├── .air.toml

```

## Overview

This project is an API for a task manager application that has all the needed endpoints for creating tasks, getting all tasks, updating tasks, deleting tasks, and getting a single task by ID. It uses MongoDB for data storage and supports two roles: admin and user. Additionally, it uses JWT for authentication and authorization.

## Setup and Installation

### Installation

1. Clone the repository:

```sh
git clone https://github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager.git
cd A2SV-Backend-Track/Track3

```

2. Fetch the dependencies:

```sh
go get ./...

```

3. Initialize the project:

```sh
go mod tidy

```
4. Before running the project monog db installation is needed since this api works with a local database

- For linux installation : https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-ubuntu/
- For windows installation : https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-windows/
- For macOS installation : https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/

## Running the project

```sh
go run main.go

```

or if using air

```sh
air 

```

- After Monog db installation and runnig the app a mongo db database called taskManager with a collection of tasks in it will be created.

There are Five Api End Points for this backend
- Post http://localhost:4000/user/register with a user request body
- Post http://localhost:4000/user/login with a user request body
- Post http://localhost:4000/user/promote/:id - needed admin token
- Get http://localhost:4000/tasks
- Get http://localhost:4000/tasks/2
- Post http://localhost:4000/tasks with a task request body - needed admin token
- Put http://localhost:4000/tasks/2 with a task request body - needed admin token
- Delete http://localhost:4000/tasks/2 - needed admin token

Detailed Api Documentation - https://documenter.getpostman.com/view/29564648/2sA3s1nXCt