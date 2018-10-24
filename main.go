package main

import (
	"flag"
)

/**
 * The command line interface to get the session with different credentials
 * @author Calvin
 */
func main() {

	var awsProfile string
	flag.StringVar(&awsProfile, "profile", "", "Profile credentials used by aws cli")

	var awsRegion string
	flag.StringVar(&awsRegion, "region", "", "Name of the region in AWS")

	var strTokenCode string
	flag.StringVar(&strTokenCode, "tokenCode", "", "MFA token code")

	var strSerialNumber string
	flag.StringVar(&strSerialNumber, "serial", "", "Serial number used in ~/.aws/credentials")

	flag.Parse()

	GetSession(awsProfile, strSerialNumber, strTokenCode)

}
