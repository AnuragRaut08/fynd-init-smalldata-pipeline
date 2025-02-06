package project

import (
    "os"
    "path/filepath"
    "testing"
)

func TestInitProject(t *testing.T) {
    projectName := "test_project"
    defer os.RemoveAll(projectName) // Clean up after test

    if err := InitProject(projectName); err != nil {
        t.Fatalf("InitProject() error = %v", err)
    }

    // Verify main project directory exists
    if _, err := os.Stat(projectName); os.IsNotExist(err) {
        t.Errorf("Project directory %s does not exist", projectName)
    }

    // Verify asset directory exists
    assetDir := filepath.Join(projectName, "asset")
    if _, err := os.Stat(assetDir); os.IsNotExist(err) {
        t.Errorf("Asset directory %s does not exist", assetDir)
    }

    // Verify project.yml exists
    projectYmlPath := filepath.Join(projectName, "project.yml")
    if _, err := os.Stat(projectYmlPath); os.IsNotExist(err) {
        t.Errorf("File %s does not exist", projectYmlPath)
    }

    // Verify pipeline.yml exists
    pipelineYmlPath := filepath.Join(projectName, "pipeline.yml")
    if _, err := os.Stat(pipelineYmlPath); os.IsNotExist(err) {
        t.Errorf("File %s does not exist", pipelineYmlPath)
    }
}
