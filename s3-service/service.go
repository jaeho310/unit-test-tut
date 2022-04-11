package s3_service

import (
	"context"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client S3Client

type S3Client interface {
	CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error)
}

func Init() {
	log.Println("load config")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("your-profile"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("create s3 client")
	SetS3Client(s3.NewFromConfig(cfg))
}

func SetS3Client(client S3Client) {
	s3Client = client
}

func CopyS3Object() error {
	sourceBucket := url.PathEscape("my-foo-test2")                              // 복사한 오브젝트가 들어갈 버킷
	destinationBucket := url.PathEscape("my-foo-test1" + "/" + "my-object.txt") // 복사할 원본의 위치. 버킷/디렉토리/파일.확장자 풀네임으로 입력
	objectName := "my-copied-object"                                            // 복사한 오브젝트의 새로운 이름

	input := &s3.CopyObjectInput{
		Bucket:     &sourceBucket,
		CopySource: &destinationBucket,
		Key:        &objectName,
		// StorageClass: , copyobject는 storage타입 등 go-sdk에서 직접 제공하지 않는 설정을 바꿀때 사용할 수 있습니다.
	}
	_, err := s3Client.CopyObject(context.TODO(), input)
	if err != nil {
		return err
	}
	log.Println("Copied " + objectName + " from " + sourceBucket + " to " + destinationBucket)
	return nil
}
