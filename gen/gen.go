package gen

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lambrospetrou/gomicroblog/post"
	"github.com/lambrospetrou/gomicroblog/view"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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
	compilePosts(filepath.Join(dir_site, POSTS_DIR_SRC), filepath.Join(dst_dir, POSTS_DIR_DST), viewBuilder)

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
	files, err := ioutil.ReadDir(src_posts_dir)
	if err != nil {
		return err
	}
	// create destination directory
	if err = os.MkdirAll(dst_posts_dir, SITE_DST_PERM); err != nil {
		return err
	}
	// iterate over all the posts
	for _, cpost := range files {
		fmt.Println("== Compiling article-post ==")
		fmt.Println(cpost.Name(), cpost.IsDir(), cpost.ModTime())
		if cpost.IsDir() {
			fmt.Println(compileDirPost(cpost, dst_posts_dir, viewBuilder))
		} else {
			fmt.Println(compileSinglePost(cpost, dst_posts_dir, viewBuilder))
		}
	}
	return nil
}

func compileSinglePost(info os.FileInfo, dst_posts_dir string, viewBuilder *view.Builder) error {
	f, err := os.Create(filepath.Join(dst_posts_dir, info.Name()[:len(info.Name())-3]))
	if err != nil {
		return err
	}
	bundle := &view.TemplateBundle{
		Footer: &view.FooterStruct{Year: time.Now().Year()},
		Header: &view.HeaderStruct{Title: "Single Post"},
		Post:   post.FromFile(info.Name()),
	}
	err = viewBuilder.Render(bufio.NewWriter(f), view.LAYOUT_POST, bundle)
	fmt.Println("Error: ", err)
	return err
}

func compileDirPost(info os.FileInfo, dst_posts_dir string, viewBuilder *view.Builder) error {
	return nil
}
