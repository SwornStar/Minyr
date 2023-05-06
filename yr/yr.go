package yr


import (
   "bufio"
   "encoding/csv"
   "fmt"
   "log"
   "os"
   "strconv"
   "strings"
   "github.com/SwornStar/funtemps/conv"
)


// Tar filnavn som argument og returnerer true dersom filen eksisterer, false ellers
// Bruker os.Stat for å hente informasjon om filen og os.IsNotExist for å sjekke om det oppstod en feil
func checkFileExists(filename string) bool {
   _, err := os.Stat(filename)
   return !os.IsNotExist(err)
}


func ConvertTemperatures() {
   // Sjekker om output filen eksisterer fra før og spør brukeren om den vil overskrive den
   outputFileName := "kjevik-temp-fahr-20220318-20230318.csv"
   //Sjekker om output filen allerede eksisterer fra før
   if checkFileExists(outputFileName) {
       // Dersom den eksisterer, spør brukeren om den vil overskrive den
       fmt.Print("Output fil eksisterer allerede, vil du overskrive? (j/n): ")
       reader := bufio.NewReader(os.Stdin)
       response, _ := reader.ReadString('\n')
       // Fjerner whitespace fra svaret
       response = strings.TrimSpace(response)
       // Dersom brukeren ikke svarer ja, avsluttes funksjonen så lenge du svarer alt annet enn "y"
       if response != "j" {
           fmt.Println("Ingen overskriving utført")
           return
       }
   }
   // Åpner filen som skal leses
   file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
   if err != nil {
       log.Fatal(err)
       return
   }
   defer file.Close()
   // Lager en CSV reader for å lese filen
   reader := csv.NewReader(file)
   reader.Comma = ';' // Set the field delimiter
   // Leser alle linjene i filen og lagrer dem til minnet
   lines, err := reader.ReadAll()
   if err != nil {
       log.Fatal(err)
       return
   }
   // Lager ny output fil (Overskriver eksisterende fil hvis den eksisterer)
   outputFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
   if err != nil {
       log.Fatal(err)
       return
   }
   defer outputFile.Close()
   // Lager en CSV writer for å skrive til filen
   writer := csv.NewWriter(outputFile)
   writer.Comma = ';'
   // Itererer gjennom alle linjene
   for i, line := range lines {
       // Dersom det er første linje, skriv den til filen uten å gjøre noe med den
       if i == 0 {
           err = writer.Write(line)
           if err != nil {
               log.Fatal(err)
               return
           }
           continue
       }
       if i == len(lines)-1 { // Hvis det er siste linje, legg til en ny linje med tekst
           newLine := []string{"Data er gyldig per 18.03.2023 (CC BY 4.0)", "Meteorologisk institutt (MET)", "endringen er gjort av", "Felix Knutsen"}
           fmt.Println(newLine)
           err = writer.Write(newLine)
           if err != nil {
               log.Fatal(err)
               return
           }
           writer.Flush()
       } else {
           celsiusTemp, _ := strconv.ParseFloat(line[3], 64)
           fahrenheitTemp := conv.CelsiusToFahrenheit(celsiusTemp)
           newLine := []string{line[0], line[1], line[2], fmt.Sprintf("%.2f", fahrenheitTemp)}
           err = writer.Write(newLine)
           if err != nil {
               log.Fatal(err)
               return
           }
           fmt.Println(newLine)
           writer.Flush()
       }
   }
   writer.Flush()
   // Sjekker om det oppstod en feil under flush
   if err := writer.Error(); err != nil {
       log.Fatal(err)
   }
}


func AverageTemperature(unit string) float64 {
   // Åpner filen som skal leses
   inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
   if err != nil {
       log.Fatal(err)
   }
   defer inputFile.Close()


   // Lager en CSV reader for å lese filen
   reader := csv.NewReader(inputFile)
   reader.Comma = ';'
   records, err := reader.ReadAll()
   if err != nil {
       log.Fatal(err)
   }


   var sum float64
   var count int
   for i, record := range records {
       // Gå forbi første linje
       if i == 0 {
           continue
       }


       if record[3] == "" {
           // Hoppe over linjer uten temperatur
           continue
       }
       // Konvertere temperatur fra string til float64
       celsius, err := strconv.ParseFloat(record[3], 64)
       if err != nil {
           log.Fatal(err)
       }


       sum += celsius
       count++
   }
   if unit == "C" {
       fmt.Printf("%.2f\n", sum/float64(count))
       return sum / float64(count)
   }
   if unit == "F" {
       fmt.Printf("%.2f\n", conv.CelsiusToFahrenheit(sum/float64(count)))
       return conv.CelsiusToFahrenheit(sum / float64(count))
   }
   return 0
}