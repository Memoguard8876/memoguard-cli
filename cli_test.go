package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestBlockAndRedact(t *testing.T) {
	var output, errors bytes.Buffer
	code := Run(context.Background(), []string{"scan", "--memo", "person@example.com", "--format", "json"}, strings.NewReader(""), &output, &errors)
	if code != ExitBlocked {
		t.Fatalf("exit=%d stderr=%s", code, errors.String())
	}
	if strings.Contains(output.String(), "person@example.com") || !strings.Contains(output.String(), "transaction.memo.text") {
		t.Fatalf("unsafe or incomplete output: %s", output.String())
	}
}

func TestClean(t *testing.T) {
	var output, errors bytes.Buffer
	code := Run(context.Background(), []string{"scan", "--memo", "reference-123"}, strings.NewReader(""), &output, &errors)
	if code != ExitClean || !strings.Contains(output.String(), "clean") {
		t.Fatalf("exit=%d stdout=%s stderr=%s", code, output.String(), errors.String())
	}
}

func TestInvalidInput(t *testing.T) {
	var output, errors bytes.Buffer
	code := Run(context.Background(), []string{"scan", "--memo", "test", "--xdr", "file"}, strings.NewReader(""), &output, &errors)
	if code != ExitInput {
		t.Fatalf("exit=%d", code)
	}
}
