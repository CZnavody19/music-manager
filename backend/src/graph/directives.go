package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

// THIS IS A CUSTOM FILE THAT WILL NOT BE REGENERATED
// reference: https://gqlgen.com/reference/directives/

type Directives struct {
	resolver *Resolver
}

func NewDirectives(resolver *Resolver) *Directives {
	return &Directives{
		resolver: resolver,
	}
}

func (d *Directives) DiscordEnabled(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	enabled := d.resolver.Discord.IsEnabled()
	if !enabled {
		return nil, fmt.Errorf("discord service is not enabled")
	}

	return next(ctx)
}
