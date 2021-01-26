package runners

// UseFlagRemote get remote from flagset
func UseFlagRemote(args ...string) (string, error) {
	return useFlag(flagRemote)
}
