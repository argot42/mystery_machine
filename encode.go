package main

import (
    "fmt"
    "os"
    "io"
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
    msg := getmsg(image_max_capacity(b.Max.X, b.Max.Y), '\000')
    //fmt.Println(msg)

    // create new image with msg encoded into it
    new_image := encode(msg, 7, im)
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

    // save image
    png.Encode(new_file, new_image)
}

func getmsg (max int, last_char byte) []byte {
    max_chars := bytes_to_char(max, 7)

    var in string
    _, e := fmt.Scanf("%s", &in)

    // if something happens when reading stdin or there's not enough space to store characters just return an empty string
    if (e != nil || max_chars-1 < 1) {
        return make([]byte, 0)
    }

    var actual_chars int = len(in)
    var msg []byte

    // get words from stdin until EOF or reaching max characters
    for e != io.EOF && actual_chars <= max_chars-1 {
        fmt.Println(actual_chars)
        for _,c := range in {
            msg = append(msg, byte(c))
        }

        _, e = fmt.Scanf("%s", &in)
        in = " " + in
        actual_chars+=len(in)
    }

    msg = append(msg, last_char)

    return msg
}

func image_max_capacity (width int, height int) int {
    // return total bytes in pixels RGBA
    // 4 bytes for every pixel
    // 3 for the moment because I can't use alpha channel
    return width * height * 3 //4
}

func encode (msg []byte, bits int, img image.Image) (*image.RGBA) {
    bounds := img.Bounds()
    rgba := image.NewRGBA(bounds)
    draw.Draw(rgba, rgba.Bounds(), img, image.Point{0,0}, draw.Src)

    var channel int = 0

    for i:=0; i<len(msg); i++ {
        for b:=0; b<bits; b++ {
            if ((channel+1) % 4 == 0) { channel++ }

            rgba.Pix[channel] = uint8(bitoperations.Changebit(int(rgba.Pix[channel]), 0, bitoperations.Getbit(int(msg[i]), uint(b))))

            channel++
        }
    }


    return rgba
}

func bytes_to_char (bytes int, char_bit int) int {
    return bytes / char_bit
}
