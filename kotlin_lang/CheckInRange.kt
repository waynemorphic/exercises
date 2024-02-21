package kotlin_lang

import org.junit.Test
import org.junit.Assert.*

//Using ranges implement a function that checks whether the date is in the range between the first
// date and the last date (inclusive).

//You can build a range of any comparable elements. In Kotlin in checks are translated to the
// corresponding contains calls and .. to rangeTo calls:

// closed ended range - i in x..y -> translates to [ i in x..<=y ] or .rangeTo()
// open ended range - i in x..<y -> translates to .rangeUntil()

// to iterate numbers in reverse order, use the function, downTo instead of ..
// i.e. for (i in 10 down to 1) print(i)

// an iteration with an arbitrary step can also be implemented using the function, step
// i.e. for (i in 0..8 step 2) print(i) -> prints 02468

class CheckInRange {
    fun checkInRange(date: MyDate, first: MyDate, last: MyDate): Boolean {
        return date in first..last
    }

    @Test
    fun main(){
        val date = MyDate(2024,2,1)
        val first = MyDate(2023, 12,1)
        val last = MyDate(2024, 20, 2)

        assertTrue(checkInRange(date, first, last))
        assertFalse(checkInRange(first, date, last))
    }
}