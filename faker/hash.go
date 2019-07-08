package faker

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func (f *Faker) SHA256() string {
	data := make([]byte, 32)
	//nolint
	f.Read(data)
	return fmt.Sprintf("%x", sha256.Sum256(data))
}

func (f *Faker) MD5() string {
	data := make([]byte, 16)
	//nolint
	f.Read(data)
	return fmt.Sprintf("%x", md5.Sum(data))
}
