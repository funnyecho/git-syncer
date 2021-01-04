package setup

type Options struct {
	Branch string `flag:"branch" usage:"Push a specific branch"`
	Remote string `flag:"remote" value:"production" usage:"Push to specific remote"`
}
