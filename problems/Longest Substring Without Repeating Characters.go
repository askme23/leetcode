import (
    "fmt"
)

func lengthOfLongestSubstring(s string) int {
    maxLen := len(s)
    if maxLen < 2 {
        return maxLen
    }
    
    cntMap := map[byte]int{ s[0]: 1 }
    start := 0
    length := 1
    maxLen = 1
  
    n := len(s)
    for i := 1; i < n; i++ {
        if val, ok := cntMap[s[i]]; ok {
            if val == 0 {
                cntMap[s[i]] = 1
                length++
            } else {
                for {
                    if v, ok := cntMap[s[i]]; ok && v == 1 {
                        cntMap[s[start]] = 0
                        start++
                        length--    
                    } else {
                        break
                    }
                }
                cntMap[s[i]] = 1
                length++
            }           
        } else {
            cntMap[s[i]] = 1
            length++
        }
        
        if length > maxLen {
            maxLen = length
        }
    }
    
    return maxLen
}
