package comma

import "strings"

func Comma(s string) string {
    _ = "breakpoint"

    var sign byte
    if s[0] == '+' || s[0] == '-' {
        sign = s[0]
        s = s[1:]
    }

    parts := strings.Split(s, ".")
    result := intComma(parts[0])

    if len(parts) > 1 {
        result += "." + parts[1]
    }
    if sign != 0x0 {
        return string(sign) + result
    }
    return result
}

func intComma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return Comma(s[:n-3]) + "," + s[n-3:]
}
