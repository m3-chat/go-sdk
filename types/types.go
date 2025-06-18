package types

type ClientOptions struct {
	Stream bool
}

type RequestParams struct {
	Model   string
	Content string
}

type BatchRequestOptions struct {
	Model string
}
