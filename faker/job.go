package faker

// JobInfo is a struct of job information
type JobInfo struct {
	Company    string
	Title      string
	Descriptor string
	Level      string
}

// Job will generate a struct with random job information
func (f *Faker) Job() *JobInfo {
	return &JobInfo{
		Company:    f.Company(),
		Title:      f.JobTitle(),
		Descriptor: f.JobDescriptor(),
		Level:      f.JobLevel(),
	}
}

// JobTitle will generate a random job title string
func (f *Faker) JobTitle() string {
	return f.getRandValue([]string{"job", "title"})
}

// JobDescriptor will generate a random job descriptor string
func (f *Faker) JobDescriptor() string {
	return f.getRandValue([]string{"job", "descriptor"})
}

// JobLevel will generate a random job level string
func (f *Faker) JobLevel() string {
	return f.getRandValue([]string{"job", "level"})
}
