package main


import (
   "bufio"
   "fmt"
   "os"
   "github.com/SwornStar/minyr/yr"
)


func main() {
   var input string
   scanner := bufio.NewScanner(os.Stdin)
   fmt.Println("Venligst velg convert, average eller exit:")


   for scanner.Scan() {
       input = scanner.Text()
       if input == "q" || input == "exit" {
           fmt.Println("Avslutter programmet.")
           os.Exit(0)
       } else if input == "convert" {
           fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
           // funksjon som åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil
           yr.ConvertTemperatures()
           // flere else-if setninger
       } else if input == "average" {
           fmt.Println("Hva ønsker du å konvertere til? (C/F)")
           scanner.Scan()
           unit := scanner.Text()
           if unit == "C" || unit == "c" {
               yr.AverageTemperature("C")
           } else if unit == "F" || unit == "f" {
               yr.AverageTemperature("F")
           } else {
               fmt.Println("Ugyldig input.")
           }


       }
   }
}

