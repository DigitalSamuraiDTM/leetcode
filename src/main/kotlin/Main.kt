fun main(args: Array<String>) {

    println(intToRoman(3749)) // MMMDCCXLIX
    println(intToRoman(58)) // LVIII
    println(intToRoman(1800)) // MDCCC
    println(intToRoman(1901)) // MCMI
    println(intToRoman(3999)) // MMMCMXCIX
    println(intToRoman(1)) // I
    println(intToRoman(404)) // CDIV

}


fun intToRoman(num: Int): String {
    var out = ""
    var input = num

    val thousands = num / 1000
    out += ("M".repeat(thousands))
    input -= 1000*thousands

    if (input >= 900) {
        out += "CM"
        input -= 900
    }

    val fivehundreds = input / 500
    out += "D".repeat(fivehundreds)
    input -= 500*fivehundreds

    if (input >= 400) {
        out += "CD"
        input -= 400
    }

    val hundreds = input / 100
    out += "C".repeat(hundreds)
    input -= 100*hundreds

    if (input >= 90) {
        out += "XC"
        input -= 90
    }

    val fifties = input / 50
    out += "L".repeat(fifties)
    input -= 50* fifties

    if (input >= 40) {
        out += "XL"
        input -= 40
    }

    val tens = input / 10
    out += "X".repeat(tens)
    input -= 10*tens

    if (input >= 9) {
        out += "IX"
        input -= 9
    }

    val fives = input / 5
    out += "V".repeat(fives)
    input -= 5*fives

    if (input >= 4) {
        out += "IV"
        input -= 4
    }
    val ones = input / 1
    out += "I".repeat(ones)
    input -= ones
    return out
}