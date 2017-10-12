package server

type optFunc func(*Server) error

func Port(port int) func(s *Server) error {
	return func(s *Server) error {
		s.port = port
		return nil
	}
}
