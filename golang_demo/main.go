package main

import (
    "fmt"
    "github.com/bradfitz/gomemcache/memcache"
    "log"
    "os"
    "errors"
)

func main() {

    if len(os.Args) < 3 {
        log.Fatalln(errors.New("Usage: main key value"))
    }

    key := os.Args[1]
    value := os.Args[2]

    mc := memcache.New("memcached:11211")
    mc.Set(&memcache.Item{Key: key, Value: []byte(value)})

    it, err := mc.Get(key)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(it)
        fmt.Println(string(it.Value[:]))
    }
}
