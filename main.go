package main

import (
    "fmt"
    "os"
    "fynd-init-smalldata-pipeline/cmd/project"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: fynd-init-smalldata-pipeline <project_name>")
        return
    }

    projectName := os.Args[1]

    if err := project.InitProject(projectName); err != nil {
        fmt.Printf("Failed to initialize project: %v\n", err)
        os.Exit(1)
    }
}
