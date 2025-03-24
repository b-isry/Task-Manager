# Task Manager

Welcome to the Task Manager project! This project was developed as part of my internship at A2SV. The primary objective of this project is to manage tasks efficiently using a RESTful API built with Go and the Gin framework.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Task Manager is a web application that allows users to create, read, update, and delete tasks. It is built using Go and the Gin framework, making it a fast and efficient solution for managing tasks.

## Features

- RESTful API for task management
- Create, read, update, and delete tasks
- JSON-based API responses
- Error handling and validation
- Lightweight and efficient

## Technologies Used

- **Go**: The primary programming language used for this project.
- **Gin**: A high-performance HTTP web framework for Go, used to build the RESTful API.

## Installation

To get started with the Task Manager, follow these steps:

1. **Clone the repository:**
    ```bash
    git clone https://github.com/b-isry/Task-Manager.git
    cd Task-Manager
    ```

2. **Install dependencies:**
    Make sure you have Go installed on your machine. Then, install the required Go packages:
    ```bash
    go mod tidy
    ```

3. **Run the application:**
    ```bash
    go run main.go
    ```

## Usage

Once the application is running, you can interact with the API using tools like `curl` or Postman. Below are some example API endpoints:

- **Create a task:**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"title": "New Task", "description": "Task description"}' http://localhost:8080/tasks
    ```

- **Get all tasks:**
    ```bash
    curl -X GET http://localhost:8080/tasks
    ```

- **Get a specific task:**
    ```bash
    curl -X GET http://localhost:8080/tasks/{task_id}
    ```

- **Update a task:**
    ```bash
    curl -X PUT -H "Content-Type: application/json" -d '{"title": "Updated Task", "description": "Updated description"}' http://localhost:8080/tasks/{task_id}
    ```

- **Delete a task:**
    ```bash
    curl -X DELETE http://localhost:8080/tasks/{task_id}
    ```

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
