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
    fmt.Println(im)
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

    fmt.Println(rgba)

    var character int = -1
    var buffer string

    for i:=0; character != int(last_character); i++ {
        character = 0

        // strange do while
        j:=i*bits
        for ok:=true; ok; ok=(j%bits != 0) {
            character += bitoperations.Getbit(int(rgba.Pix[j]), 0) * int(math.Pow(2, float64(j-(i*bits))))
            j++
        }

        buffer += string(character)
    }

    return buffer
}
