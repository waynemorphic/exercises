package kotlin_lang

import org.junit.Test
import org.junit.Assert.*;

//Learn about operator overloading and how the different conventions for operations like ==, <, + work
//  in Kotlin. Add the function compareTo to the class MyDate to make it comparable. After this, the
//  code below date1 < date2 should start to compile.

//Note that when you override a member in Kotlin, the override modifier is mandatory.

data class MyDate(val year: Int, val month: Int, val dayOfMonth: Int) : Comparable<MyDate> {
    // Solution step 1: Overload the compareTo method from Comparable class by following Kotlin convention
    //   that allows Operator overloading
    override fun compareTo(other: MyDate): Int {
        // Solution step 2: Return integer that compares two objects of type MyDate
        return compareValuesBy(this, other, MyDate::year, MyDate::month, MyDate::dayOfMonth)
    }
}

fun test(date1: MyDate, date2: MyDate): Boolean {
    // this code should compile:
    return date1 < date2
}

// Solution 3: Sample test
class Comparisons {
    @Test
    fun main(): Unit {
        val date1 = MyDate(2023, 12, 12)
        val date2 = MyDate(2024, 1, 1)

        assertTrue(test(date1, date2) )
        assertFalse(test(date2, date1))
    }
}


