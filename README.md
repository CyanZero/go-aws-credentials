# go-aws-credentials
Use GO AWS SDK connect to AWS and support different kind of credentials, e.g. profile, temporary session token and etc.

## Example
Go get
```
go get github.com/CyanZero/go-aws-credentials/sessionutil
```

In the code to import and create a new session
```
import (
	"github.com/CyanZero/go-aws-credentials/sessionutil"
)

// To get a temporary session via MFA token
func test() {
  sess := sessionutil.GetSession("", "SerialNumber of your MFA device", "Token Code from you MFA device")
}
```

** This is still in beta version **
