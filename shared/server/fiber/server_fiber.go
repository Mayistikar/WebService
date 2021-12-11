package fiber

type ServerFiber struct {

}

func NewServer() *ServerFiber {
	return new(ServerFiber)
}

func (s *ServerFiber) Run(port string) {}
