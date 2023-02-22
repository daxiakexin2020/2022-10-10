package server

import (
	"errors"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Email struct {
	client    *email.Email
	option    Option
	plainAuth PlainAuth
}

type Option struct {
	From    string
	To      []string
	Body    []byte
	Subject string
	Html    []byte
}

type PlainAuth struct {
	Addr     string
	Username string
	Pwd      string
	Host     string
	Identity string
}

type Opt func(e *Email)

func WithSubject(subject string) Opt {
	return func(e *Email) {
		e.option.Subject = subject
	}
}

func WithHtml(html []byte) Opt {
	return func(e *Email) {
		e.option.Html = html
	}
}

func WithIdentity(identity string) Opt {
	return func(e *Email) {
		e.plainAuth.Identity = identity
	}
}

func (e *Email) apply(opts ...Opt) {
	for _, opt := range opts {
		opt(e)
	}
}

func NewEmail(from string, to []string, body []byte, opts ...Opt) Email {
	o := Option{From: from, To: to, Body: body}
	e := Email{client: email.NewEmail(), option: o, plainAuth: PlainAuth{}}
	e.apply(opts...)
	return e
}

func (e *Email) check() error {
	if e.plainAuth.Addr == "" || e.plainAuth.Username == "" || e.plainAuth.Pwd == "" || e.plainAuth.Host == "" {
		return errors.New("sendoption config err")
	}
	return nil
}

func (e *Email) Send(auth PlainAuth) error {

	e.plainAuth = auth
	if err := e.check(); err != nil {
		return err
	}

	e.client.From = e.option.From
	e.client.To = e.option.To
	e.client.Text = e.option.Body
	if e.option.Subject != "" {
		e.client.Subject = e.option.Subject
	}
	if len(e.option.Html) > 0 {
		e.client.HTML = e.option.Html
	}
	return e.client.Send(auth.Addr, smtp.PlainAuth(auth.Identity, auth.Username, auth.Pwd, auth.Host))
}
