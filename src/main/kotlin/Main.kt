import java.lang.NullPointerException

fun main(args: Array<String>) {
    println(isValid("()"))
    println(isValid("()[]{}"))
    println(isValid("(]"))
    println(isValid("(((())))"))
    println(isValid("()(()){}{{[}]}"))
    println(isValid("(("))
    println(isValid("){"))
}

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

val skobkaMap = mapOf(')' to '(', ']' to '[', '}' to '{')
val skobkaOpenSet = setOf('(','[','{')
var skobkaStack = ArrayDeque<Char>()
fun isValid(s: String): Boolean {
    if (s.length == 1) return false
    s.forEach { char ->
        if (char in skobkaOpenSet) {
            skobkaStack.addLast(char)
        } else {
            val reversedSkobka = skobkaMap[char] ?: throw NullPointerException("Task contains only skobka")
            if (reversedSkobka == skobkaStack.lastOrNull()) {
                skobkaStack.removeLast()
            } else {
                return false
            }
        }
    }
    return skobkaStack.isEmpty()
}