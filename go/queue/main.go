// ref:
// https://medium.com/smsjunk/handling-1-million-requests-per-minute-with-golang-f70ac505fcaa

package main

func main() {

}

func StartProcessor() {
	for _, payload := range content.Payloads {
		Queue <- payload
	}

	for {
		select {
		case job := <-Queue:
			job.payload.UploadToS3() // <-- STILL NOT GOOD
		}
	}
}
