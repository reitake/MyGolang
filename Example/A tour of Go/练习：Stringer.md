通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。  

例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。  

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ipa IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ipa[0], ipa[1], ipa[2], ipa[3])
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}

```

输出：
```
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```
