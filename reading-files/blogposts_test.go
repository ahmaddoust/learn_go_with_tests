package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/ahmaddoust/reading-files"
)

func TestBlogPosts(t *testing.T) {
	//Given our file system
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hello, world")},
		"hello-twitch.md": {Data: []byte("hello, twitch")},
	}

	// When we call our posts from fs
	posts := blogposts.PostsFromFS(fs)

	//Then we expect to have the same number of posts as we do files in the file system
	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d posts", len(fs), len(posts))
	}

}
