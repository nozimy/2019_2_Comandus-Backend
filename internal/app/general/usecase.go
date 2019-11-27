package general

import (
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/app/general/delivery/grpc/auth_grpc"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
)

type Usecase interface {
	CreateUser(*model.User) (*auth_grpc.User, error)
	VerifyUser(*model.User) (int64, error)
}