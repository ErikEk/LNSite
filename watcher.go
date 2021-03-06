package main

import (
	"fmt"
	//"log"
	//"time"
	"reflect"
	//"cloud.google.com/go/firestore"
	"github.com/lightningnetwork/lnd/lnrpc"
	"golang.org/x/net/context"
)

/*
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
*/

func checkPayments(invoice string) (payed bool, value int64) {
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

	decoded, err := c.DecodePayReq(context.Background(), &lnrpc.PayReqString{PayReq: invoice})
	if err != nil {
		fmt.Print(err)
	}
	res, err := c.GetInfo(context.Background(), &lnrpc.GetInfoRequest{})
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(res)
	fmt.Println(decoded.GetPaymentHash())
	fmt.Println(reflect.TypeOf(decoded))
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

	return lnInvoice.GetSettled(), lnInvoice.GetValue()
}
