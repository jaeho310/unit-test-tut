package main

import s3Service "unit-test-tut/s3-service"

func main() {
	err := s3Service.CopyS3Object()
	if err != nil {
		panic(err)
	}
}
