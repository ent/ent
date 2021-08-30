package sqlcommenter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentEscapse(t *testing.T) {
	cmts := SqlComments{
		"route": `/param first`,
		"num":   `1234`,
		"query": `DROP TABLE FOO'`,
	}

	assert.Equal(t, `num='1234',query='DROP%20TABLE%20FOO%27',route='%2Fparam%20first'`, cmts.Marshal())
}
