package storage

import (
	"errors"
	"github.com/lambrospetrou/gomicroblog/post"
	"io/ioutil"
	"os"
	"time"
)

// postDateFilename returns the Date part of the post filename
func postDateFilename(dt time.Time) string {
	return dt.Format("20060102")
}

// postFileName returns the filename to be used for the blog post
// when stored on disk.
func postFilename(data_dir string, bp *post.BPost) string {
	return data_dir + "/" + postDateFilename(bp.DateCreated) + "-" + id + ".dat"
}

///////////////////////////////////////////////////
//----------------
///////////////////////////////////////////////////

type fssClientBundleAdd struct {
	ChResult chan error
	PostNew  *post.BPost
}

type fssClientBundleRem struct {
	ChResult chan error
	PostId   string
}

type FileSavingService struct {
	ChPostSave   chan fssClientBundleAdd
	ChPostRemove chan fssClientBundleRem
	ChTerminate  chan bool
	DataDir      string
}

func newFileSavingService(dt_dir string) *FileSavingService {
	return &FileSavingService{
		make(chan fssClientBundleAdd),
		make(chan fssClientBundleRem),
		make(chan bool),
		dt_dir,
	}
}

func (fss *FileSavingService) run() {
	go func(ChTerminate <-chan bool,
		ChNewPost <-chan fssClientBundleAdd,
		ChRemPost <-chan fssClientBundleRem) {
		for {
			select {
			case <-ChTerminate:
				return
			case np := <-ChNewPost:
				// save the file and return any error at the response channel
				jsonBytes, err := json.Marshal(np.PostNew)
				if err != nil {
					np.ChResult <- errors.New("Could not convert post to JSON format!")
				}
				ioutil.WriteFile(postFileName(fss.DataDir, np.PostNew), jsonBytes, 0777)
				np.ChResult <- nil
			case dp := <-ChRemPost:
				src_f := postFileName(fss.DataDir, dp.PostRem)
				if os.Remove(src_f) != nil {
					dp.ChResult <- errors.New("Could not delete post: " + id)
				}
				dp.ChResult <- nil
			}
		}
	}(fss.ChTerminate, fss.ChPostSave, fss.ChPostRemove)
}

// FileStorager is the main object that will be handling the post saving on disk
// using a file-based approach instead of using a database to allow later easy addition
// of posts without requiring updates in a database.
type FileStorager struct {
	data_dir string
	fss      *FileSavingService
}

// New returns a new FileStorager that will use the given dir as a root directory
// to store, load and delete the blog posts
func New(dir string) FileStorager {
	fs := FileStorager{dir}
	fs.fss = NewFileSavingService()
}

func (fs *FileStorager) Store(p *post.BPost) error {
	// we have to allocate a channel and send it to the FileSavingService
	chResult := make(chan error)
	fs.fss <- fssClientBundleAdd{chResult, p}
	if err := <-chResult; err != nil {
		return errors.New("Could not store new post!")
	}
	return nil
}

func (fs *FileStorager) Delete(id string) error {

}

func (fs *FileStorager) Load(id string) (*post.BPost, error) {
	src_f := postFileName(fs, id)
	if bytes, err := ioutil.ReadFile(src_f); err != nil {
		return nil, errors.New("Could not find post: " + id)
	}
	return post.FromJson(bytes), nil
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
