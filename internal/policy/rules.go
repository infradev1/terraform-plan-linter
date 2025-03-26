package policy

type Violation struct {
	Resource string
	Message  string
}

func CheckForPublicS3(resources map[string]interface{}) []Violation {
	var violations []Violation
	// TODO: walk through resources["aws_s3_bucket"] and check acl/public_access_block
	return violations
}
