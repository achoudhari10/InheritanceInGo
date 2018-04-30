package main

import "fmt"

func main() {
	A1 :=Author{"Stephen","life science"}
	p1 :=Post{"Black Bang","Book Name",A1}
	c1 :=content{"c1",p1}

	A2 :=Author{"Markow","data science"}
	p2 :=Post{"Big data","Book Name2",A2}
	c2 := content{"c2",p2}
	wd :=website{
		contents :[]content{c1,c2},
	}
	wd.display()
}

type Author struct{
	name string
	bio string
}

type Post struct{
	Title string
	book string
	Author
}

type content struct{
	nameofc string
	Post
}
func (p content)Details (){
	fmt.Printf("Title of Post %s\n",p.Post.Title)
	fmt.Printf("Bio %s\n",p.Post.Author.bio)
	fmt.Printf("Content %s\n",p.nameofc)
}

type website struct{
	contents []content
}

func (w website)display(){
	fmt.Println("Contents of Website")
	for _,post :=range w.contents{
		post.Details()
	}
}