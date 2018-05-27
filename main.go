package main

import (
	"net/http"
	"github.com/joomtriggers/ideamart"
	"io/ioutil"
	"github.com/kr/pretty"
)

func main() {
	sms := ideamart.SMS()
	smss := sms.Sender
	smss.SetApplication("APP_041418")
	smss.SetPassword("23ff355d399a1567e64297e502912fe7")
	smss.SetServer("https://api.dialog.lk/sms/send/")
	smss.SetSourceAddress("GOLANGTEST")
	http.HandleFunc("/receive/", func(w http.ResponseWriter, r *http.Request) {
		defer smss.Send()
		rp, _ := ioutil.ReadAll(r.Body)
		smsReceiver := sms.Receiver
		smsReceiver.Receive(rp)
		smss.AddReceiver(smsReceiver.Sender)
		smss.SetMessage(smsReceiver.Message)
		pretty.Print(smsReceiver)
	})
	http.ListenAndServe(":8084", nil)
}
