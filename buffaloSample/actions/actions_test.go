package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/jimxshaw/testing/buffaloSample/actions"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(App())}
	suite.Run(t, as)
}
