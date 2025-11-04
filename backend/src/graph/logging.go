package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
)

type LoggingExtension struct{}

func (l LoggingExtension) ExtensionName() string {
	return "LoggingExtension"
}
func (l LoggingExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}
func (l LoggingExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	zap.L().Info("GraphQL "+string(rc.Operation.Operation)+": "+rc.OperationName, zap.String("auth_header", rc.Headers.Get("Authorization")), zap.Any("variables", rc.Variables))
	return next(ctx)
}
