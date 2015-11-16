 package main

 import (
	 "fmt"
	 "io"
	 "os"
	 "runtime"
	 "time"
	 "net/http"
 )

 func download_file(rawURL string, fileName string) {
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

 func main() {
	 serverName := "http://developer.toradex.com/files/toradex-dev/uploads/media/Colibri/Linux/Images/"
	 fileName := "Apalis_T30_LinuxImageV2.5Beta2_20151106.tar.bz2"

	 runtime.GOMAXPROCS(runtime.NumCPU())
	 fmt.Printf("Number of available CPUs: %d\n", runtime.NumCPU())

	 // Remove already downloaded file
	 if _, err := os.Stat(fileName); err == nil {
		 err := os.Remove(fileName)
		 if err != nil {
			 fmt.Println(err)
			 return
		 }
	 }

	 start_time := time.Now()

	 download_file(serverName, fileName)

	 download_time := time.Now().Sub(start_time)
	 fmt.Printf("Download time: %f\n", download_time.Seconds())
 }
