package gen

import (
	"bufio"
	"errors"
	"fmt"
	lpio "github.com/lambrospetrou/go-utils/io"
	"github.com/lambrospetrou/gomicroblog/config"
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
func GenerateSite(dir_site string, viewBuilder *view.Builder, confPath string) error {
	fmt.Fprintln(os.Stdout, "dir:", dir_site)
	// prepare the destination site dir
	dst_dir := filepath.Join(dir_site, SITE_DST)
	if err := prepareSiteDest(dst_dir); err != nil {
		return err
	}

	// iterate over the posts directory and compile each post
	if err := compilePosts(filepath.Join(dir_site, POSTS_DIR_SRC), filepath.Join(dst_dir, POSTS_DIR_DST), viewBuilder); err != nil {
		return err
	}

	conf := config.FromConfiguration(confPath)
	// copy all the static paths
	for _, path := range conf.StaticPaths {
		fmt.Println("path", path)
		if err := lpio.Copy(path, filepath.Join(dst_dir, filepath.Base(path)), SITE_DST_PERM); err != nil {
			return err
		}
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
			fmt.Println(compileDirPost(src_posts_dir, cpost, dst_posts_dir, viewBuilder))
		} else {
			fmt.Println(compileSinglePost(src_posts_dir, cpost, dst_posts_dir, viewBuilder))
		}
	}
	return nil
}

func compileSinglePost(src_post_dir string, info os.FileInfo, dst_posts_dir string, viewBuilder *view.Builder) error {
	postName := getPostNameFromRaw(info)
	postDir := filepath.Join(dst_posts_dir, postName)
	// create the directory of the post
	if err := os.MkdirAll(postDir, SITE_DST_PERM); err != nil {
		return err
	}

	// get the markdown filename
	postMarkdownPath := filepath.Join(src_post_dir, info.Name())

	// copy the markdown file to the directory
	if err := lpio.CopyFile(postMarkdownPath, filepath.Join(postDir, info.Name()), SITE_DST_PERM); err != nil {
		return err
	}
	return generateHTMLFromMarkdown(postMarkdownPath, filepath.Join(postDir, "index.html"), viewBuilder)
}

func compileDirPost(src_post_dir string, info os.FileInfo, dst_posts_dir string, viewBuilder *view.Builder) error {
	postName := getPostNameFromRaw(info)
	srcPostDir := filepath.Join(src_post_dir, info.Name())
	dstPostDir := filepath.Join(dst_posts_dir, postName)
	if err := lpio.Copy(srcPostDir, dstPostDir, SITE_DST_PERM); err != nil {
		return err
	}

	return generateHTMLFromMarkdown(filepath.Join(srcPostDir, info.Name()+".md"),
		filepath.Join(dstPostDir, "index.html"), viewBuilder)
}

// getPostNameFromRaw() returns the post name by discarding the date at the front and any extension
func getPostNameFromRaw(info os.FileInfo) string {
	if info.IsDir() {
		return info.Name()[11:]
	}
	return info.Name()[11 : len(info.Name())-3]
}

// generateHTMLFromMarkdown() generates the HTML of the given markdown file 'postMarkdownPath' and stores it in the file denoted by 'postDstPath'.
func generateHTMLFromMarkdown(postMarkdownPath string, postDstPath string, viewBuilder *view.Builder) error {
	p := post.FromFile(postMarkdownPath)
	// create the actual HTML file for the post
	bundle := &view.TemplateBundle{
		Footer: &view.FooterStruct{Year: time.Now().Year()},
		Header: &view.HeaderStruct{Title: p.Title},
		Post:   p,
	}
	f, err := os.Create(postDstPath)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = viewBuilder.Render(w, view.LAYOUT_POST, bundle)
	w.Flush()
	return err
}
