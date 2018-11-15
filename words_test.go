package fastfaker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Word() {
	Global.Seed(11)
	fmt.Println(Global.Word())
	// Output: quia
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Word()
	}
}

func ExampleFaker_Sentence() {
	Global.Seed(11)
	fmt.Println(Global.Sentence(5))
	// Output: Quia quae repellat consequatur quidem.
}

func BenchmarkSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Sentence(10)
	}
}

func ExampleFaker_Paragraph() {
	Global.Seed(11)
	fmt.Println(Global.Paragraph(3, 5, 12, "\n"))
	// Output: Quia quae repellat consequatur quidem nisi quo qui voluptatum accusantium quisquam amet. Quas et ut non dolorem ipsam aut enim assumenda mollitia harum ut. Dicta similique veniam nulla voluptas at excepturi non ad maxime at non. Eaque hic repellat praesentium voluptatem qui consequuntur dolor iusto autem velit aut. Fugit tempore exercitationem harum consequatur voluptatum modi minima aut eaque et et.
	// Aut ea voluptatem dignissimos expedita odit tempore quod aut beatae ipsam iste. Minus voluptatibus dolorem maiores eius sed nihil vel enim odio voluptatem accusamus. Natus quibusdam temporibus tenetur cumque sint necessitatibus dolorem ex ducimus iusto ex. Voluptatem neque dicta explicabo officiis et ducimus sit ut ut praesentium pariatur. Illum molestias nisi at dolore ut voluptatem accusantium et fugiat et ut.
	// Explicabo incidunt reprehenderit non quia dignissimos recusandae vitae soluta quia et quia. Aut veniam voluptas consequatur placeat sapiente non eveniet voluptatibus magni velit eum. Nobis vel repellendus sed est qui autem laudantium quidem quam ullam consequatur. Aut iusto ut commodi similique quae voluptatem atque qui fugiat eum aut. Quis distinctio consequatur voluptatem vel aliquid aut laborum facere officiis iure tempora.
}

func BenchmarkParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Paragraph(3, 5, 12, "\n")
	}
}

func TestSentence(t *testing.T) {
	for _, count := range []int{-100, -1, 0} {
		if Global.Sentence(count) != "" {
			t.Errorf("result should be blank for %d words", count)
		}
	}
}
func TestParagraph(t *testing.T) {
	for _, count := range []struct{ parag, sent, words int }{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, -100},
		{1, -100, 1},
		{-100, 1, 1},
		{0, 0, 0},
	} {
		if Global.Paragraph(count.parag, count.sent, count.words, " ") != "" {
			t.Errorf("result should be blank for %v input", count)
		}
	}
}
