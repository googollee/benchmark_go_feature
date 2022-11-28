package reflectcall

import (
	"context"
	"io"
)

func F(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", io.EOF
	}
	return "string: " + s, nil
}
