package freeling

import (
	"io"
	"os/exec"
)

type Server struct {
	cmd *exec.Cmd
}

func NewServer(logFile io.Writer) (*Server, error) {
	cmd := exec.Command("analyze", "-f", "es.cfg", "--flush", "--output", "json", "--server", "--port", "50005")

	if logFile != nil {
		cmd.Stdout = logFile
		cmd.Stderr = logFile
	}

	return &Server{cmd}, nil
}

func (s *Server) Start() error {
	return s.cmd.Start()
}

func (s *Server) Stop() error {
	return s.cmd.Process.Kill()
}
