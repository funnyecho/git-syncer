package runner

// Options basic options for command
type Options struct {
	Base   string `flag:"base" value:"" usage:"Base dir path to run syncer"`
	Branch string `flag:"branch" value:"" usage:"Push a specific branch"`
	Remote string `flag:"remote" value:"" usage:"Push to specific remote"`
}
