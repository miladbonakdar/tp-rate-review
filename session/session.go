package session

import "github.com/aws/aws-sdk-go/aws/session"

func New() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}
