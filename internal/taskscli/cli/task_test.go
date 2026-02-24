package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/hollis-labs/forge/internal/taskscli/config"
	"github.com/hollis-labs/forge/internal/taskscli/index"
)

const testConfig = `version: 1
project:
  name: forge
`

func TestCreateStartDoneFlow(t *testing.T) {
	repo := setupRepo(t)
	cfg, err := config.Load(repo)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	runner := newRunner(cfg)
	runner.WithStdout(ioDiscard{})
	runner.WithStderr(ioDiscard{})
	fixed := time.Date(2026, 2, 24, 0, 0, 0, 0, time.UTC)
	runner.WithNow(func() time.Time { return fixed })

	if err := runner.Dispatch([]string{"create", "Test task"}); err != nil {
		t.Fatalf("create: %v", err)
	}

	taskPath := filepath.Join(repo, ".forge", "tasks", "TASK-20260224-001.md")
	data, err := os.ReadFile(taskPath)
	if err != nil {
		t.Fatalf("read task: %v", err)
	}
	if !bytes.Contains(data, []byte("status: todo")) {
		t.Fatalf("expected todo status:\n%s", data)
	}

	if err := runner.Dispatch([]string{"start", "TASK-20260224-001"}); err != nil {
		t.Fatalf("start: %v", err)
	}
	if err := runner.Dispatch([]string{"done", "TASK-20260224-001"}); err != nil {
		t.Fatalf("done: %v", err)
	}

	data, _ = os.ReadFile(taskPath)
	if !bytes.Contains(data, []byte("status: done")) {
		t.Fatalf("expected done status:\n%s", data)
	}

	store, err := index.Open(context.Background(), cfg.DBPath)
	if err != nil {
		t.Fatalf("open index: %v", err)
	}
	defer store.Close()
	rows, err := store.ListTasks(context.Background(), index.Filters{})
	if err != nil || len(rows) != 1 || rows[0].Status != "done" {
		t.Fatalf("unexpected rows: %+v err=%v", rows, err)
	}
}

func TestReindexAndList(t *testing.T) {
	repo := setupRepo(t)
	cfg, err := config.Load(repo)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}
	runner := newRunner(cfg)
	runner.WithStdout(ioDiscard{})
	runner.WithStderr(ioDiscard{})
	fixed := time.Date(2026, 2, 24, 0, 0, 0, 0, time.UTC)
	runner.WithNow(func() time.Time { return fixed })

	if err := runner.Dispatch([]string{"create", "First task"}); err != nil {
		t.Fatalf("create first: %v", err)
	}
	if err := runner.Dispatch([]string{"create", "Second task"}); err != nil {
		t.Fatalf("create second: %v", err)
	}

	os.Remove(cfg.DBPath)
	if err := runner.Dispatch([]string{"reindex"}); err != nil {
		t.Fatalf("reindex: %v", err)
	}

	var out bytes.Buffer
	runner.WithStdout(&out)
	if err := runner.Dispatch([]string{"list"}); err != nil {
		t.Fatalf("list: %v", err)
	}
	text := out.String()
	if !strings.Contains(text, "TASK-20260224-001") || !strings.Contains(text, "TASK-20260224-002") {
		t.Fatalf("list output missing tasks:\n%s", text)
	}
}

type ioDiscard struct{}

func (ioDiscard) Write(p []byte) (int, error) { return len(p), nil }

func setupRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "forge.yaml"), []byte(testConfig), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(dir, ".forge", "tasks"), 0o755); err != nil {
		t.Fatalf("mkdir tasks: %v", err)
	}
	return dir
}
