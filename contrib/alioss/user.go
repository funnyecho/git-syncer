package alioss

import (
	"fmt"
	"strings"
)

func (a *Alioss) getExecutor() string {
	user := a.opts.UserName()
	email := a.opts.UserEmail()

	if email != "" {
		email = fmt.Sprintf("<%s>", email)
	}

	return strings.Join([]string{user, email}, " ")
}
