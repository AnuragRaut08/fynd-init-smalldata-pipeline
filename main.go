package main

import (
    "fmt"
    "os"
    "path/filepath"

    
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: fynd-init-smalldata-pipeline <project_name>")
        return
    }

    projectName := os.Args[1]

    // Create the main project directory
    err := os.MkdirAll(projectName, 0755)
    if err != nil {
        fmt.Printf("Error creating project directory: %v\n", err)
        return
    }

    // Create the asset subdirectory
    assetDir := filepath.Join(projectName, "asset")
    err = os.MkdirAll(assetDir, 0755)
    if err != nil {
        fmt.Printf("Error creating asset directory: %v\n", err)
        return
    }

    // Create project.yml with name and version
    projectYmlPath := filepath.Join(projectName, "project.yml")
    projectYmlContent := fmt.Sprintf("name: %s\nversion: 1.0.0\n", projectName)
    err = os.WriteFile(projectYmlPath, []byte(projectYmlContent), 0644)
    if err != nil {
        fmt.Printf("Error creating project.yml: %v\n", err)
        return
    }

    // Create an empty pipeline.yml file
    pipelineYmlPath := filepath.Join(projectName, "pipeline.yml")
    err = os.WriteFile(pipelineYmlPath, []byte{}, 0644)
    if err != nil {
        fmt.Printf("Error creating pipeline.yml: %v\n", err)
        return
    }

    fmt.Printf("Project '%s' has been initialized successfully.\n", projectName)
}
