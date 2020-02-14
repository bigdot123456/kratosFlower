// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"aflower/internal/dao"
	"aflower/internal/service"
	"aflower/internal/server/grpc"
	"aflower/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
