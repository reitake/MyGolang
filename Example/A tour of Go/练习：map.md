实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。  

你会发现 strings.Fields 很有帮助。  

```go
package main

import (
    "golang.org/x/tour/wc"

    "strings"
)

func WordCount(s string) map[string]int {
    res := make(map[string]int, 0)
    strSlice := strings.Fields(s)
    for _, str := range strSlice {
        if _, ok := res[str]; ok {
            res[str] += 1
        } else {
            res[str] = 1
        }
    }

    return res
}

func main() {
    wc.Test(WordCount)
}
```

输出：  

```go
PASS
 f("I am learning Go!") = 
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") = 
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") = 
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") = 
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
```

ps：实际操作中应该考虑去掉标点符号？判断大小写？  
