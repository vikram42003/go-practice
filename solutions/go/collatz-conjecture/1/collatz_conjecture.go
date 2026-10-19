package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("")
    }
    
	res := 0
    for n != 1 {
        if (n & 1) == 0 {
            n = n / 2
        } else {
            n = (n * 3) + 1
        }
        res++
    }
    
    return res, nil
}
