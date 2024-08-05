## Project structure

library_management/
├── main.go
├── controllers/
│ └── library_controller.go
├── models/
│ └── book.go
│ └── member.go
├── services/
│ └── library_manager.go
| └── library_service.go
├── docs/
│ └── documentation.md
└── go.mod

## Overview

This project is a console-based library management system implemented in Go. It allows users to add, remove, borrow, and return books, as well as list available and borrowed books. The system uses a simple command-line interface for interaction.

## Setup and Installation

### Installation

1. Clone the repository:

```sh
git clone https://github.com/Johna210/A2SV-Backend-Track/Track3.git
cd backend_assessment/Track3

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
