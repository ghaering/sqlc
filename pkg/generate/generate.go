// Package generate runs sqlc's code generation in process for fernglas:
// the generated files come back as a map instead of being written, and a
// failure is an error instead of os.Exit, which sqlc's own generate
// command calls.
//
// Local changes to sqlc v1.31.1 in this copy, nothing else differs:
// this package, and a recover in the per-package goroutine of
// internal/cmd/process.go, so a panic in the compiler or code generator
// fails that package instead of the process. Examples, docs, test files,
// test data and the command mains of the module were left out.
package generate

import (
	"context"
	"io"

	"github.com/sqlc-dev/sqlc/internal/cmd"
)

// Generate compiles the sqlc configuration dir/filename and returns the
// generated files by absolute path. Remote generation is off and
// process plugins stay disabled (sqlc's default); stderr receives sqlc's
// messages.
func Generate(ctx context.Context, dir, filename string, stderr io.Writer) (map[string]string, error) {
	return cmd.Generate(ctx, dir, filename, &cmd.Options{Env: cmd.Env{NoRemote: true}, Stderr: stderr})
}
