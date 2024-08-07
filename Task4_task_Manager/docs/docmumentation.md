## Project structure

```
library_management/
├── main.go
├── controllers/
│ └── task_controller.go
├── models/
│ └── task.go
├── data/
│ └── task_service.go
├── docs/
│ └── documentation.md
├── router/
| └── router.go
└── go.mod
└── go.sum
```

## Overview

This project is an Api for a task manager application that has all the needed end point of creating tasks, getting all tasks, updating task, deleting task and getting a single task by id.

## Setup and Installation

### Installation

1. Clone the repository:

```sh
git clone https://github.com/Johna210/A2SV-Backend-Track/Track4_Task_Manager.git
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

## Running the project

```sh
go run main.go

```

There are Five Api End Points for this backend

- Get http://localhost:4000/tasks
- Get http://localhost:4000/tasks/2
- Post http://localhost:4000/tasks with a request body
- Put http://localhost:4000/tasks/2 with a request boyd
- Delete http://localhost:4000/tasks/2

Detailed Api Documentation - https://documenter.getpostman.com/view/29564648/2sA3rzJsV2