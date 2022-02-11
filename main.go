package main

import (
	"fmt"
	"io"
	"net"

	"github.com/gliderlabs/ssh"
	"github.com/pires/go-proxyproto"
	gossh "golang.org/x/crypto/ssh"
)

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:2222")
	listener := &proxyproto.Listener{Listener: l}
	defer listener.Close()

	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	})

	go ssh.Serve(listener, nil)

	// basic client
	cliCfg := &gossh.ClientConfig{
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		User:            "r00t",
		Auth: []gossh.AuthMethod{
			gossh.Password("53cr3t"),
		},
	}
	gossh.Dial("tcp", listener.Addr().String(), cliCfg)
}
