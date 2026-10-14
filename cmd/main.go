package main

import (
	"log/slog"

	"blog-ssg/internal/generator"
)

func main() {

	//TODO: DECIDE HOW TO HANDLE ERRS
	dir := "content/"
	posts, err := generator.MakePosts(dir)
	if err != nil {
		slog.Error("creating html documents", "err", err)
	} else {
		slog.Info("html documents created", "posts", posts)
	}
}
