package helper

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type GraphQLError struct {
	Message error
	ErrCode string
}

type GError interface {
	Error() string
}

func GqlError(ctx context.Context, err *GraphQLError) {
	graphql.AddError(ctx, &gqlerror.Error{
		Message: err.Message.Error(),
		Extensions: map[string]interface{}{
			"code": err.ErrCode,
		},
	})
}
