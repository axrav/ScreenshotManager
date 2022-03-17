package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Uploading Screenshot [POST REQUEST]
func UploadScreenshot(w http.ResponseWriter, r *http.Request){

	// Checking Password

	if r.Header.Get("PASSWORD") != os.Getenv("PASSWORD"){
		json.NewEncoder(w).Encode("Your password is incorrect, Try again with valid password as header")
		return
	}
	r.ParseMultipartForm(10 * 1024 * 1024) // This limits the maximum size to 10MB
	
	//Making a seperate folder
	os.MkdirAll("screenshots", os.ModePerm)
	file, _, err := r.FormFile("file")
	if err != nil{
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	defer file.Close()
	saved, err2 := ioutil.TempFile("uploads", "*.jpg")
	if err2 != nil{
		fmt.Printf("[ERROR] %v\n", err2)
		return
	}
	filename := strings.Replace(saved.Name(), "uploads/" ,"", 1)
	fmt.Println("[INFO] A File has been uploaded", filename)
	defer saved.Close()


	// Checking The Host
	var myhost string
	if os.Getenv("HOST") == ""{
		myhost = "https://localhost:"
	} else {
		myhost = os.Getenv("HOST")
	}
	
	// Reading the bytes of the screen
	finalfile, err3 := ioutil.ReadAll(file)
	if err3 != nil{
 	fmt.Printf("[ERROR] %v\n", err2)
		return
	}
	// writing bytes to the file
	saved.Write(finalfile)
	imgUrl := fmt.Sprintf("%v:%v/screen/%v", myhost, ThePort(), filename)
	
	// returning the path link to User
	json.NewEncoder(w).Encode(imgUrl)

}
// Getting The Screenshot

func SendScreenshot(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "image/jpeg")
	imagePath := params["path"]
	finalImage := fmt.Sprintf("uploads/%v", imagePath)

	// Reading the image bytes
	imagebytes, err := ioutil.ReadFile(finalImage)
	if err != nil{
		fmt.Printf("[ERROR] %v\n", err)
	}
	// writing the image bytes
	w.Write(imagebytes)
}

// Checking Port entered by user
func ThePort() string {
	var port string
	if os.Getenv("DEFAULT_PORT") == ""{
		port = "8080"
	} else {
		port= os.Getenv("DEFAULT_PORT")
	}
	return port
}
