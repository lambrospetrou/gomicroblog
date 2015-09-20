package post

import (
	"encoding/json"
	//"github.com/russross/blackfriday"
	//"html/template"
	"time"
)

type BPost struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	UrlPermalink    string    `json:"url_permalink"`
	DateCreated     time.Time `json:"date_created"`
	DateEdited      time.Time `json:"date_edited"`
	ContentMarkdown string    `json:"content_markdown"`
	ContentHtml     string    `json:"content_html"`
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
	return a[i].DateEdited.Unix() >= a[j].DateEdited.Unix()
}

func (p *BPost) IdStr() string {
	return p.Id
}

func (p *BPost) FormattedEditedTime() string {
	return p.DateEdited.Format("January 02, 2006 | Monday -- 15:04PM")
}

func (p *BPost) FormattedCreatedTime() string {
	return p.DateCreated.Format("January 02, 2006 | Monday")
}

func (p *BPost) HTML5CreatedTime() string {
	return p.DateCreated.Format("2006-01-02")
}

// FromFile reads a post folder (that follows our special structure) and creates a new
// post structure with fields filled from the file loaded.
func FromFile(pathname string) *BPost {
	//p.ContentHtml = string(blackfriday.MarkdownCommon([]byte(p.ContentMarkdown)))
	return nil
	//return &BPost{DateCreated: time.Now()}
}

func FromJson(b []byte) *BPost {
	bp := &BPost{}
	if err := json.Unmarshal(b, bp); err != nil {
		return nil
	}
	return bp
}

/*
// New creates a new blog post returns it empty setting its creation date to time.Noe
func New() *BPost {
	return &BPost{DateCreated: time.Now()}
}

func (p *BPost) PrepareSave() {
	// update the HTML content
	p.DateEdited = time.Now()
	p.ContentHtml = string(blackfriday.MarkdownCommon([]byte(p.ContentMarkdown)))

	// set the ID the same as the URL friendly link and it will be updated
	// by the storager used if necessary
	p.Id = p.UrlPermalink
}
*/
