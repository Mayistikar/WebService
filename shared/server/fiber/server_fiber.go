package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ServerFiber struct {}

func NewServer() *ServerFiber {
	return new(ServerFiber)
}

func (s *ServerFiber) Run(port string) {
	server := fiber.New()

	server.Listen(fmt.Sprintf(":%s", port))
}
