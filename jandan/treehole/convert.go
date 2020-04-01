package treehole

import "strings"

type template struct {
	Post    string
	Comment string
}

var Template = &template{
	Post:    "{{body}}\n==========\n{{comments}}",
	Comment: "{{author}}: {{content}}\n----------",
}

func (th *Treehole) String() string {
	var comments string
	for _, v := range *th.Comments {
		comments += v.String() + "\n"
	}

	body := Template.Post
	body = strings.ReplaceAll(body, "{{body}}", th.Content)
	body = strings.ReplaceAll(body, "{{comments}}", comments)

	return body
}

func (comment *Comment) String() string {
	str := Template.Comment
	str = strings.ReplaceAll(str, "{{author}}", comment.Author)
	str = strings.ReplaceAll(str, "{{content}}", comment.Content)
	return str
}
