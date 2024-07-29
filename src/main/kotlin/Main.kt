import kotlin.math.log10
import kotlin.math.pow

fun main(args: Array<String>) {

    println(isPalindrome(12321)) // true
    println(isPalindrome(12345)) // false
    println(isPalindrome(1232)) // false
    println(isPalindrome(1221)) // true
    println(isPalindrome(-121)) // false
    println(isPalindrome(0)) // true
    println(isPalindrome(123456)) // false
}

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

/**
 * !!! Follow up: Could you solve it without converting the integer to a string?
 * ```
 *  return with(x.toString()) { return this == this.reversed() }
 * ```
 */

fun isPalindrome(x: Int): Boolean {
    // signed always false
    if (x<0) return false

    var palindrom = x
    var numSize = (log10(x.toDouble()) + 1).toInt()
    repeat(numSize) {
        val right = palindrom % 10
        val left = palindrom / 10.0.pow(numSize - 1).toInt()
        if (left != right) return false
        palindrom = palindrom % (10.0.pow(numSize-1).toInt()) / 10
        numSize -= 2
        if (numSize <=0) return true
    }
    return true
}