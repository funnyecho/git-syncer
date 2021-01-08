package setup

type Options struct {
	Base   string `flag:"base" usage:"Base dir path to run syncer"`
	Branch string `flag:"branch" usage:"Push a specific branch"`
	Remote string `flag:"remote" value:"" usage:"Push to specific remote"`
}
