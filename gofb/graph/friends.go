package graph

const (
	BASE_GRAPH_URL = "https://graph.facebook.com/%v/friends"
)

type Friends struct {
}

func (f *Friends) Query() string {
	return ""
}
