package org.nkcoder.array_hash

/** EASY: https://leetcode.com/problems/valid-anagram/description/
  *
  * An Anagram is a word or phrase formed by rearranging the letters of a
  * different word or phrase, using all the original letters exactly once.
  */
object ValidAnagram:

  // hash: T-O(n), S-O(n)
  // Can also use 2 hash maps and compare the count for each letter.
  def validAnagram(s: String, t: String): Boolean =
    val hash = collection.mutable.Map[Char, Int]()

    for (c <- s)
      hash.put(c, hash.getOrElse(c, 0) + 1)

    for (c <- t)
      hash.put(c, hash.getOrElse(c, 0) - 1)

    hash.forall((_, v) => v == 0)

  // sort: T-O(nlogn), S-O(n) (depends on the sort algorithm)
  def validAnagram2(s: String, t: String): Boolean =
    if s.length != t.length then false
    else s.sorted == t.sorted