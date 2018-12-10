package twilio

import (
	"github.com/sfreiberg/gotwilio"
)

func SendTextWithMessage(msg string) {

	//authenticate
	from := "+12148172417"
	to := "+12148429453"
	//trial acct
	accountSID, authToken := "AC1f7ba4c766e2fb463f4fac0d85d9ec87", "5d53fbced274c991841d1a34a71d519d"
	twilio := gotwilio.NewTwilioClient(accountSID, authToken)

	//send msg
	twilio.SendSMS(from, to, msg, "", "")
}
