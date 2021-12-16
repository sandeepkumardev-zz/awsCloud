package utils

func CreateBucketName(username string, userId string) string {
	return username + "-buckets-" + userId
}
