package main

import(
  "fmt"
)

func main(){
  arr := []string{"a","b","c"}

  //Iterating through array
  for index, value := range arr{
    fmt.Println("index:", index, "value:", value)
}
  m := make(map[string]string)
  m["a"] = "alpha"
  m["b"] = "beta"
  for key, value := range m{
    fmt.Println("key:", key, "value:", value)

  }
}
