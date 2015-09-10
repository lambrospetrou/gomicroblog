package post

import (
	"encoding/json"
	"errors"
	"github.com/lambrospetrou/gomicroblog/storage"
	"log"
	"time"
)

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
