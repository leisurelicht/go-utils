package shift

import "strconv"

func Int642StringSlice(s []int64) []string {
    result := make([]string, len(s))
    for i, v := range s {
        result[i] = strconv.FormatInt(v, 10)
    }
    return result
}
