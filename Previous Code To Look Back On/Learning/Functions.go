package main

import(
  "fmt"
  "errors"
  "math"
)

func main(){
  result := sum(5, 3)
  fmt.Println("Result of calculation was:", result)

  result2, error := sqrt(16)
  if error!= nil{
    fmt.Println(error)
  }else{
    fmt.Println(result2)
  }

}

// Creating a function expecting two integer values and returning an integer
func sum(x int, y int) int{
  return x + y
}

func sqrt(x float64) (float64, error){
  if x < 0{
    return 0, errors.New("Undefined for negative numbers!")
  }
  return math.Sqrt(x), nil
}
