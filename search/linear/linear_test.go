package linear

import "testing"

func TestLinearSearch(t *testing.T) {
    haystack := [10]int{9,0,8,1,7,2,6,3,5,4}

    t.Run("Should return false if there's no needle", func (t *testing.T) {
        hasNeedle := LinearSearch(haystack[:], 69)

        if hasNeedle {
            t.Error("it wasn't expected to find the needle")
        }
    })

    t.Run("Should return true if there's a needle in the haystack", func (t *testing.T) {
        hasNeedle := LinearSearch(haystack[:], 2)

        if !hasNeedle {
            t.Error("it was expected to find the needle")
        }
    })
}
