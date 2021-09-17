package main

import "fmt"

type Repository struct {
	repoLink    string
	cloneMethod typeOfLink
}

func (r *Repository) setCloneMethod(linkType typeOfLink) {
	r.cloneMethod = linkType
}

func (r *Repository) callClone() {
	fmt.Println("Start cloning")
	r.cloneMethod.clone(r)
}

type typeOfLink interface {
	clone(*Repository)
}

type https struct {}

func (h https) clone(rep *Repository) {
	fmt.Println("Cloning repository via HTTPS")
}

type ssh struct {}

func (s ssh) clone(rep *Repository){
	fmt.Println("Cloning repository via SSH")
}

func main() {
	var rep Repository

	doClone := func(s string) {
		if s == "ssh" {
			rep.setCloneMethod(ssh{})
		} else {
			rep.setCloneMethod(https{})
		}
		rep.callClone()
	}

	doClone("ssh")
}