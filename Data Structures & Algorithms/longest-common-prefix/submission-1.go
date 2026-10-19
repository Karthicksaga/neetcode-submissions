func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }

    sort.Strings(strs)
    first := strs[0]
    last := strs[len(strs)-1]
    n := len(first)
    if len(last) < n {
        n = len(last)
    }

    for i := 0; i < n; i++ {
        if first[i] != last[i] {
            return first[:i]
        }
    }
    return first
}