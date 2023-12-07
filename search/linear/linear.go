package linear

func LinearSearch(haystack []int, needle int) bool {
    for _, num := range haystack {
        if num == needle {
            return true
        }
    }
    return false
}
