package main

import (
    "fmt"
    "os"
    "io"
    "errors"
    "math"
    "log"
    "image"
    _ "image/png"
    "draw"
)

func main() {
    if (len(os.Args) < 2) {
        log.Fatal("Too few arguments")
    }

    // open image
    file, err := os.Open(Args[1])
    if (err != nil) {
        log.Fatal(err)
    }

    // decode image
    im, _, err := image.Decode(file)
    if (err != nil) {
        log.Fatal(err)
    }
    defer file.Close()

    msg := getmsg(image_max_capacity(im.Width, im.Height))

    // create new image with msg encoded into it
    new_image, err := encode(msg, im)
    if (err != nil) {
        log.Fatal(err)
    }

    // create new file
    if (len(os.Args) < 3) {
        new_file, err := os.Create(os.Args[1] + ".encoded")
    } else {
        new_file, err := os.Create(os.Args[2])
    }

    if (err != nil) {
        log.Fatal(err)
    }
    defer new_file.Close()

    // save image
    image.Encode(new_file, new_image)
}

func getmsg (max int) string {
    max_chars := bits_to_char(max)
    msg := make(string, max_chars)

    var in string
    _, e := fmt.Scanf("%s", &in)

    // if something happens when reading stdin or there's not enough space to store characters just return an empty string
    if (e != nil || max_chars-1 < 1) {
        return ""
    }

    var actual_chars int = len(in)
    var i int = 1

    for e != io.EOF && actual_chars < max_chars-1 {
        for j,c := range in {
            msg[j+(actual_chars*i)] = c
        }

        fmt.Scanf("%s", &in)
        i++
        in = ' ' + in
        actual_chars+=len(in)
    }

    return msg
}

func image_max_capacity (width int, height int) int {
    // return total bits in pixels RGBA
    // 32 bits for every pixel
    return width * height * 4 * 8
}

func encode (msg string, img image.Image) (image.RGBA, err error) {
    
}
