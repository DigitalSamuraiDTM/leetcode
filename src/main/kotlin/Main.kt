import java.lang.NullPointerException

fun main(args: Array<String>) {
    println(romanToInt("IV"))
    println(romanToInt("LVIII"))
    println(romanToInt("MCMXCIV"))
    println(romanToInt("III"))
}

//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

//  Example 1:
// Input: s = "III"
// Output: 3
// Explanation: III = 3.

// Example 2:
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.

// Example 3:
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
val romanMap = mapOf<Char, Int>(
    'I' to 1,
    'V' to 5,
    'X' to 10,
    'L' to 50,
    'C' to 100,
    'D' to 500,
    'M' to 1000
)
fun romanToInt(s: String): Int {
    var out = 0
    var storage = 0
    var previousValue: Int? = null
    s.forEach { char ->
        val value = romanMap[char] ?: throw NullPointerException("task always has char from map")
        if (previousValue == null) {
            storage = value
        } else {
            // CM -> 100, 1000
            if (previousValue!! < value) {
                out += (value - storage)
                storage = 0
            } else if (previousValue!! > value) {
                // MC -> 1000, 100
                out += storage
                storage = value
            } else {
                // CC -> 100, 100
                storage += value
            }
        }
        previousValue = value
    }
    return out + storage
}