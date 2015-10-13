package gen

import (
	"errors"
	"fmt"
	"github.com/lambrospetrou/gomicroblog/view"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	SITE_DST      = "_site"
	SITE_DST_PERM = 0755

	POSTS_DIR_SRC = "_posts"
	POSTS_DIR_DST = "articles"
)

// GenerateHandler is called by the website when we want to execute the generator
func GenerateSite(dir_site string, viewBuilder *view.Builder) error {
	fmt.Fprintln(os.Stdout, "dir:", dir_site)
	// prepare the destination site dir
	dst_dir := filepath.Join(dir_site, SITE_DST)
	if err := prepareSiteDest(dst_dir); err != nil {
		return err
	}

	// iterate over the posts directory and compile each post
	compilePosts(filepath.Join(dir_site, POSTS_DIR_SRC), filepath.Join(SITE_DST, POSTS_DIR_DST), viewBuilder)

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

func compilePosts(src_posts_dir string, dst_posts_dir string, viewBuilder *view.Builder) error {
	//filepath.Walk(src_posts_dir, createWalker())

	files, err := ioutil.ReadDir(src_posts_dir)
	if err != nil {
		return err
	}
	// iterate over all the posts
	for _, cpost := range files {
		fmt.Println("== Visiting file or dir ==")
		fmt.Println(cpost.Name(), cpost.IsDir(), cpost.ModTime())
	}
	return nil
}

func createWalker() filepath.WalkFunc {
	var counter int64 = 0
	return func(path string, info os.FileInfo, err error) error {
		counter++
		if counter == 1 {
			// do not do anything in the root directory
			return nil
		}
		return walkFn(path, info, err)
	}
}

func walkFn(path string, info os.FileInfo, err error) error {
	fmt.Println("== Visiting file or dir ==")
	fmt.Println(path, info.IsDir(), info.Name(), info.ModTime())
	// skip the post directories recursion
	if info.IsDir() {
		return filepath.SkipDir
	}
	return nil
}
