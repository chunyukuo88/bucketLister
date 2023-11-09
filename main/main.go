package main

import (
	listbuckets "bucketLister"
	"fmt"
)

func main() {
	buckets, err := listbuckets.ListBuckets()

	if err != nil {
		fmt.Println("Wasn't able to list the buckets.")
	}
	for _, name := range buckets {
		fmt.Println(*name)
	}
}
