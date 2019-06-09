package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/writeas/go-writeas"
	"github.com/writeas/wf-import"
	"io/ioutil"
	"os"
)

func main() {
	// Get parameters
	u := flag.String("u", "", "WriteFreely username")
	host := flag.String("h", "write.as", "WriteFreely host")
	flag.Parse()

	// Validate parameters
	args := flag.Args()
	if *u == "" || len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: wfimport -u username [-h example.com] file1\n")
		os.Exit(1)
	}
	fn := args[0]

	// Get password
	fmt.Print("Password: ")
	pass, err := gopass.GetPasswdMasked()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading pass: %v\n", err)
		os.Exit(1)
	}
	// Validate password
	if len(pass) == 0 {
		fmt.Fprintf(os.Stderr, "Please enter your password.\n")
		os.Exit(1)
	}

	// Create Write.as client
	cl := writeas.NewClientWith(writeas.Config{
		URL: "https://" + *host + "/api",
	})

	// Log user in
	fmt.Printf("Logging in to %s...", *host)
	_, err = cl.LogIn(*u, string(pass))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("OK\n")
	defer func() {
		fmt.Print("Logging out...")
		cl.LogOut()
		fmt.Print("OK\n")
	}()

	// Read file contents
	// TODO: validate
	fmt.Print("Reading file...")
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s: %v\n", fn, err)
		os.Exit(1)
	}
	fmt.Print("OK\n")

	imp := wfimport.Import{}

	fmt.Print("Parsing file...")
	err = json.Unmarshal(content, &imp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s: %v\n", fn, err)
		os.Exit(1)
	}
	fmt.Print("OK\n")
	fmtln("Read user %s export.", imp.User.Username)
	fmtln("Found %d collection(s).", len(imp.Collections))
	fmtln("Found %d draft post(s).", len(imp.Posts))

	// Create collections and their posts
	for _, coll := range imp.Collections {
		fmt.Printf("%s has %d post(s). ", coll.Alias, len(*coll.Posts))
		if len(*coll.Posts) == 0 {
			fmt.Print("Skipping.\n")
			continue
		}
		fmt.Print("\n")

		fmt.Printf("Creating collection %s...", coll.Alias)
		_, err = cl.CreateCollection(&writeas.CollectionParams{
			Alias:       coll.Alias,
			Title:       coll.Title,
			Description: coll.Description,
			// TODO:
			//Stylesheet:  coll.Stylesheet,
			//Public: coll.Public,
		})
		if err != nil {
			// TODO: handle alias collisions
			// TODO: handle hitting collection allowance limit
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		fmt.Print("OK\n")

		// Create posts
		for _, p := range *coll.Posts {
			fmt.Printf("Creating post %s...", p.Slug)
			_, err = cl.CreatePost(&writeas.PostParams{
				Slug:       p.Slug,
				Title:      p.Title,
				Content:    p.Content,
				Font:       p.Font,
				Language:   p.Language,
				IsRTL:      p.RTL,
				Created:    &p.Created,
				Updated:    &p.Updated,
				Collection: coll.Alias,
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			fmt.Print("OK\n")
		}
	}

	// Create anonymous / draft posts
	for _, p := range imp.Posts {
		fmt.Printf("Creating draft post from %s...", p.ID)
		_, err = cl.CreatePost(&writeas.PostParams{
			Title:    p.Title,
			Content:  p.Content,
			Font:     p.Font,
			Language: p.Language,
			IsRTL:    p.RTL,
			Created:  &p.Created,
			Updated:  &p.Updated,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		fmt.Print("OK\n")
	}
}

func fmtln(s string, v ...interface{}) {
	fmt.Printf(s+"\n", v...)
}
