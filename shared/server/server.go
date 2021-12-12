package server

type WebServer interface {
	Run(port string)
}
