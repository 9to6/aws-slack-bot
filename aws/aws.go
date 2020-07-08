package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/9to6/aws-slack-bot/config"
)

type Sessions []*session.Session

func NewSessions(cfg *config.Config) (Sessions, error) {
	sessions := Sessions{}
	for _, a := range cfg.Aws {
		session, err := newSession(a.Region, a.AccessKey, a.SecretAccessKey)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func newSession(region, accessKey, secretAccessKey string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretAccessKey, ""),
	}

	options := session.Options{
		Config: *awsConfig,
	}

	sess, err := session.NewSessionWithOptions(options)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func NewEc2WithRole(role string, cfg *config.Config) (*ec2.EC2, error) {
	sess := session.Must(session.NewSession())

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN.
	creds := stscreds.NewCredentials(sess, role)

	// Create service client value configured for credentials
	// from assumed role.
	conf := &aws.Config{
		Credentials: creds,
		Region: aws.String(cfg.Aws[0].Region),
	}
	return ec2.New(sess, conf), nil
}
