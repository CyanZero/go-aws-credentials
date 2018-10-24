package sessionutil

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

type Session session.Session

func GetSession(awsProfile string, strSerialNumber string, strTokenCode string) *session.Session {
	var sess *session.Session
	// Specify profile to load for the session's config
	if awsProfile != "" {
		sess, _ = session.NewSessionWithOptions(session.Options{
			Profile: awsProfile,
		})
	} else if strSerialNumber != "" {
		sess = GetTemporySessionWithMFA(strSerialNumber, strTokenCode)
	}

	if sess == nil {
		log.Panic("AWS credentials are not valid, please check ~/.aws/config and ~/.aws/credentials")
		return nil
	} else {
		log.Println("A new connection session is created and ready for use")
	}

	return sess
}

func GetTemporySessionWithMFA(strSerialNumber string, strTokenCode string) *session.Session {
	svc := sts.New(session.New())
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(3600),
		SerialNumber:    aws.String(strSerialNumber),
		TokenCode:       aws.String(strTokenCode),
	}

	result, err := svc.GetSessionToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				log.Fatal(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				log.Fatal(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Fatal(err.Error())
		}
		return nil
	}

	fmt.Println("Generate temporary session successfully, and it can be used for 3600 seconds..")
	fmt.Println(result)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Prints out full error message, including original error if there was one.
			log.Fatal("Error:", awsErr.Error())
		} else {
			log.Fatal("Error:", err.Error())
		}
	}

	// Start new session using temporary, MFA-authenticated credentials
	sess := session.New(&aws.Config{
		Region: aws.String(endpoints.EuWest1RegionID),
		Credentials: credentials.NewStaticCredentials(
			*result.Credentials.AccessKeyId,
			*result.Credentials.SecretAccessKey,
			*result.Credentials.SessionToken),
	})

	return sess
}
