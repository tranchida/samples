package main

import (
	"encoding/json"
	"github.com/fatih/color"

	bolt "go.etcd.io/bbolt"
)

type message struct {
	MessageId             string            `json:"messageId"`
	BusinessId            string            `json:"businessId"`
	BusinessCorrelationId string            `json:"businessCorrelationId"`
	Headers               map[string]string `json:"headers"`
	ServiceDestination    string            `json:"serviceDestination"`
	Body                  string            `json:"body"`
}


func main() {

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			panic(err)
		}

		m := message{
			MessageId:             "1",
			BusinessId:            "1",
			BusinessCorrelationId: "1",
			Headers: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			ServiceDestination: "service1",
			Body:               "<test/>",
		}

		j, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		b.Put([]byte("2"), j)

		return nil
	})

	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("MyBucket"))
		if b == nil {
			return nil
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			
			color.Red("key -> %s\n", string(k))
			color.Blue("value -> %s\n", string(v))
		}

		return nil
	
	})

}
