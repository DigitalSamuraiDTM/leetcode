fun main(args: Array<String>) {
    println(convert("PAYPALISHIRING",1)) // PINALSIGYAHRPI
    println(convert("A",1)) // A
    println(convert("PAYPALISHIRING",3)) // PAHNAPLSIIGYIR
    println(convert("PAYPALISHIRING",4)) // PINALSIGYAHRPI
    println(convert("PAYPALISHIRING",14)) // PINALSIGYAHRPI
    println(convert("PAYPALISHIRING",13)) // PINALSIGYAHRPI
    println(convert("PAYPALISHIRING",500)) // PINALSIGYAHRPI
}

fun convert(s: String, numRows: Int): String {
    if (numRows == 1) return s
    var zigzagCounter = 0
    var metronome = 1
    val outArray = Array(numRows) { "" }
    s.forEach { char ->
        outArray[zigzagCounter] = outArray[zigzagCounter] + char
        zigzagCounter += metronome
        if (zigzagCounter == numRows-1) {
            metronome = -1
        } else if (zigzagCounter == 0) {
            metronome = 1
        }
    }
    return outArray.joinToString("")
}
