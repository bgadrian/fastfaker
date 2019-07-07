package faker

import "fmt"

func ExampleFaker_SHA256() {
	Global.Seed(11)
	fmt.Println(Global.SHA256())
	// Output: e33789b87c2858fbfece97c81533b496493700ff3845d3f80e413ab588859dff
}

func ExampleFaker_MD5() {
	Global.Seed(11)
	fmt.Println(Global.MD5())
	// Output: 605410f97df9dc83d1e469fb3ee96119
}
