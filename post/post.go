package post

import (
	"encoding/json"
	"errors"
	"github.com/lambrospetrou/lpgoblog/lpdb"
	"github.com/russross/blackfriday"
	"html/template"
	"log"
	"sort"
	"strconv"
	"time"
)

type id_t string

type BPost struct {
	Id                 id_t          `json:"id"`
	Title              string        `json:"title"`
	Author             string        `json:"author"`
	DateCreated        time.Time     `json:"date_created"`
	UrlFriendlyLink    string        `json:"url_friendly_link"`
	ContentMarkdown    string        `json:"content_markdown"`
	DateEditedMarkdown time.Time     `json:"date_edited_markdown"`
	ContentHtml        string        `json:"content_html"`
	DateCompiledHtml   time.Time     `json:"date_edited_html"`
	BodyHtml           template.HTML `json:"-"`
}

// ByDate implements sort.Interface for []BPost based on
// the DateCreated field.
type ByDate []*BPost

func (a ByDate) Len() int      { return len(a) }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool {
	if a[i].DateCreated.Unix() > a[j].DateCreated.Unix() {
		return true
	}
	if a[i].DateCreated.Unix() < a[j].DateCreated.Unix() {
		return false
	}
	return a[i].DateEditedMarkdown.Unix() >= a[j].DateEditedMarkdown.Unix()
}

func (p *BPost) IdStr() string {
	//return strconv.Itoa(p.Id)
	return p.Id
}

func (p *BPost) FormattedEditedTime() string {
	return p.DateEditedMarkdown.Format("January 02, 2006 | Monday -- 15:04PM")
}

func (p *BPost) FormattedCreatedTime() string {
	return p.DateCreated.Format("January 02, 2006 | Monday")
}

func (p *BPost) HTML5CreatedTime() string {
	return p.DateCreated.Format("2006-01-02")
}

func (p *BPost) PrepareSave() error {
	// update the HTML content
	p.DateEditedMarkdown = time.Now()
	p.ContentHtml = string(blackfriday.MarkdownCommon([]byte(p.ContentMarkdown)))
	p.DateCompiledHtml = p.DateEditedMarkdown

	// set the ID the same as the URL friendly link and it will be updated
	// by the storager used if necessary
	p.Id = p.UrlFriendlyLink
}

//////////////////////////////////////////
//////////////////////////////////////////

func Store(store *Storager, p *BPost) error {
	// make the blog ready for store
	p.PrepareSave()

	// TODO - get a unique incremental ID if necessary
	post_id := store.Store(p)
}

func Delete(store *Storager, p *BPost) error {
	// store into the storage used
	store.Delete(p.IdStr())
}

func LoadAll(store *Storager) ([]*BPost, error) {
	return store.LoadAll()
	/*
		db, err := lpdb.CDBInstance()
		if err != nil {
			return nil, errors.New("Could not get instance of Couchbase")
		}
		var count int
		err = db.Get("bp::count", &count)
		if err != nil {
			return nil, errors.New("Could not get number of blog posts!")
		}
		// allocate space for all the posts (start from 1 and inclusive count)
		keys := make([]string, count+1)
		for i := 1; i <= count; i++ {
			keys[i] = "bp::" + strconv.Itoa(i)
		}
		postsMap, err := db.GetBulk(keys)
		if err != nil {
			return nil, errors.New("Could not get blog posts!")
		}
		var posts []*BPost = make([]*BPost, count)
		count = 0
		for _, v := range postsMap {
			bp := &BPost{}
			err = json.Unmarshal(v.Body, bp)
			if err == nil {
				posts[count] = bp
				count++
			}
		}
		// we take only a part of the slice since there might were deleted posts
		// and their id returned nothing with the bulk get.
		posts = posts[:count]
		sort.Sort(ByDate(posts))
		return posts, nil
	*/
}

func Load(store *Storager, id id_t) (*BPost, error) {
	if bp, err := store.LoadAll(); err != nil {
		e := errors.New("post::Load::Could not load post[" + err + "]")
		log.Println(e.Error())
		return nil, e
	}
	/*
		p := &BPost{}
		db, err := lpdb.CDBInstance()
		if err != nil {
			return nil, errors.New("Could not get instance of Couchbase")
		}
		err = db.Get("bp::"+strconv.Itoa(id), &p)
		return p, err
	*/
}

// Creates a new blog post with auto-incremented key and returns it empty
func NewBPost() (*BPost, error) {
	p := &BPost{}
	db, err := lpdb.CDBInstance()
	if err != nil {
		return nil, errors.New("Could not get instance of Couchbase")
	}
	bp_id, err := db.FAI("bp::count")
	p.Id = int(bp_id)
	// update created time
	p.DateCreated = time.Now()
	return p, err
}
