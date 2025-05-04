// https://leetcode.com/problems/length-of-last-word

/** Run: scala-cli run LengthOfLastWord.scala
  */
object LengthOfLastWord:
  @main def main(): Unit =
    val s = "Hello World"
    assert(lengthOfLastWord(s) == 5)
    assert(lengthOfLastWord2(s) == 5)
    assert(lengthOfLastWord3(s) == 5)
    assert(lengthOfLastWord4(s) == 5)

  // scan from the beginning
  def lengthOfLastWord(s: String): Int = {
    var i = 0
    var count = 0
    val n = s.length
    while (i < n) {
      if (s(i) != ' ') {
        count += 1
        i += 1
      } else {
        while (i < n && s(i) == ' ') {
          i += 1
        }

        if (i == n) {
          return count
        } else {
          count = 0
        }
      }
    }
    count;
  }

  // scan from the end
  def lengthOfLastWord2(s: String): Int = {
    val n = s.length();
    var i = n - 1;
    var count = 0;
    while (i >= 0 && s(i) == ' ') {
      i -= 1;
    }

    while (i >= 0 && s(i) != ' ') {
      count += 1;
      i -= 1;
    }

    count;
  }

  def lengthOfLastWord3(s: String): Int = {
    s.trim.split(" ").last.length
  }

  def lengthOfLastWord4(s: String): Int = {
    s.trim.split(" ").reverse.head.length
  }
