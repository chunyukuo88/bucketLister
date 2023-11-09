package listbuckets

//begin import
import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client

func init() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("Configuration error! It is: " + err.Error())
	}
	client = s3.NewFromConfig(config)
}

func ListBuckets() ([]*string, error) {
	bucketArray := make([]*string, 0, 10)
	params := &s3.ListBucketsInput{}
	res, err := client.ListBuckets(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	for _, bucket := range res.Buckets {
		bucketArray = append(bucketArray, bucket.Name)
	}

	return bucketArray, nil
}
