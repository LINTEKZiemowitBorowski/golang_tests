 package main

 import (
	 "fmt"
	 "io"
	 "os"
	 "runtime"
	 "time"
	 "net/http"
	 "os/exec"
 )

 const (
	 SERVER_NAME = "http://developer.toradex.com/files/toradex-dev/uploads/media/Colibri/Linux/Images/"
	 FILE_NAME = "Apalis_T30_LinuxImageV2.5Beta2_20151106.tar.bz2"
 )

 func downloadFile1(rawURL string, fileName string) {
	 fmt.Printf("Downloading file...\n")

	 file, err := os.Create(fileName)
	 if err != nil {
		 fmt.Println(err)
		 panic(err)
	 }
	 defer file.Close()

	 check := http.Client{
		 CheckRedirect: func(r *http.Request, via []*http.Request) error {
			 r.URL.Opaque = r.URL.Path
			 return nil
		 },
	 }

	 resp, err := check.Get(rawURL + fileName)
	 if err != nil {
		 fmt.Println(err)
		 panic(err)
	 }
	 defer resp.Body.Close()
	 fmt.Printf("Response: %v\n", resp.Status)

	 size, err := io.Copy(file, resp.Body)
	 if err != nil {
		 panic(err)
	 }

	 fmt.Printf("%s with %v bytes downloaded\n", fileName, size)
 }

 func downloadFile2(rawURL string, fileName string) {
	 fmt.Printf("Downloading file...\n")

	 _, err := exec.Command("wget", rawURL + fileName).Output()
	 if err != nil {
		 fmt.Printf("%s\n", err)
	 }

	 fmt.Printf("File downloaded\n")
 }

 func removeFile(fileName string) {
	 if _, err := os.Stat(fileName); err == nil {
		 err := os.Remove(fileName)
		 if err != nil {
			 fmt.Println(err)
			 return
		 }
	 }
 }

 func main() {

	 runtime.GOMAXPROCS(runtime.NumCPU())
	 fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	 // Remove already downloaded file
	 removeFile(FILE_NAME)

	 startTime1 := time.Now()

	 // Legacy method
	 downloadFile1(SERVER_NAME, FILE_NAME)

	 legacyDownloadTime := time.Now()

	 // Remove already downloaded file
	 removeFile(FILE_NAME)

	 startTime2 := time.Now()

	 // WGET method
	 downloadFile2(SERVER_NAME, FILE_NAME)

	 wgetDownloadTime := time.Now()

	 fmt.Printf("Download time using legacy method: %f\n", legacyDownloadTime.Sub(startTime1).Seconds())
	 fmt.Printf("Download time using wget method: %f\n", wgetDownloadTime.Sub(startTime2).Seconds())
 }
