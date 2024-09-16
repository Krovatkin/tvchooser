//go:build linux
// +build linux

package mounted

func GetDrivePaths() ([]string, error) {
	return []string{"/"}, nil
}
