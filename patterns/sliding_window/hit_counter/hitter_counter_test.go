package main

import "testing"

func TestHitCounter(t *testing.T) {
	hitCounter := NewHitCounter()
	hitCounter.hit(1)
	hitCounter.hit(2)
	hitCounter.hit(3)
	hitsAt4 := hitCounter.getHits(4)
	if hitsAt4 != 3 {
		t.Errorf("hitsAt4, got: %d, want: 3", hitsAt4)
	}
	hitCounter.hit(300)
	hitCounter.hit(300)
	hitCounter.hit(300)
	hitsAt300 := hitCounter.getHits(300)
	if hitsAt300 != 6 {
		t.Errorf("hitsAt300, got: %d, want: 6", hitsAt300)
	}
}
