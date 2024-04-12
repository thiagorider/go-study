package main

import (
    "fmt"
    "strings"
)

func main() {
    var temperatura float64
    var escala string

    fmt.Print("Insira a temperatura: ")
    fmt.Scan(&temperatura)

    fmt.Print("Insira a escala  (C, F ou K): ")
    fmt.Scan(&escala)

    escala = strings.ToUpper(escala)

    switch escala {
    case "C":
        fahrenheit := (temperatura * 9/5) + 32
        kelvin := temperatura + 273.15
        fmt.Printf("Fahrenheit: %.2f\n", fahrenheit)
        fmt.Printf("Kelvin: %.2f\n", kelvin)
    case "F":
        celsius := (temperatura - 32) * 5/9
        kelvin := (temperatura + 459.67) * 5/9
        fmt.Printf("Celsius: %.2f\n", celsius)
        fmt.Printf("Kelvin: %.2f\n", kelvin)
    case "K":
        celsius := temperatura - 273.15
        fahrenheit := (temperatura * 9/5) - 459.67
        fmt.Printf("Celsius: %.2f\n", celsius)
        fmt.Printf("Fahrenheit: %.2f\n", fahrenheit)
    default:
        fmt.Println("Escala inválida.")
    }
}
