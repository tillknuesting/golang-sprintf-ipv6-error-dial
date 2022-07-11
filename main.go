package main

import (
	"fmt"
	"net"
)

func DialWithFormat(host string, port string) error {
	connectionString := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.Dial("tcp", connectionString)
	if err != nil {
		return fmt.Errorf("error while dialing: %w", err)
	}
	defer conn.Close()

	return nil
}

func DialWithNetJoin(host string, port string) error {
	connectionString := net.JoinHostPort(host, port)

	conn, err := net.Dial("tcp", connectionString)
	if err != nil {
		return fmt.Errorf("error while dialing: %w", err)
	}
	defer conn.Close()

	return nil
}
