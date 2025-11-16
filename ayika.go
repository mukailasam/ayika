package ayika

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var once sync.Once
var debug bool

// Load .env automatically ONLY ONCE
func init() {
	once.Do(func() {
		debug = strings.ToLower(os.Getenv("AYIKA_DEBUG")) == "true"
		log("[ayika] starting automatic environment load…")
		err := LoadEnv(".env")
		if err != nil {
			log("[ayika] warning: .env loaded with error:", err)
		} else {
			log("[ayika] .env loaded successfully")
		}
	})
}

// LoadEnv loads a .env file manually
func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// split KEY=VALUE
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // ignore malformed lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		value = strings.Trim(value, `"'`)

		os.Setenv(key, value)
	}

	return scanner.Err()
}

func log(message ...interface{}) {
	if debug {
		t := time.Now().Format("15:04:05")
		fmt.Println(t, message)
	}
}
