package main

import (
  "fmt"
  "encoding/json"
  "github.com/go-redis/redis"
)

type Data struct {
  Message string
}

func main() {
  
    var data Data

    db := redis.NewClient(&redis.Options{
      Addr: "redis:6379",
      Password: "",
      DB: 0,
    })

    data = Data{ Message: "Set By Go." }
    payload, _ := json.Marshal(data)
    db.Set("foo", payload, 0)

    val, _ := db.Get("foo").Result()
    json.Unmarshal([]byte(val), &data)
    fmt.Println("Data:", data)

}
