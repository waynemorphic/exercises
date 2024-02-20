package kotlin_lang

// Given a number, make it negative. However, if it is already negative, return the number
class MakeNegative {
    fun makeNegative(x: Int): Int {
        return if (x < 0) x else x * -1
    }

}
