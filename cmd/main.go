package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/xidongc-wish/mgo"
	"github.com/xidongc-wish/mgo/bson"
	"time"
)

type OrderStatus string

type InvoiceNumber int32

type ShippingAddr string

type MileStone string

type OrderFee float64

type Channel int8

type SKU string

type Order struct {

	// customer fields
	Id_ bson.ObjectId `bson:"_id"`
	CustomerId int `bson:"cid"`
	OrderStatus OrderStatus `bson:"status"`
	OrderPlaced time.Time `bson:"placed"`
	InvoiceNumber InvoiceNumber `bson:"invoice"`
	ShippingAddr ShippingAddr `bson:"addr"`
	LastUpdated time.Time `bson:"last"`
	OrderFee OrderFee `bson:"fee"`

	// warehouse field
	TrackingNumber string `bson:"tid"`
	WarehouseId int64 `bson:"wid"`
	MileStone MileStone `bson:"ms"`
	Weight float32 `bson:"weight"`
	Channel Channel `bson:"channel"`
	Skus []SKU `bson:"sku"`
}

func main() {
	dialInfo := &mgo.DialInfo{
		Addrs:					[]string{"10.88.2.211", "10.88.2.83"},
		Direct:                true,
		FailFast:              false,
		LoadBalancingStrategy: mgo.Fastest,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
	c := session.DB("mpc").C("mpc")
	var result Order
	err = c.Find(bson.M{"cid": 30}).One(&result)
	if err == nil {
		log.Info(result)
	} else {
		log.Info(err)
	}
	defer session.Close()
}
