import (
    "fmt"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func longestPalindrome(s string) string {
    palindrom := ""
    l, r, n := 0, -1, len(s)
    if n > 0 {
        palindrom = string(s[0])
    }
    d1, d2 := make([]int, n), make([]int, n)
    
    for i := 0; i < n; i++ {
        var k int
        if i > r {
            k = 1
        } else {
            k = min(d1[l+r-i], r-i+1)
        }
        
        for i+k < n && i-k >= 0 {
            if s[i+k] == s[i-k] {
                currentPalindrom := s[i-k:i+k+1]
                if len(currentPalindrom) >= len(palindrom) {
                    palindrom = currentPalindrom
                }
                k++    
            } else {
                break
            }
        }
        
        d1[i] = k
        if i+k-1 > r {
            l, r = i-k+1, i+k-1
        }
    }
    
    l, r = 0, -1
    for i := 0; i < n; i++ {
        var k int
        if i > r {
            k = 0
        } else {
            k = min(d1[l+r-i+1], r-i+1)
        }
        
        for i+k < n && i-k-1 >= 0 {
            if s[i+k] == s[i-k-1] {
                currentPalindrom := s[i-k-1:i+k+1]
                if len(palindrom) < len(currentPalindrom) {
                    palindrom = currentPalindrom
                }
                k++    
            } else {
                break
            }
        }
        
        d2[i] = k
        if i+k-1 > r {
            l, r = i-k, i+k-1
        }
    }
    
    return palindrom
}
