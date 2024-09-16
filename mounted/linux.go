//go:build linux
// +build linux

package mounted

func GetDriveLetters() ([]string, error) {
	return []string{"/"}, nil
}
