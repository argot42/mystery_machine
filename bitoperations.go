package bitoperations

func Changebit (n int, pos uint, b int) int {
    return (n & ^(1 << pos)) | (b << pos)
}

func Getbit (n int, pos uint) int {
    return (n >> pos) & 1
}
