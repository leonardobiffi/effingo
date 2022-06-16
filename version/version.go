/*
Copyright © 2022 Leonardo Biffi <leonardobiffi@outlook.com>
*/
package version

var Version = "v0.0.0-dev"
var template = `Effingo {{printf "%s" .Version}}
`

func String() string {
	return Version
}

func Template() string {
	return template
}
