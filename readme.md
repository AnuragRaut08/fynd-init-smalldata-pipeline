
# fynd-init-smalldata-pipeline

fynd-init-smalldata-pipeline` is a command-line tool designed to initialize a standardized directory structure for small data pipeline projects.t streamlines the setup process by creating necessary directories and configuration files, providing a consistent starting point for your projects.
## Features

- reates a main project directory with a specified name.- enerates an `asset/` subdirectory for storing assets.- reates a `project.yml` file containing the project's name and version.- enerates an empty `pipeline.yml` file for future pipeline configurations.- rovides a Terminal User Interface (TUI) summary of the initialization process using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework.
## Prerequisites

- Go](https://go.dev/dl/) 1.16 or later installed on your system.
## Installation

1. **Clone the Repository:**

   ``bash
   git clone https://github.com/AnuragRaut08/fynd-init-smalldata-pipeline.git
   cd fynd-init-smalldata-pipeline
   ```
2. **Initialize Go Modules:**

   ``bash
   go mod tidy
   ```
3. **Build the Application:**

   ``bash
   go build -o fynd-init-smalldata-pipeline ./cmd/fynd-init-smalldata-pipeline
   ```
   his command compiles the application and creates an executable named `fynd-init-smalldata-pipeline` in your current directory.
## Usage

un the application with the desired project name as an argument:
``bash
./fynd-init-smalldata-pipeline <project_name>
```
eplace `<project_name>` with your desired project name.
**Example:**

``bash
./fynd-init-smalldata-pipeline project123
```
## Expected Output

pon execution, the application initializes the following directory structure:
``
project123/
├── asset/
├── project.yml
└── pipeline.yml
```
he Terminal User Interface (TUI) will display a summary of the initialization process:
``
Project 'project123' has been initialized successfully.

The following items were created:
- project123/
- project123/asset/
- project123/project.yml
- project123/pipeline.yml

Press 'q' or 'esc' to exit.
```
