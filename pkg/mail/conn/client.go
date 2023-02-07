package conn

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
)

type SmtpAddress string

type SmtpClient struct {
	*smtp.Client
	hostname string
}

// best way to implement singleton?
// What if the first client will close function, while the rest will still be using it?
// Maybe it can use some kind of counter, which if has the value of 0 just closes the connection?
// That would give a lot of boilerplate code
var smtpClientSingleton *SmtpClient

func NewSmtpClient(addr SmtpAddress) (*SmtpClient, func(), error) {
	if smtpClientSingleton != nil {
		return smtpClientSingleton, func() {
			return
		}, nil
	}

	host, _, _ := net.SplitHostPort(string(addr))
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", string(addr), tlsconfig)
	if err != nil {
		return nil, nil, fmt.Errorf("smtp client init failed: %v", err)
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, nil, fmt.Errorf("smt client init failed: %v", err)
	}
	cleanup := func() {
		client.Close()
	}

	smtpClientSingleton = &SmtpClient{
		client,
		string(addr),
	}
	return smtpClientSingleton, cleanup, nil
}

func (sl *SmtpClient) Host() string {
	return sl.hostname
}
