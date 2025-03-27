package extensions

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	appErr "github.com/ezex-io/gopkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// FormatGQLError converts internal app errors into a structured GraphQL error.
// If the error is not of type *appErr.Error, it falls back to the default presenter.
func FormatGQLError(ctx context.Context, err error) *gqlerror.Error {
	var appErr *appErr.Error
	if errors.As(err, &appErr) {
		ext := map[string]interface{}{
			"code": appErr.Code,
		}
		if len(appErr.Meta) > 0 {
			ext["meta"] = appErr.Meta
		}

		return &gqlerror.Error{
			Message:    appErr.Message,
			Extensions: ext,
		}
	}

	return graphql.DefaultErrorPresenter(ctx, err)
}
