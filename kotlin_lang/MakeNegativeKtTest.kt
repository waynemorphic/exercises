package kotlin_lang

import org.junit.Test
import org.junit.Assert.*;

class MakeNegativeKtTest {

    @Test
    fun makeNegativeTest() {
        assertEquals(-1, MakeNegative().makeNegative(-1))
        assertEquals(-1, MakeNegative().makeNegative(1))
    }
}