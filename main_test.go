package main

import (
	"io/ioutil"
	"log"
	"net"
	"testing"
)

func init() {
	go func() {
		tcpServer("localhost:2033")
	}()
	go func() {
		tcpServer("[0:0:0:0:0:0:0:1]:2033")
	}()
}

func tcpServer(address string) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		_, err = ioutil.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestDialWithFormat(t *testing.T) {
	type args struct {
		host string
		port string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"DialWithFormat", args{"0:0:0:0:0:0:0:1", "2033"}, false},
		{"DialWithFormat", args{"localhost", "2033"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			if err := DialWithFormat(tt.args.host, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("DialWithFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDialWithNetJoin(t *testing.T) {
	type args struct {
		host string
		port string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"DialWithFormat", args{"0:0:0:0:0:0:0:1", "2033"}, false},
		{"DialWithFormat", args{"localhost", "2033"}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := DialWithNetJoin(tt.args.host, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("DialWithNetJoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
