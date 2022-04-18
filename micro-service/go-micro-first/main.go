package go_micro_first

import (
	"github.com/asim/go-micro/v3"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)
}
