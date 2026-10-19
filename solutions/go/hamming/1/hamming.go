package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("DNA strainds are not of equal length")
    }
    
	total := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            total++
        }
    }
    return total, nil
}
