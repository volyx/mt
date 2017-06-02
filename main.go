package main

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "os"
        "io/ioutil"
        "encoding/json"
)

type Person struct {
        Name string
        Phone string
}

type Config struct {
        Uri string
}

func main() {

        file, _ := os.Open("config.json")

        contents,_ := ioutil.ReadFile("config.json")
        println(string(contents))

        decoder := json.NewDecoder(file)
        config := Config{}
        err := decoder.Decode(&config)
        if err != nil {
                log.Panic(err)
        }

        session, err := mgo.Dial(config.Uri)
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	               &Person{"Cla", "+55 53 8402 8510"})
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
}
