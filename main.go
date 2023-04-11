package main

import (
	"fmt"
	"forum/repo"
)

func main() {
	fmt.Println("kirdim")
	r, e := repo.New()
	fmt.Println(r, e)
	// if e != nil {
	// 	fmt.Println("oibai owibka: ", e)
	// 	os.Exit(1)
	// }
	// fmt.Println("repo awtim")

	// req := &entity.Post{
	// 	UserId:    1,
	// 	Title:     "Ekinwi bolma",
	// 	Content:   "asdfasfasfdasfdsafd safdsafdsafdsafdsafd sadfsadf",
	// 	CreatedAt: time.Now(),
	// }
	// fmt.Println("post dayin", req)

	// fmt.Println(r.CreatePost(req))
}
