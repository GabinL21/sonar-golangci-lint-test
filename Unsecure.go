package samples

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	data := "Hello, World!"
	hash := md5.Sum([]byte(data))
	fmt.Printf("MD5 hash of '%s': %s\n", data, hex.EncodeToString(hash[:]))
}
