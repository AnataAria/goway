package http

func (s *Server) Start() {
	s.fuego.Run()
}
