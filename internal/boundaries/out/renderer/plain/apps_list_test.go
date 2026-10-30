package plain

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/Scalingo/go-scalingo/v11"
)

func TestAppsListRendererRender(t *testing.T) {
	ctx := context.Background()
	renderer := NewAppsList()
	renderer.SetData(ctx, []*scalingo.App{
		{Name: "my-app"},
		{Name: "my-app-2"},
	})

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	renderErr := renderer.Render(ctx)
	w.Close()
	os.Stdout = oldStdout

	if renderErr != nil {
		t.Fatalf("Render returned error: %v", renderErr)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatal(err)
	}

	want := "my-app\nmy-app-2\n"
	if buf.String() != want {
		t.Fatalf("got %q, want %q", buf.String(), want)
	}
}

func TestAppsListRendererRenderEmpty(t *testing.T) {
	ctx := context.Background()
	renderer := NewAppsList()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	renderErr := renderer.Render(ctx)
	w.Close()
	os.Stdout = oldStdout

	if renderErr != nil {
		t.Fatalf("Render returned error: %v", renderErr)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatal(err)
	}

	if buf.String() != "" {
		t.Fatalf("got %q, want empty output", buf.String())
	}
}
