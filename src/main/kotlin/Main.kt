fun main(args: Array<String>) {
    var out = mergeTwoLists(list1 = ListNode(1).apply { this.next = ListNode(2).apply { this.next = ListNode(3) } },
        list2 = ListNode(1).apply { this.next = ListNode(2).apply { this.next = ListNode(3) } })
    while (out != null) {
        println(out.value)
        out = out.next
    }

    println("CASE 2")
    var out2 = mergeTwoLists(
        list1 = ListNode(1).apply { this.next = ListNode(2).apply { this.next = ListNode(3) } },
        list2 = null
    )
    while (out2 != null) {
        println(out2.value)
        out2 = out2.next
    }
}
// 21-merge-two-sorted-lists
class ListNode(var value: Int) {
    var next: ListNode? = null
}

fun mergeTwoLists(list1: ListNode?, list2: ListNode?): ListNode? {
    var left = list1
    var right = list2
    // extra case
    if (left == null) return right
    if (right == null) return left

    var outNode: ListNode? = null
    var outRoot: ListNode? = null
    while (true) {
        // out condition
        if (left == null && right == null)  break

        if (left == null) {
            outNode?.next = ListNode(right?.value!!)
            outNode = outNode?.next
            right = right.next
            continue
        }
        if (right == null) {
            outNode?.next = ListNode(left.value)
            outNode = outNode?.next
            left = left.next
            continue
        }

        if (left.value!! < right?.value!!) {
            if (outRoot == null) {
                outNode = ListNode(left.value)
                outRoot = outNode
            } else {
                outNode?.next = ListNode(left.value)
                outNode = outNode?.next
            }
            left = left.next
        } else {
            if (outRoot == null) {
                outNode = ListNode(right.value)
                outRoot = outNode
            } else {
                outNode?.next = ListNode(left.value)
                outNode = outNode?.next
            }
            right = right.next
        }
    }
    return outRoot
}