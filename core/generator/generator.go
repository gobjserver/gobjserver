package generator

import (
	"time"
)

// CurrentDateTime generates current datetime.
func CurrentDateTime() string {
	return time.Now().Format(time.RFC3339)
}
