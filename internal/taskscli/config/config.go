package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	defaultTasksDir    = ".forge/tasks"
	defaultStateDir    = ".forge/state"
	defaultDBFile      = "forge.db"
	defaultProjectName = "forge"
	defaultContext     = "dev"
)

// Config carries resolved absolute paths for CLI operations.
type Config struct {
	RepoPath       string
	TasksDir       string
	StateDir       string
	DBPath         string
	TasksDirRel    string
	StateDirRel    string
	DBFile         string
	ProjectName    string
	DefaultContext string
}

type rootConfig struct {
	Project struct {
		Name string `yaml:"name"`
	} `yaml:"project"`
	Tasks struct {
		TasksDir string `yaml:"tasks_dir"`
		StateDir string `yaml:"state_dir"`
		DBFile   string `yaml:"db_file"`
	} `yaml:"tasks"`
}

// Load reads forge.yaml at repoPath and returns resolved paths.
func Load(repoPath string) (*Config, error) {
	configPath := filepath.Join(repoPath, "forge.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read forge.yaml: %w", err)
	}

	var rc rootConfig
	if err := yaml.Unmarshal(data, &rc); err != nil {
		return nil, fmt.Errorf("failed to parse forge.yaml: %w", err)
	}

	tasksDirRel := rc.Tasks.TasksDir
	if tasksDirRel == "" {
		tasksDirRel = defaultTasksDir
	}
	stateDirRel := rc.Tasks.StateDir
	if stateDirRel == "" {
		stateDirRel = defaultStateDir
	}
	dbFile := rc.Tasks.DBFile
	if dbFile == "" {
		dbFile = defaultDBFile
	}

	tasksDir := filepath.Join(repoPath, filepath.Clean(tasksDirRel))
	stateDir := filepath.Join(repoPath, filepath.Clean(stateDirRel))
	dbPath := filepath.Join(stateDir, dbFile)

	return &Config{
		RepoPath:       repoPath,
		TasksDir:       tasksDir,
		StateDir:       stateDir,
		DBPath:         dbPath,
		TasksDirRel:    tasksDirRel,
		StateDirRel:    stateDirRel,
		DBFile:         dbFile,
		ProjectName:    projectName(rc.Project.Name),
		DefaultContext: defaultContext,
	}, nil
}

func projectName(name string) string {
	if strings.TrimSpace(name) == "" {
		return defaultProjectName
	}
	return name
}
