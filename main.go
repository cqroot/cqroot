package main

import (
	"os"
	"text/template"
)

type Readme struct {
	Repos  [][]Repo
	Badges []string
}

type Repo struct {
	User string
	Name string
}

var readme Readme = Readme{
	Repos: [][]Repo{
		{
			Repo{
				"cqroot",
				"joplin-outline",
			},
			Repo{
				"cheatsheets-cn",
				"cheatsheets-cn",
			},
		},
		{
			Repo{
				"cqroot",
				"openstack-swift-exporter",
			},
			Repo{
				"cqroot",
				"openstack-swift-dashboard",
			},
		},
		{
			Repo{
				"cqroot",
				"zookeeper-ansible",
			},
			Repo{
				"cqroot",
				"kafka-ansible",
			},
		},
	},
	Badges: []string{
		"Go",
		"Kubernetes",
		"Docker",
		"Linux",
		"Neovim",
		"OpenStack",
		"Ceph",
		"Prometheus",
		"Python",
		"Tmux",
	},
}

func main() {
	tmpl, err := template.ParseFiles("./README.tmpl")
	if err != nil {
		panic(err)
	}

	output, err := os.OpenFile("./README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	tmpl.Execute(output, readme)
}
