package storage

import (
	"github.com/lambrospetrou/gomicroblog/post"
	"io/ioutil"
)

type FileStorager struct {
	data_dir string
}

func (fs *FileStorager) Store(p *post.BPost) (string, error) {
	return "blog_id_here", nil
}

func (fs *FileStorager) Load(id string) (*post.BPost, error) {

}

func (fs *FileStorager) Delete(id string) error {

}

func (fs *FileStorager) LoadAll() []post.BPost {
	return make([]post.BPost, 10)
}
