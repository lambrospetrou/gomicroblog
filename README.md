# gomicroblog
Micro blog service powered by Go - blog.lambrospetrou.com will be based on this!

## Architecture

This section describes the architecture of the micro-blog service.

### Storage/DB

Package *storage* contains the interface **Storager** that defines the functionality that all storage methods should implement (either file-based or DB-based).

Later a *Proxy* can be implemented to wrap the concrete storager to provide cache or speed-up post loading.

**File Storager**

There is gonna be a separate goroutine, that will be responsible to store on-disk new posts to avoid locking in each request. Communication with this service will be done using channels.


## Post creation

For the moment, we will only create posts using the online form but the goal is to be able to write the markdown content of the post off-line using any editor and automatically being available on the website. That's why I reverse to file-based storage system instead of Couchbase.

### Off-line creation - manual file editing

I have to find a way to easily edit the files and make them available to the blog.
This will be done by creating the post directory in the *pending* posts directory and then calling the proper API call on the server (through an admin dashboard) to update the static website with the new post.

#### Idea 1

* Have a folder that will contain all the un-posted/pending/in-progress posts.
* Once a post is done I should either move it manually into the _posts directory or call a function of the server to update the blog posts.
* Still not sure if I want a static blog or a go-powered one
