package push

// options options for config command
type options struct {
	WorkingDir  string `flag:"wd" value:"" usage:"working dir path to run syncer"`
	WorkingHead string `flag:"wh" value:"" usage:"working head to run syncer"`
	Remote      string `flag:"remote" value:"" usage:"Use specific remote config"`
}
