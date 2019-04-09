package main

import (
	"fmt"
	"io"
	"net/http"
	"errors"
	"bytes"
	"os"
	"strings"
)

func main() {
	urls := []string{"https://raw.githubusercontent.com/alekskivuls/rosalind/master/dna/main.go","https://raw.githubusercontent.com/alekskivuls/rosalind/master/binarySearch/main.go","https://raw.githubusercontent.com/alekskivuls/rosalind/master/dag/main.go"}
	_, printOut := downloadMultipleFiles(urls)

	fmt.Printf("%v", printOut)
}

func DownloadFile(filepath string, url string) ([]byte, error) {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return nil, err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return nil, err
}

func downloadFile(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
        defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func downloadMultipleFiles(urls []string) ([][]byte, error) {
	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))
	for _, URL := range urls {
		go func(URL string) {
			b, err := DownloadFile(fmt.Sprintf("/Users/cheriewoo/go/src/test/%v",strings.ReplaceAll(strings.ReplaceAll(URL, "https://raw.githubusercontent.com/alekskivuls/rosalind/master/", ""), "/main.go", "")), URL)
			// b, err := downloadFile(URL)
			if err != nil {
				errch <- err
				done <- nil
				return
			}
			done <- b
			errch <- nil
		}(URL)
	}
	bytesArray := make([][]byte, 0)
	var errStr string
	for i := 0; i < len(urls); i++ {
		bytesArray = append(bytesArray, <-done)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}
	var err error
	if errStr!=""{
		err = errors.New(errStr)
	}
	return bytesArray, err
}

// func (this *Zipnik) zipData(files [][]byte) {

//     // Create a buffer to write our archive to.
//     fmt.Println("we are in the zipData function")
//     buf := new(bytes.Buffer)

//     // Create a new zip archive.
//     zipWriter := zip.NewWriter(buf)

//     // Add some files to the archive.
//     // var files = []struct {
//     //     Name, Body string
//     // }{
//     //     {"readme.txt", "This archive contains some text files."},
//     //     {"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
//     //     {"todo.txt", "Get animal handling licence.\nWrite more examples."},
//     // }
//     for _, file := range files {
//     zipFile, err := zipWriter.Create(file.Name)
//         if err != nil {
//             fmt.Println(err)
//         }
//         _, err = zipFile.Write([]byte(file))  
//         if err != nil {
//             fmt.Println(err)
//         }
//     }

//     // Make sure to check the error on Close.
//     err := zipWriter.Close()
//     if err != nil {
//         fmt.Println(err)
//     }

//     //write the zipped file to the disk
//     ioutil.WriteFile("demo.zip", buf.Bytes(), 0777)    

// }