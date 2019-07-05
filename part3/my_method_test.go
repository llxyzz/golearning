package part3

import (
	"fmt"
	"testing"
	"time"
)

func TestMyMethod(t *testing.T) {

	const hour = time.Hour

	fmt.Println(hour.Seconds())

	p := Sub{1, 2}

	fmt.Println(p.getSubtraction(p))

}

func TestCache(t *testing.T) {

}
