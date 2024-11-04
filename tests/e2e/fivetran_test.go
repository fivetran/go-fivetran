package fivetran_test

import (
	testutils "github.com/fivetran/go-fivetran/test_utils"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	testutils.InitE2E()
}
