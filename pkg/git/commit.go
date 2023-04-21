package git

import "time"

type commitBody struct {
	Sha         string      `json:"sha"`
	NodeID      string      `json:"node_id"`
	Commit      commit      `json:"commit"`
	URL         string      `json:"url"`
	HTMLURL     string      `json:"html_url"`
	CommentsURL string      `json:"comments_url"`
	Author      interface{} `json:"author"`
	Committer   interface{} `json:"committer"`
	Parents     []parent    `json:"parents"`
}

type parent struct {
	Sha     string `json:"sha"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type verification struct {
	Verified  bool        `json:"verified"`
	Reason    string      `json:"reason"`
	Signature interface{} `json:"signature"`
	Payload   interface{} `json:"payload"`
}

type commitUser struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  time.Time `json:"date"`
}

type tree struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}

type commit struct {
	Author       commitUser   `json:"author"`
	Committer    commitUser   `json:"committer"`
	Message      string       `json:"message"`
	Tree         tree         `json:"tree"`
	URL          string       `json:"url"`
	CommentCount int          `json:"comment_count"`
	Verification verification `json:"verification"`
}
