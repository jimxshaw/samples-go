package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/jimxshaw/testing/buffaloSample/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
