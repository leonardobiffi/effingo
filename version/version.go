package version

var version = "v0.0.0-dev"
var versionTemplate = `Effingo {{printf "%s" .Version}}
`

func Version() string {
	return version
}

func VersionTemplate() string {
	return versionTemplate
}
