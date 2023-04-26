package git

type fileBody struct {
	Sha         string `json:"sha"`
	Filename    string `json:"filename"`
	Status      string `json:"status"`
	Additions   int    `json:"additions"`
	Deletions   int    `json:"deletions"`
	Changes     int    `json:"changes"`
	BlobURL     string `json:"blob_url"`
	RawURL      string `json:"raw_url"`
	ContentsURL string `json:"contents_url"`
	Path        string `json:"patch"`
}

type version struct {
	Name          string
	OldVersion    string
	NewVersion    string
	OldIntVersion int
	NewIntVersion int
}

type fileResponse struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	GitUrl      string `json:"git_url"`
	DownloadUrl string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		Html string `json:"html"`
	} `json:"_links"`
}

type updatedFileBody struct {
	Message   string    `json:"message"`
	Content   string    `json:"content"`
	Committer committer `json:"committer"`
	Sha       string    `json:"sha"`
	Branch    string    `json:"branch"`
}

type committer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TreeResponseStruct struct {
	Sha  string `json:"sha"`
	Url  string `json:"url"`
	Tree []struct {
		Path string `json:"path"`
		Mode string `json:"mode"`
		Type string `json:"type"`
		Size int    `json:"size"`
		Sha  string `json:"sha"`
		Url  string `json:"url"`
	} `json:"tree"`
	Truncated bool `json:"truncated"`
}
