package main

import (
	"fmt"
	"math/rand"
	"time"
)

func l() []string {
	return []string{
		"怒りは無謀をもって始まり、後悔をもって終わる。",
		"怒を敵と思へ。",
		"憤怒の火の燃えるのを弱めよ。",
		"憤怒は他人にとって有害であるが、憤怒に駆られている当人にはもっと有害である。",
		"人間は、理性のうちに負けたものの埋め合わせを怒りの中でするものである。",
		"怒りと愚行は相並んで歩み、悔恨が両者のかかとを踏む。",
		"怒りは一時の狂気なり。汝が怒りを制さざれば、怒りが汝を制せん。",
	}
}

func g() string {
	rand.Seed(time.Now().UnixNano())
	l := l()
	r := rand.Intn(len(l))
	return l[r]
}

func main() {
	fmt.Println(g())
}
