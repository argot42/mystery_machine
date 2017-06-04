package bitoperations

func Changebit (n int, pos uint, b int) int {
    return (n & ^(1 << pos)) | (b << pos)
}

func Getbit (n int, pos int) int {
    return (n >> 0) & 1
}
