package main

import (
	"net/http"
	"io/ioutil"
	"github.com/joomtriggers/ideamart"
)

func main() {
	sms := ideamart.SMS()
	configuration := map[string]string{
		"applicationId": "APP_041418",
		"password":      "23ff355d399a1567e64297e502912fe7",
		"server":        "https://api.dialog.lk/sms/send/",
	}
	sms.Configure(configuration)
	sms.Sender.SetSourceAddress("GOLANGTEST")
	http.HandleFunc("/receive/", func(w http.ResponseWriter, r *http.Request) {
		defer sms.Sender.Send()
		rp, _ := ioutil.ReadAll(r.Body)
		sms.Receiver.Receive(rp)
		sms.Sender.AddReceiver(sms.Receiver.Sender)
		sms.Sender.SetMessage(sms.Receiver.Message)
	})
	http.ListenAndServe(":8084", nil)
}
