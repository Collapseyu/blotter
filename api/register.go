package api

import (
	"github.com/OhYee/blotter/register"
)

func Register() {
	register.Register(
		"friends",
		Friends,
	)
	register.Register(
		"menus",
		Menus,
	)
	register.Register(
		"post",
		Post,
	)
	register.Register(
		"posts",
		Posts,
	)
	register.Register(
		"markdown",
		Markdown,
	)
	register.Register(
		"comments",
		Comments,
	)
	register.Register(
		"variables",
		Variables,
	)
	register.Register(
		"layout",
		Layout,
	)
	register.Register(
		"tags",
		Tags,
	)
	register.Register(
		"avatar",
		Avatar,
	)
	register.Register(
		"comment/add",
		CommentAdd,
	)
	register.Register(
		"mail",
		Mail,
	)
}
