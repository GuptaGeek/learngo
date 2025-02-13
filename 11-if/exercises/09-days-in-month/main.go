// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main
import(
  "fmt"
  )

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {
  args1:=os.Args[1]
  args2:=os.Args[2]
  length:=len(os.Args)
  if length<3 ||length>3{
    fmt.Println("invalid input")
    return
    }
  
  b:=strings.ToLower(args2)
  switch b {
    case "january":
     fmt.Printf("\"%s\" has 31 days\n",args2)
    case "february":
    if  isleapyear(args1) {
     fmt.Printf("\"%s\" has 29 days\n",args2)
      }else{
       fmt.Printf("\"%s\" has 28 days\n",args2)
      }
    case "march":
     fmt.Printf("\"%s\" has 31 days\n",args2)
    case "april":
     fmt.Printf("\"%s\" has 30 days\n",args2)
    case "may":
     fmt.Printf("\"%s\" has 31 days\n",args2)
    case "june":
     fmt.Printf("\"%s\" has 30 days\n",args2)
    case "july":
    fmt.Printf("\"%s\" has 31 days\n",args2)
    case "august":
    fmt.Printf("\"%s\" has 31 days\n",args2)
    case "september":
    fmt.Printf("\"%s\" has 30 days\n",args2)
    case "october":
    fmt.Printf("\"%s\" has 31 days\n",args2)
    case "november":
    fmt.Printf("\"%s\" has 30 days\n",args2)
    
    case "december":
    fmt.Printf("\"%s\" has 31 days\n",args2)
    default:
    fmt.Println("INAVLID INPUT")
  
}
  func isleapyear(year int)bool{
    if year%4 == 0 {
        if year%100 == 0 {
            if year%400 == 0 {
                return true
            }
            return false
        }
        return true
    }
    return false
  }
    
