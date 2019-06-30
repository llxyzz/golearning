package part3

import (
	"fmt"
	"testing"
	"time"
)

func TestMyMethod(t *testing.T) {

	const hour = time.Hour

	fmt.Println(hour.Seconds())

}
