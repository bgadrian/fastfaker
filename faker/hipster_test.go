package faker

import (
	"fmt"
	"testing"

	"github.com/bgadrian/fastfaker/data"
)

func ExampleFaker_HipsterWord() {
	Global.Seed(11)
	fmt.Println(Global.HipsterWord())
	// Output: microdosing
}

func BenchmarkHipsterWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.HipsterWord()
	}
}

func ExampleFaker_HipsterSentenceAvg() {
	Global.Seed(42)
	fmt.Println(Global.HipsterSentenceAvg())
	// Output: Gastropub five dollar toast mumblecore kinfolk loko bushwick semiotics pickled tote bag neutra loko tote bag +1 salvia organic umami polaroid synth.
}

func ExampleFaker_HipsterSentence() {
	Global.Seed(11)
	fmt.Println(Global.HipsterSentence(5))
	// Output: Microdosing roof chia echo pickled.
}

func BenchmarkHipsterSentence(b *testing.B) {
	Global.Seed(42)
	for i := 0; i < b.N; i++ {
		Global.HipsterSentence(10)
	}
}

func ExampleFaker_HipsterParagraphAvg() {
	Global.Seed(42)
	fmt.Println(Global.HipsterParagraphAvg())
	// Output: Gastropub five dollar toast mumblecore kinfolk loko bushwick semiotics pickled tote bag neutra loko tote bag +1 salvia organic umami polaroid synth. Wolf roof PBR&B asymmetrical craft beer taxidermy everyday venmo shabby chic art party marfa flexitarian intelligentsia mumblecore lumbersexual occupy roof you probably haven't heard of them.
	//Kitsch wayfarers hashtag pitchfork shabby chic echo cronut gentrify wolf jean shorts pabst normcore XOXO DIY 3 wolf moon plaid roof fanny pack. Thundercats typewriter deep v literally lo-fi cardigan hammock umami keytar polaroid schlitz yr YOLO sriracha small batch neutra truffaut sustainable.
	//Slow-carb fanny pack neutra wayfarers helvetica polaroid vinegar biodiesel chambray street meh cleanse drinking cliche blog pickled kale chips meggings. Selvage tattooed squid Wes Anderson fanny pack hoodie tofu yr sustainable chillwave slow-carb beard gluten-free carry stumptown hoodie actually williamsburg.
	//Fanny pack austin ethical portland church-key narwhal next level cleanse distillery mixtape deep v humblebrag umami gentrify kogi distillery art party Yuccie. Gentrify tumblr Yuccie VHS humblebrag swag franzen direct trade squid venmo post-ironic chillwave direct trade ethical meh vice post-ironic gentrify.
}

func ExampleFaker_HipsterParagraph() {
	Global.Seed(11)
	fmt.Println(Global.HipsterParagraph(3, 5, 12, "\n"))
	// Output: Microdosing roof chia echo pickled meditation cold-pressed raw denim fingerstache normcore sriracha pork belly. Wolf try-hard pop-up blog tilde hashtag health butcher waistcoat paleo portland vinegar. Microdosing sartorial blue bottle slow-carb freegan five dollar toast you probably haven't heard of them asymmetrical chia farm-to-table narwhal banjo. Gluten-free blog authentic literally synth vinyl meh ethical health fixie banh mi Yuccie. Try-hard drinking squid seitan cray VHS echo chillwave hammock kombucha food truck sustainable.
	// Pug bushwick hella tote bag cliche direct trade waistcoat yr waistcoat knausgaard pour-over master. Pitchfork jean shorts franzen flexitarian distillery hella meggings austin knausgaard crucifix wolf heirloom. Crucifix food truck you probably haven't heard of them trust fund fixie gentrify pitchfork stumptown mlkshk umami chambray blue bottle. 3 wolf moon swag +1 biodiesel knausgaard semiotics taxidermy meh artisan hoodie +1 blue bottle. Fashion axe forage mixtape Thundercats pork belly whatever 90's beard selfies chambray cred mlkshk.
	// Shabby chic typewriter VHS readymade lo-fi bitters PBR&B gentrify lomo raw denim freegan put a bird on it. Raw denim cliche dreamcatcher pug fixie park trust fund migas fingerstache sriracha +1 mustache. Tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid stumptown organic schlitz. Flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy. Chia mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade.
}

func BenchmarkHipsterParagraph(b *testing.B) {
	Global.Seed(42)
	for i := 0; i < b.N; i++ {
		Global.HipsterParagraph(3, 5, 12, "\n")
	}
}

func TestFaker_HipsterSetCacheNil(t *testing.T) {
	faker := NewFastFaker()
	old := data.Data["hipster"]["word"]
	data.Data["hipster"]["word"] = nil

	if faker.HipsterSentence(3) != "" {
		t.Error("should return empty result with null randomizer")
	}

	if faker.HipsterParagraph(3, 3, 3, "\n") != "" {
		t.Error("should return empty result with null randomizer")
	}

	data.Data["hipster"]["word"] = old
}
