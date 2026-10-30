package plain

import (
	"context"
	"fmt"
	"os"

	"github.com/Scalingo/cli/internal/boundaries/out/renderer"
	"github.com/Scalingo/go-scalingo/v11"
)

type appsListRenderer struct {
	apps []*scalingo.App
}

func NewAppsList() renderer.Renderer[[]*scalingo.App] {
	return &appsListRenderer{}
}

func (r *appsListRenderer) Render(ctx context.Context) error {
	for _, app := range r.apps {
		_, err := fmt.Fprintln(os.Stdout, app.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *appsListRenderer) SetData(ctx context.Context, apps []*scalingo.App) {
	r.apps = apps
}
