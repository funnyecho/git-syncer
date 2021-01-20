package alioss

import (
	"fmt"
	"strings"
)

func (a *Alioss) getExecutor() string {
	user, _ := a.GetRawConfig("user.name")
	email, _ := a.GetRawConfig("user.email")

	if email != "" {
		email = fmt.Sprintf("<%s>", email)
	}

	return strings.Join([]string{user, email}, " ")
}
