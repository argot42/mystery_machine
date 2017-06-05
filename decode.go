package main

import (
    "fmt"
    "os"
    "log"
    "math"
    "image"
    "image/draw"
    "image/png"
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
    im, err := png.Decode(file)
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
    rgba := image.NewRGBA(im.Bounds())
    draw.Draw(rgba, rgba.Bounds(), im, image.Point{0,0}, draw.Src)

    var character int = -1
    var buffer string
    var channel int = 0

    for character != int(last_character) {
        character = 0

        for i:=0; i<bits; i++ {

            if ((channel+1) % 4 == 0) { channel++ }

            character += bitoperations.Getbit(int(rgba.Pix[channel]), 0) * int(math.Pow(2, float64(i)))

            channel++
        }

        buffer += string(character)
    }

    return buffer
}
