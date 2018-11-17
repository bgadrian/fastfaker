package faker

import (
	"fmt"
	"testing"
)

func ExampleFaker_Job() {
	Global.Seed(11)
	jobInfo := Global.Job()
	fmt.Println(jobInfo.Company)
	fmt.Println(jobInfo.Title)
	fmt.Println(jobInfo.Descriptor)
	fmt.Println(jobInfo.Level)
	// Output: Moen, Pagac and Wuckert
	// Developer
	// National
	// Integration
}

func BenchmarkJob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Job()
	}
}

func ExampleFaker_JobTitle() {
	Global.Seed(11)
	fmt.Println(Global.JobTitle())
	// Output: Director
}

func BenchmarkJobTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.JobTitle()
	}
}

func ExampleFaker_JobDescriptor() {
	Global.Seed(11)
	fmt.Println(Global.JobDescriptor())
	// Output: Central
}

func BenchmarkJobDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.JobDescriptor()
	}
}

func ExampleFaker_JobLevel() {
	Global.Seed(11)
	fmt.Println(Global.JobLevel())
	// Output: Assurance
}

func BenchmarkJobLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.JobLevel()
	}
}
