//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos-layout/internal/helloworld/domain"
	"github.com/go-kratos/kratos-layout/internal/helloworld/conf"
	"github.com/go-kratos/kratos-layout/internal/helloworld/repo"
	"github.com/go-kratos/kratos-layout/internal/helloworld/server"
	"github.com/go-kratos/kratos-layout/internal/helloworld/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, repo.ProviderSet, domain.ProviderSet, service.ProviderSet, newApp))
}
