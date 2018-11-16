package subtext

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bgadrian/fastfaker/randomizer"
)

func ExampleFaker_Word() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	fmt.Println(text.Word())
	// Output: quia
}

func BenchmarkWord(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.Word()
	}
}

func ExampleFaker_Sentence() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	fmt.Println(text.Sentence(5))
	// Output: Quia quae repellat consequatur quidem.
}

func BenchmarkSentence(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.Sentence(10)
	}
}

func ExampleFaker_SentenceAvg() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	fmt.Println(text.SentenceAvg())
	// Output: At illum ut est sit soluta nulla numquam nobis sunt quaerat ea dolores facere deleniti culpa numquam ut.
}

func ExampleFaker_Paragraph() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(11)))
	fmt.Println(text.Paragraph(3, 5, 12, "\n"))
	// Output: Quia quae repellat consequatur quidem nisi quo qui voluptatum accusantium quisquam amet. Quas et ut non dolorem ipsam aut enim assumenda mollitia harum ut. Dicta similique veniam nulla voluptas at excepturi non ad maxime at non. Eaque hic repellat praesentium voluptatem qui consequuntur dolor iusto autem velit aut. Fugit tempore exercitationem harum consequatur voluptatum modi minima aut eaque et et.
	// Aut ea voluptatem dignissimos expedita odit tempore quod aut beatae ipsam iste. Minus voluptatibus dolorem maiores eius sed nihil vel enim odio voluptatem accusamus. Natus quibusdam temporibus tenetur cumque sint necessitatibus dolorem ex ducimus iusto ex. Voluptatem neque dicta explicabo officiis et ducimus sit ut ut praesentium pariatur. Illum molestias nisi at dolore ut voluptatem accusantium et fugiat et ut.
	// Explicabo incidunt reprehenderit non quia dignissimos recusandae vitae soluta quia et quia. Aut veniam voluptas consequatur placeat sapiente non eveniet voluptatibus magni velit eum. Nobis vel repellendus sed est qui autem laudantium quidem quam ullam consequatur. Aut iusto ut commodi similique quae voluptatem atque qui fugiat eum aut. Quis distinctio consequatur voluptatem vel aliquid aut laborum facere officiis iure tempora.
}

func BenchmarkParagraph(b *testing.B) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for i := 0; i < b.N; i++ {
		text.Paragraph(3, 5, 12, "\n")
	}
}

func ExampleFakerText_ParagraphAvg() {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	fmt.Println(text.ParagraphAvg())
	// Output: At illum ut est sit soluta nulla numquam nobis sunt quaerat ea dolores facere deleniti culpa numquam ut. Distinctio maxime consequatur est qui corporis sunt officia odit et quia odit molestias voluptas porro repellendus magnam ipsa.
	//Corporis eos rem non hic esse optio quisquam hic natus earum molestias iste architecto porro et blanditiis iste. Eum repellendus nostrum qui eius suscipit fugit quia quo et nesciunt quod fuga ut vel pariatur libero sequi.
	//Rerum omnis soluta facilis voluptatem possimus et voluptas eaque possimus harum voluptatibus aperiam voluptatibus qui autem quam veniam. Ea voluptas facilis est autem illum esse amet in id doloribus ab architecto hic at aut aliquam impedit.
	//Id unde et porro est repudiandae omnis beatae iusto pariatur quia sed laboriosam voluptate earum dolores facilis aspernatur. Quam aperiam et nihil explicabo voluptates officia ad porro animi officiis et quam in voluptatem eveniet temporibus fuga.
}

func TestSentence(t *testing.T) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for _, count := range []int{-100, -1, 0} {
		if text.Sentence(count) != "" {
			t.Errorf("result should be blank for %d words", count)
		}
	}
}
func TestParagraph(t *testing.T) {
	text := New(randomizer.NewRandWrapper(false, rand.NewSource(42)))
	for _, count := range []struct{ parag, sent, words int }{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, -100},
		{1, -100, 1},
		{-100, 1, 1},
		{0, 0, 0},
	} {
		if text.Paragraph(count.parag, count.sent, count.words, " ") != "" {
			t.Errorf("result should be blank for %v input", count)
		}
	}
}
