package gen

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	SITE_DST      = "_site"
	SITE_DST_PERM = 0755
)

// GenerateHandler is called by the website when we want to execute the generator
func GenerateSite(dir_site string) error {
	fmt.Fprintln(os.Stdout, "dir:", dir_site)
	// prepare the destination site dir
	dst_dir := filepath.Join(dir_site, SITE_DST)
	if err := prepareSiteDest(dst_dir); err != nil {
		return err
	}
	return nil
}

// prepareSiteDest ensures that the destination directory is created and is empty.
// If it already exists the all its contents are deleted!
func prepareSiteDest(dst string) error {
	var path_exists bool = true
	var info os.FileInfo
	var err error
	if info, err = os.Stat(dst); err != nil {
		if os.IsNotExist(err) {
			path_exists = false
		} else {
			return errors.New(fmt.Sprintln("Could not open site destination directory: "+dst, err))
		}
	}
	if path_exists {
		// delete everything that exists in the destination directory
		if !info.IsDir() {
			return errors.New("Destination specified is not directory: " + dst)
		}
		if err = os.RemoveAll(dst); err != nil {
			return err
		}
	}
	if err = os.MkdirAll(dst, SITE_DST_PERM); err != nil {
		return err
	}
	return nil
}
