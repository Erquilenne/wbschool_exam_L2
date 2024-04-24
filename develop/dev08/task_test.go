package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestShellCd(t *testing.T) {
	// Test cd command
	cmd := exec.Command("go", "run", "task.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Send cd command
	fmt.Fprint(stdin, "cd /tmp")
	stdin.Close()

	// Read output
	output, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for command to exit
	if err := cmd.Wait(); err != nil {
		t.Fatal(err)
	}

	// Check output
	expectedOutput := "$ "
	if string(output) != expectedOutput {
		t.Errorf("Expected output '%s', got '%s'", expectedOutput, string(output))
	}
}
func TestShellEcho(t *testing.T) {
	// Test echo command
	cmd := exec.Command("go", "run", "task.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Send echo command
	fmt.Fprint(stdin, "echo Hello, World!\n")
	stdin.Close()

	// Read output
	output, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for command to exit
	if err := cmd.Wait(); err != nil {
		t.Fatal(err)
	}

	// Check output
	expectedOutput := "$ Hello, World!\n$ "
	if string(output) != expectedOutput {
		t.Errorf("Expected output '%s', got '%s'", expectedOutput, string(output))
	}
}

func TestShellPwd(t *testing.T) {
	// Test pwd command
	cmd := exec.Command("go", "run", "task.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Send pwd command
	fmt.Fprint(stdin, "pwd\n")
	stdin.Close()

	// Read output
	output, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for command to exit
	if err := cmd.Wait(); err != nil {
		t.Fatal(err)
	}

	// Check output
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := "$ " + currentDir + "\n$ "
	if string(output) != expectedOutput {
		t.Errorf("Expected output '%s', got '%s'", expectedOutput, string(output))
	}
}
func TestShellKill(t *testing.T) {
	// Test kill command
	cmd := exec.Command("go", "run", "task.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Send kill command (replace PID with a valid PID on your system)
	fmt.Fprint(stdin, "kill [PID]")
	stdin.Close()

	// Wait for command to exit
	if err := cmd.Wait(); err != nil {
		t.Fatal(err)
	}
	// No need to check output for kill command
}

func TestShellPS(t *testing.T) {
	// Test ps command
	cmd := exec.Command("go", "run", "task.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Send ps command
	fmt.Fprint(stdin, "ps")
	stdin.Close()

	// Read output
	_, err = io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for command to exit
	if err := cmd.Wait(); err != nil {
		t.Fatal(err)
	}

	// No need to check specific output for ps command
}

func TestShellNC(t *testing.T) {
	// Test nc command
	cmd := exec.Command("go", "run", "task.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Send nc command (replace host and port with valid values)
	fmt.Fprint(stdin, "nc [host] [port]\n")
	stdin.Close()

	// No need to check output for nc command
}
