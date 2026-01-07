// Package maximumpoints: https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/description/
package maximumpoints

// We can take total of k elements from either left side or right side:
// 1. we take 0 from left, then k from right
// 2. we take 1 from left, then k-1 from right
// ...
// n. we take k-1 from left, then take 1 from right
// Time complexity: O(n), 13ms, only beats 7%
// Space complexity: O(n)
func maxScore(cardPoints []int, k int) int {
	start, maxScore := 0, 0

	validPoints := make([]int, 2*k)
	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		validPoints = append(validPoints, cardPoints[i])
	}
	for i := range k {
		validPoints = append(validPoints, cardPoints[i])
	}

	currentScore := 0
	for end := 0; end < len(validPoints); end++ {
		currentScore += validPoints[end]
		if end-start+1 == k {
			maxScore = max(maxScore, currentScore)
			currentScore -= validPoints[start]
			start++
		}
	}

	return maxScore
}

// When we pick `k` cards from either end of the array, we're actually leaving behind `n-k` consecutive cards in the middle that're not picked.
// If we know the minimum total points of the unpicked cards, then we know the maximum total points of the picked cards.
// Time complexity: O(n), 0ms, beats 100%
// Space complexity: O(1)
func maxScore2(cardPoints []int, k int) int {
	total := 0
	for _, p := range cardPoints {
		total += p
	}

	if k >= len(cardPoints) {
		return total
	}

	start, currentTotal, maxPoints := 0, 0, 0
	for end := range len(cardPoints) {
		currentTotal += cardPoints[end]
		if end-start+1 == len(cardPoints)-k {
			maxPoints = max(maxPoints, total-currentTotal)

			currentTotal -= cardPoints[start]
			start++
		}
	}

	return maxPoints
}
