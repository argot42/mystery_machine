package main

import (
    "fmt"
    "os"
    "log"
    "math"
    "image"
    "image/draw"
    _ "image/png"
    "./bitoperations"
)

func main() {
    if (len(os.Args) < 2) {
        log.Fatal("Too few arguments")
    }

    file, err := os.Open(os.Args[1])
    if (err != nil) {
        log.Fatal(err)
    }

    // decode image
    im, _, err := image.Decode(file)
    if (err != nil) {
        log.Fatal(err)
    }
    defer file.Close()

    msg := decode(im, 7, '\000')
    if (err != nil) {
        log.Fatal(err)
    }

    fmt.Println(msg)
}

func decode (im image.Image, bits int, last_character rune) string {

    // get RGBA from image
    b := im.Bounds()
    rgba := image.NewRGBA(b)
    draw.Draw(rgba, b, im, b.Min, draw.Src)

    fmt.Println(rgba)

    var character int = -1
    var buffer string

    for i:=0; character != int(last_character); i++ {
        character = 0

        for j:=i*bits; j%bits != 0; j++ {
            character += bitoperations.Getbit(int(rgba.Pix[j]), 0) * int(math.Pow(2, float64(j-(i*bits))))
        }

        buffer += string(character)
    }

    return buffer
}
