package internal

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
)

type SmtpAddress string

type SmtpClient struct {
	*smtp.Client
	addr string
}

func NewSmtpClient(addr SmtpAddress) (*SmtpClient, error) {
	host, _, _ := net.SplitHostPort(string(addr))
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", string(addr), tlsconfig)
	if err != nil {
		return nil, fmt.Errorf("smtp client init failed: %v", err)
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, fmt.Errorf("smt client init failed: %v", err)
	}
	return &SmtpClient{
		client,
		string(addr),
	}, nil
}

func (sl *SmtpClient) Address() string {
	return sl.addr
}
