import kotlin.math.*

fun main(args: Array<String>) {
    println(reverse(900000)) // 9
    println(reverse(-900000)) // -9
    println(reverse(1)) // 1
    println(reverse(10)) // 1
    println(reverse(-9)) // 1
    println(reverse(0)) // 0
    println(reverse(2147483641)) // 1463847412
    println(reverse(2147483641)) // 1463847412
    println(reverse(1463847412)) // 2147483641
    println(reverse(-10101010)) // -1010101
    println(reverse(-123)) // -321
    println(reverse(-1000000000)) //  -1
    println(reverse(-2147483648)) // 0
    println(reverse(123456789)) // 987654321
    println(reverse(2147483647)) // 0
    println(reverse(-1000000003)) // 0
    println(reverse(1000000003)) // 0
    println(reverse(1000000000)) // 1
    println(reverse(123)) // 312
    println(reverse(120)) // 21
}

val minValue = -2147483648
val maxValue = 2147483647
fun reverse(x: Int): Int {
    var num = x
    var index = 10 - ceil(log10(abs(x.toDouble()))).toInt()
    var shouldCheckLimit = index == 0
    var out = 0
    if (num <= 9 && num >= -9) return num
    do {
        val last = num % 10
        if (!shouldCheckLimit) {
            out += abs(last)
            out *= 10
            num /= 10
            continue
        }

        if (x<0) {
            val pivo = (minValue / (10.0.pow(10-index-1)).toInt()) % 10
            if (last < pivo) {
                return 0
            } else {
                out += abs(last)
                out *= 10
                if (last > pivo) shouldCheckLimit = false
            }
        } else {
            val pivo = (maxValue / (10.0.pow(10-index-1)).toInt()) % 10
            if (last > pivo) {
                return 0
            } else {
                out += last
                out *= 10
                if (last < pivo) shouldCheckLimit = false
            }
        }
        num /= 10
        index ++
    } while (num > 9 || num < -9)

    val result = out + abs(num)
    return if (x<0) result * -1 else result
}