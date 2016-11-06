package commands

import (
	"testing"

	"github.com/mikkeloscar/flise/backend"
	"github.com/mikkeloscar/flise/context"
)

func TestExecRun(t *testing.T) {
	ctx := context.Context(map[string]interface{}{
		"backend": backend.Mock{},
	})
	exec := Exec("")

	err := exec.Exec(ctx)
	if err != nil {
		t.Errorf("failed to execute command: %s", err)
	}
}

func TestExecString(t *testing.T) {
	cmd := "command"
	exec := Exec(cmd)

	if exec.String() != cmd {
		t.Errorf("expected %s, got %s", cmd, exec.String())
	}
}

func TestParseExec(t *testing.T) {
	cmds := map[string]bool{
		"command":               true,
		"command,":              true,
		"command,command2":      true,
		`"command with spaces"`: true,
		"": false,
	}

	for cmd, valid := range cmds {
		lexer := lex(cmd)
		_, err := parseExec(lexer)
		if err != nil && valid {
			t.Errorf("parsing '%s' should not fail", cmd)
		}

		if err == nil && !valid {
			t.Errorf("parsing '%s' should not succeed", cmd)
		}
	}
}