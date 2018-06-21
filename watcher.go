package main

import (
	"fmt"
	//"log"
	"time"
	//"reflect"
	//"cloud.google.com/go/firestore"
	"github.com/lightningnetwork/lnd/lnrpc"
	"golang.org/x/net/context"
)

type Message struct {
	Invoice string `json:"invoice,omitempty"`
	Settled bool   `json:"settled,omitempty"`
}

func watchPayments() {
	//TODO: A better way is to watch for payments and then
	// update firebase.
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				checkPayments()
			}
		}
	}()
}



func checkPayments() {
	c, clean := getClient()
	defer clean()

	// 1st get unsettled message payment hashes
	//it := firebaseDb.Collection("messages").Where("settled", "==", false).Documents(context.Background())
	//snapshot, err := it.GetAll()
	//if err != nil {
	//	log.Fatalln("Failed to get documents ", err)
	//	return
	//}
	//for _, s := range snapshot {
		//invoice := s.Data()["invoice"].(string)
	invoice := "lntb69690n1pdjgw2wpp5xakfaufhth55srq03mfnrj800q0dkx7aaxeuvhky4s236405676qdqagysxxmmxvejk2grxdaezq5n0vajhycqzys0set6gt5lm3p4lnz5d3dv8cx4cqayzjzupe0r55g9h4nrmd6jt2hvhccnu42pg9wlla3zqgtx0fv7ewlw6tlmxvjz4cut3k0x4ykc4cp03f6vn"
	
	decoded, err := c.DecodePayReq(context.Background(), &lnrpc.PayReqString{PayReq: invoice})
	//res, err := c.GetInfo(context.Background(), &lnrpc.GetInfoRequest{})
	//fmt.Println(res)
	//fmt.Println(decoded.GetPaymentHash())
	//fmt.Println(reflect.TypeOf(decoded))
	if err != nil {
		fmt.Println("Failed to decode payreq")
		//continue
	}
	
	lnInvoice, err := c.LookupInvoice(context.Background(), &lnrpc.PaymentHash{RHashStr: decoded.GetPaymentHash()})
	if err != nil {
		// It's possible that invoice generated with a test lnd won't appear in prod lnd.
		// Best approach is to separate them in the DB, but for now, just ignore them.
		fmt.Println("Failed to find invoice ", err)
	}

	//fmt.Println("--",lnInvoice.GetSettled())
	//fmt.Println(lnInvoice)
	//fmt.Println(lnInvoice.GetValue())
	//fmt.Println(lnInvoice.GetPaymentRequest())
	//fmt.Println()
	
	
	t := time.Unix(lnInvoice.GetCreationDate(), 0)
	fmt.Println(t)
	/*if err != nil {
		// It's possible that invoice generated with a test lnd won't appear in prod lnd.
		// Best approach is to separate them in the DB, but for now, just ignore them.
		fmt.Println("Failed to find invoice ", err)
	} else {
		if lnInvoice.GetSettled() {
			_, err := s.Ref.Update(context.Background(), []firestore.Update{{Path: "settled", Value: true}})
			if err != nil {
				log.Println("Update failed ", err)
			} else {
				log.Println("Updated ", invoice)
			}
		}
	}
	*/
	
	//}
}