
// https://leetcode.com/problems/reverse-string
import scala.annotation.tailrec

/** Run: scala-cli run ReverseString.scala
  */
object ReverseString:
  @main def main(): Unit =
    val s = Array('h', 'e', 'l', 'l', 'o')
    reverseString(s)
    assert(s == Array('o', 'l', 'l', 'e', 'h'))

  // use two pointers: O(n) time and O(1) space
  def reverseString(s: Array[Char]): Unit = {
    var i = 0
    var n = s.length - 1
    while (i < n) {
      val temp = s(i)
      s(i) = s(n)
      s(n) = temp
      i += 1
      n -= 1
    }
  }

  // use one pointer: O(n) time and O(1) space
  def reverseString2(s: Array[Char]): Unit = {
    var i = 0
    while (i < s.length / 2) {
      val temp = s(i)
      s(i) = s(s.length - i - 1)
      s(s.length - i - 1) = temp
      i += 1
    }
  }

  // use Scala for loop: O(n) time and O(1) space
  def reverseStringUsingForLoop(s: Array[Char]): Unit = {
    for (i <- 0 until s.length / 2) {
      val temp = s(i)
      s(i) = s(s.length - i - 1)
      s(s.length - i - 1) = temp
    }
  }

  // use tail recursion: O(n) time and O(1) space (best submission result)
  def reverseStringUsingRecursion(s: Array[Char]): Unit = {
    @tailrec
    def helper(left: Int, right: Int): Unit = {
      if (left < right) {
        val temp = s(left)
        s(left) = s(right)
        s(right) = temp
        helper(left + 1, right - 1)
      }
    }

    helper(0, s.length - 1)
  }

  // (N/A) use stack: O(n) time and space
  def reverseStringUsingStack(s: Array[Char]): Unit = {
    import scala.collection.mutable.Stack
    val stack = Stack[Char]()
    for (c <- s) {
      stack.push(c)
    }

    for (i <- s.indices) {
      s(i) = stack.pop()
    }
  }


/** Summary:
  *   1. the recursion solution is the best submission result 2. the stack
  *      solution is easy to understand but not efficient 3. Scala has a
  *      built-in method: s.reverse which returns a new reversed string
  */
