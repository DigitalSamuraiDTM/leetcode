fun main(args: Array<String>) {

    val e11 = ListNode(2).also { it.next = ListNode(4).also { it.next = ListNode(3) } }
    val e12 = ListNode(5).also { it.next = ListNode(6).also { it.next = ListNode(4) } }
    addTwoNumbers(e11,e12)?.print() ?: println("E1 NULL")
    println("EXAMPLE 2")
    val e21 = ListNode(9).also { it.next = ListNode(9).also { it.next = ListNode(9).also { it.next = ListNode(9).also { it.next = ListNode(9).also { it.next = ListNode(9).also { it.next = ListNode(9) } } } } } }
    val e22 = ListNode(9).also { it.next = ListNode(9).also { it.next = ListNode(9).also { it.next = ListNode(9) } } }
    addTwoNumbers(e21,e22)?.print()
}
fun ListNode.print() {
    println(this.`val`)
    if (this.next != null) this.next?.print()
}

// Example 1:
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.

//Example 2:
//Input: l1 = [0], l2 = [0]
//Output: [0]

//Example 3:
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]

fun addTwoNumbers(l1: ListNode?, l2: ListNode?): ListNode? {
    return recursively(l1,l2,false)
}

fun recursively(l1: ListNode?, l2: ListNode?, extraOne: Boolean): ListNode? {
    var sum = (l1?.`val` ?: 0) + (l2?.`val` ?: 0) + (if (extraOne) 1 else 0)
    val hasNextExtraOne = if (sum >= 10) {
        sum -= 10
        true
    } else {
        false
    }
    val newListNode = ListNode(sum)
    newListNode.next = if (l1?.next == null && l2?.next == null) {
        if (hasNextExtraOne) ListNode(1) else null
    } else {
        recursively(l1?.next, l2?.next, hasNextExtraOne)
    }
    return newListNode
}