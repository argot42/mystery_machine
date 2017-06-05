package main

import (
    "fmt"
    "os"
    "io"
    //"math"
    "log"
    "image"
    "image/png"
    "image/draw"
    "./bitoperations"
)

func main() {
    if (len(os.Args) < 2) {
        log.Fatal("Too few arguments")
    }

    // open image
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

    b := im.Bounds()
    msg := getmsg(image_max_capacity(b.Max.X, b.Max.Y))

    // create new image with msg encoded into it
    new_image := encode(msg, im)
    //if (err != nil) {
    //    log.Fatal(err)
    //}

    // create new file
    var new_file *os.File
    if (len(os.Args) < 3) {
        new_file, err = os.Create(os.Args[1] + ".encoded")
    } else {
        new_file, err = os.Create(os.Args[2])
    }

    if (err != nil) {
        log.Fatal(err)
    }
    defer new_file.Close()

    fmt.Println(new_image)
    // save image
    png.Encode(new_file, new_image, nil)
}

func getmsg (max int) []byte {
    max_chars := bytes_to_char(max, 7)
    msg := make([]byte, max_chars)

    var in string
    _, e := fmt.Scanf("%s", &in)

    // if something happens when reading stdin or there's not enough space to store characters just return an empty string
    if (e != nil || max_chars-1 < 1) {
        return make([]byte, 0)
    }

    var actual_chars int = len(in)
    var i int = 0

    // get words from stdin until EOF or reaching max characters
    for e != io.EOF && actual_chars <= max_chars-1 {
        for j,c := range in {
            msg[j+(actual_chars*i)] = byte(c)
        }

        _, e = fmt.Scanf("%s", &in)
        i++
        in = " " + in
        actual_chars+=len(in)
    }

    msg[max_chars-1] = '\000'
    return msg
}

func image_max_capacity (width int, height int) int {
    // return total bytes in pixels RGBA
    // 4 bytes for every pixel
    return width * height * 4
}

func encode (msg []byte, img image.Image) (*image.RGBA) {
    rgba := image.NewRGBA(img.Bounds())
    draw.Draw(rgba, rgba.Bounds(), img, image.Point{0,0}, draw.Src)

    for i,n := range msg {

        //fmt.Println("char: ", n, "\n-------")
        for j:=(7*i); j<7*(i+1); j++ {
            // change last bit from color rgba value
            // matching byte from every part of msg letter
            //fmt.Printf("%d", bitoperations.Getbit(int(n), uint(j-(7*i))))
            //fmt.Println(uint8(bitoperations.Changebit(int(rgba.Pix[j]), 0, bitoperations.Getbit(int(n), uint(j-(7*i))))))
            rgba.Pix[j] = uint8(bitoperations.Changebit(int(rgba.Pix[j]), 0, bitoperations.Getbit(int(n), uint(j-(7*i)))))
        }
        //fmt.Printf("\n\n")
    }

    //fmt.Println(rgba)
    return rgba
}

func bytes_to_char (bytes int, char_bit int) int {
    return bytes / char_bit
}
