package main

import (
	//"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var count uint64 = 0
	//data := [][]string{}
	//data = append(data, []string{"Path", "Size", "Filename"})
	err := filepath.Walk("E:\\",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				//log.Println(err)
				return filepath.SkipDir
			}
			parts := strings.Split(path, string(os.PathSeparator))
			if strings.EqualFold(parts[1], "$RECYCLE.BIN") {
				return nil
			}
			if strings.EqualFold(filepath.Ext(filepath.Base(path)), ".JPG") {
				count++
				//fmt.Println(path + "\t" + strconv.FormatInt(info.Size(), 10) + "\t\t" + info.Name())
				fmt.Printf("Count=%s\tPath=%s\tSize=%s\tFilename=%s\n", strconv.FormatUint(count, 10), path, strconv.FormatInt(info.Size(), 10), info.Name())
				//data = append(data, []string{strconv.FormatUint(count, 10), path, strconv.FormatInt(info.Size(), 10), info.Name()})
			}
			//if err := csvExport(data); err != nil {
			//	log.Fatal(err)
			//}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

/* func csvExport(data [][]string) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
		}
	}
	return nil
} */
