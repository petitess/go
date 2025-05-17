## go
```go
go mod init site.org/abc
go mod tidy
go run .
```
## time
```go
fmt.Println(time.Now().Format("2006-01-02 15:04:05"))                     //2025-05-17 09:40:16
fmt.Println(time.Now().Format("2006-01-02-15-04-05"))                     //2025-05-17-09-40-16
fmt.Println(time.Now().Format("15-04-05"))                                //09-40-16
fmt.Println(time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")) //2025-05-18 09:40:16
fmt.Println(time.Now().Clock())                                           //9 40 16
fmt.Println(time.Now().Day())                                             //17
fmt.Println(time.Now().Weekday())                                         //Saturday
fmt.Println(time.Now().Month())                                           //May
fmt.Println(time.Now().Format("2006-January-02"))                         //2025-May-17
fmt.Println(time.Now().Compare(time.Now().Add(time.Hour * 24)))           //-1
fmt.Println(time.Now().UnixMicro())                                       //1684314016000000
fmt.Println(time.Now().Round(time.Hour).Format("15:04:05"))               //10:00:00
fmt.Println(time.Now().Truncate(time.Hour).Format("15:04:05"))            //09:00:00
fmt.Println(time.Now().ISOWeek())                                         //2025 20
fmt.Println(time.Now().Zone())                                            //CEST 7200
fmt.Println(time.December.String())                                       //December
fmt.Println(time.NewTimer(time.Hour * 24).C)                              //0xc000186000	
// Get the current time in Chicago
loc, err := time.LoadLocation("America/Chicago")
if err != nil {
  fmt.Println("Error loading location:", err)
  return
}
chicagoTime := time.Now().In(loc)
fmt.Println("Current time in Chicago:", chicagoTime.Format("2006-01-02 15:04:05"))
```
