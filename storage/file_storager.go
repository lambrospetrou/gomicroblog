package storage

import (
	"errors"
	"github.com/lambrospetrou/gomicroblog/post"
	"io/ioutil"
	"os"
)

type FileStorager struct {
	data_dir string
}

// New returns a new FileStorager that will use the given dir as a root directory
// to store, load and delete the blog posts
func New(dir string) FileStorager {
	return FileStorager{dir}
}

func (fs *FileStorager) Store(p *post.BPost) (string, error) {
	id := p.IdStr()
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New("Could not convert post to JSON format!")
	}
	ioutil.WriteFile(postFileName(fs, id), jsonBytes, 0777)
	return id, nil
}

func (fs *FileStorager) Load(id string) (*post.BPost, error) {
	src_f := postFileName(fs, id)
	if bytes, err := ioutil.ReadFile(src_f); err != nil {
		return nil, errors.New("Could not find post: " + id)
	}
	return post.FromJson(bytes), nil
}

func (fs *FileStorager) Delete(id string) error {
	src_f := postFileName(fs, id)
	if os.Remove(src_f) != nil {
		return errors.New("Could not delete post: " + id)
	}
	return nil
}

func (fs *FileStorager) LoadAll() ([]post.BPost, error) {
	if files, err := ioutil.ReadDir(fs.data_dir); err != nil {
		return nil, errors.New("Could not load blog posts")
	}
	posts := make([]post.BPost, len(files))
	for _, f := range files {
		posts = append(posts, post.FromJson(ioutil.ReadFile(fs.data_dir+"/"+f.Name())))
	}
	return posts
}

func postFileName(fs *FileStorager, id string) string {
	return fs.data_dir + "/" + id + ".dat"
}
