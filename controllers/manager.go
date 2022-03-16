package manager

import (
	"fmt"
	"net/http"
)

func UploadScreenshot(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(10 * 1024 * 1024) // This limits the maximum size to 10MB
	file, header, err := r.FormFile("file")
	if err != nil{
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	fmt.Println(file)
	fmt.Println(header)
	// defer file.Close()
	// // File name 
	// fmt.Println("[INFO] FILE NAME", header.Filename)

}