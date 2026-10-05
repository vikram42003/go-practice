package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	total := 0
    for _, x := range cb[file] {
        if x {
            total += 1
        }
    }
    return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    fmt.Println(cb)
	if rank < 1 || rank > 8 {
        return 0
    }
    
	files := []string{ "A", "B", "C", "D", "E", "F", "G", "H" }
    total := 0
    for _, f := range files {
        if cb[f][rank - 1] {
            total += 1
        }
    }
    return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	total := 0
    for _, v := range cb {
        total += len(v)
    }
    return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total := 0
    for _, v := range cb {
        for _, vv := range v {
            if vv {
                total += 1
            }
        }
    }
    return total
}
