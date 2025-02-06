package project

import (
    "fmt"
    "os"
    "path/filepath"
)

// InitProject initializes the project directory structure.
func InitProject(projectName string) error {
    // Create the main project directory
    err := os.MkdirAll(projectName, 0755)
    if err != nil {
        return fmt.Errorf("error creating project directory: %w", err)
    }

    // Create the asset subdirectory
    assetDir := filepath.Join(projectName, "asset")
    err = os.MkdirAll(assetDir, 0755)
    if err != nil {
        return fmt.Errorf("error creating asset directory: %w", err)
    }

    // Create project.yml with name and version
    projectYmlPath := filepath.Join(projectName, "project.yml")
    projectYmlContent := fmt.Sprintf("name: %s\nversion: 1.0.0\n", projectName)
    err = os.WriteFile(projectYmlPath, []byte(projectYmlContent), 0644)
    if err != nil {
        return fmt.Errorf("error creating project.yml: %w", err)
    }

    // Create an empty pipeline.yml file
    pipelineYmlPath := filepath.Join(projectName, "pipeline.yml")
    err = os.WriteFile(pipelineYmlPath, []byte{}, 0644)
    if err != nil {
        return fmt.Errorf("error creating pipeline.yml: %w", err)
    }

    return nil
}
