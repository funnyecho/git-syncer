package options

// BasicOptions is basic options that all command shall support
type BasicOptions struct {
	WorkingDir  string `flag:"working-dir,wd" value:"" usage:"working dir path to run syncer"`
	WorkingHead string `flag:"working-head,wh" value:"" usage:"working head to run syncer"`
	Remote      string `flag:"remote" value:"" usage:"Use specific remote config"`
	Verbose     string `flag:"verbose" value:"info" usage:"Verbose level"`
}
