package main
import (
  "fmt"
  "time"
)

func joinTwoStrings(prefix string, suffix string) string {
  return prefix + suffix
}

func printOutDate() {
  t := time.Now()
  fmt.Println(t)
}

func waitForIt(message string) {
  defer fmt.Println("Done!")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println(message)
}

func main() {
  printOutDate()
  waitForIt(joinTwoStrings("Hi", " there"))

  side := 6.98
  areaOnly := getArea(side)
  fmt.Printf("Area of Square with side %f is %.2f\n", side, areaOnly)

  perimeterOnly := getPerimeter(side)
  fmt.Printf("Area of Square with side %f is %.2f\n", side, perimeterOnly)

  side = 71.89
  area, perimeter := getAreaNPerimeter(side)
  fmt.Printf("Area and Perimeter of Square with side %f is %.2f and %.2f respectively\n", side, area, perimeter)

  var enteredSide float64
  var enteredType string
  fmt.Println("How long is the side of the square")
  fmt.Scanln(&enteredSide)

  fmt.Println("What do you want to calculate for the square? area or perimeter?")
  fmt.Scanln(&enteredType)
  fmt.Printf("The %s of the square with side %f is %.2f\n", enteredType, enteredSide, getX(enteredSide, enteredType) )
}

func getArea(side float64) float64 {
  return side*side
}

func getPerimeter(side float64) float64{
  return 4*side
}

func getAreaNPerimeter(side float64) (float64, float64){
  return getArea(side), getPerimeter(side)
}

func getX(side float64, geometryType string) float64{
  defer fmt.Printf("returned %s of square of side %f\n", geometryType, side)

  if geometryType == "area"{
    return getArea(side)
  }else if geometryType == "perimeter"{
    return getPerimeter(side)
  }else{
    fmt.Println("Invalid type")
    return 0
  }
}
