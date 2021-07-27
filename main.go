package main

import (
    "fmt"
	"flag"
	"regexp"
    "io/ioutil"
    "log"
	"path"
)

func main() {
	
	var limit int
	flag.IntVar(&limit, "l", 10, "")
	flag.IntVar(&limit, "limit", 10, "")
	flag.Parse()

	dir := flag.Arg(0)
	if flag.NArg() < 1 {
		fmt.Println("Usage:\n\t\tzenpli-prueba DIR")
		fmt.Println("Optional:\n\t\t-l, --limit\tSets a limit of blocks per drone run. Default: 10")
		return
	}

	files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

	re := regexp.MustCompile(`in([\d]+).txt`,)
    for _, file := range files {
		matched := re.MatchString(file.Name())
		inPath := path.Join(dir, file.Name())
		if !file.IsDir() && matched {
			nbr := re.FindAllSubmatch([]byte(inPath), -1)[0][1]
			outPath :=  path.Join(dir, fmt.Sprintf("out%s.txt", nbr))
			processDroneRoute(inPath, outPath, limit)
		}
    }
}

