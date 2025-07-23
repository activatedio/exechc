package exechc_test

import (
	"testing"

	"github.com/activatedio/exechc"
	"github.com/stretchr/testify/assert"
)

func TestChecker_Check(t *testing.T) {

	a := assert.New(t)

	cases := map[string]struct {
		cmd        string
		expression string
		assert     func(got bool, err error)
	}{
		"match": {
			cmd:        "echo -n 'test'",
			expression: "stdout == 'test'",
			assert: func(got bool, err error) {
				a.True(got)
				a.Nil(err)
			},
		},
		"no match": {
			cmd:        "echo -n 'test'",
			expression: "stdout == 'other'",
			assert: func(got bool, err error) {
				a.False(got)
				a.Nil(err)
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {

			unit := exechc.NewChecker(&exechc.Runtime{
				CheckCmd:        v.cmd,
				CheckExpression: v.expression,
			})

			v.assert(unit.Check())
		})
	}
}
