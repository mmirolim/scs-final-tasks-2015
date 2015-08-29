// hide and search for file 
// walk each dir by seperate goroutine
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	root := "dirtree"
	fname := "find.me"
	err := hideInDirForest(root, fname)
	if err != nil {
		log.Fatal(err)
	}
	found := make(chan string)
	go func() {
		// find file
		err = findInDirForest(root, fname, found)
		fatalOnErr(err)
	}()

	fmt.Println("found", <-found)
	fmt.Println("Execution time", time.Since(start))
	start = time.Now()
	seqSearch(root, fname)
	fmt.Println("Execution time of seqSearch", time.Since(start))
}

func findInDirForest(root, fname string, found chan string) error {
	dir, err := os.Open(root)
	defer dir.Close()
	if err != nil {
		return err
	}

	fls, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, d := range fls {
		go crawler(root+"/"+d.Name(), fname, found)
	}

	return nil
}

func crawler(root, target string, found chan string) {
	f, err := os.Open(root)
	defer f.Close()
	if err != nil {
		log.Println("Open by name", root, err)
		return
	}
	fls, err := f.Readdir(-1)
	if err != nil {
		log.Println("ReadDir", f.Name(), err)
		return
	}

	if len(fls) > 1 {
		for _, d := range fls {
			go crawler(root+"/"+d.Name(), target, found)
		}

	}

	if len(fls) == 1 && fls[0].Name() == target {
		if err != nil {
			log.Println("getwd err", err)
			return
		}
		found <- root
	}
}
func hideInDirForest(root, fname string) error {
	rootdir := root
	err := os.Mkdir(rootdir, os.ModePerm)
	if err != nil {
		return err
	}
	max := 10
	for i := 0; i < max; i++ {
		lvl1 := rootdir + "/branch" + strconv.Itoa(i)
		err := os.Mkdir(lvl1, os.ModePerm)
		fatalOnErr(err)
		for j := 0; j < max; j++ {
			lvl2 := lvl1 + "/branch" + strconv.Itoa(j)
			err := os.Mkdir(lvl2, os.ModePerm)
			fatalOnErr(err)
			for n := 0; n < max; n++ {
				lvl3 := lvl2 + "/branch" + strconv.Itoa(n)
				err := os.Mkdir(lvl3, os.ModePerm)
				fatalOnErr(err)
			}
		}
	}
	// hide file
	path := ""
	for i := 0; i < 3; i++ {
		path += "/branch" + strconv.Itoa(rand.Intn(max))
	}
	_, err = os.Create(rootdir + path + "/" + fname)
	fmt.Println(rootdir + path + "/" + fname)
	fatalOnErr(err)
	return err

}

func seqSearch(root, fname string) {
	max := 10
	for i := 0; i < max; i++ {
		lvl1 := root + "/branch" + strconv.Itoa(i)
		for j := 0; j < max; j++ {
			lvl2 := lvl1 + "/branch" + strconv.Itoa(j)
			for n := 0; n < max; n++ {
				lvl3 := lvl2 + "/branch" + strconv.Itoa(n)
				f, err := os.Open(lvl3)
				fatalOnErr(err)
				fls, err := f.Readdir(-1)
				fatalOnErr(err)
				if len(fls) == 1 && fls[0].Name() == fname {
					fmt.Println("found in ", lvl3)
					return
				}
			}
		}
	}
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
