package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Country struct {
	Id    string
	Short string
	Code  string
}

func (c *Country) shout() string {
	result := c.Id + ": " + strings.ToLower(c.Short) + " " + c.Code
	return fmt.Sprint(result)
}

func (c *Country) name() string {
	return strings.ToLower(c.Short)
}

func (c *Country) code() string {
	return strings.ToLower(c.Code)
}

func main() {

	// 1. Read CSV file
	lines, err := ReadCsv("countries.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		value := strings.Split(line[0], ";")
		data := Country{
			Id:    value[0],
			Short: value[1],
			Code:  value[2],
		}
		data.shout()
		if fileExists("svg/" + data.name() + ".svg") {
			fmt.Println(data.shout() + " file exists")
			copy("svg/"+data.name()+".svg", "svg2/flags-"+data.name()+"-"+data.code()+".svg")
		} else {
			fmt.Println(data.shout() + " file does not exist")
		}
	}

	// 2. Find SVG file

	// 3. Rename it

}

// ReadCsv accepts a file and
// returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func copy(src string, dst string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
