// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudHSM_AddTagsToResource() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.AddTagsToResourceInput{
		ResourceArn: aws.String("String"), // Required
		TagList: []*cloudhsm.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_CreateHapg() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.CreateHapgInput{
		Label: aws.String("Label"), // Required
	}
	resp, err := svc.CreateHapg(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_CreateHsm() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.CreateHsmInput{
		IamRoleArn:       aws.String("IamRoleArn"),       // Required
		SshKey:           aws.String("SshKey"),           // Required
		SubnetId:         aws.String("SubnetId"),         // Required
		SubscriptionType: aws.String("SubscriptionType"), // Required
		ClientToken:      aws.String("ClientToken"),
		EniIp:            aws.String("IpAddress"),
		ExternalId:       aws.String("ExternalId"),
		SyslogIp:         aws.String("IpAddress"),
	}
	resp, err := svc.CreateHsm(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_CreateLunaClient() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.CreateLunaClientInput{
		Certificate: aws.String("Certificate"), // Required
		Label:       aws.String("ClientLabel"),
	}
	resp, err := svc.CreateLunaClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_DeleteHapg() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.DeleteHapgInput{
		HapgArn: aws.String("HapgArn"), // Required
	}
	resp, err := svc.DeleteHapg(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_DeleteHsm() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.DeleteHsmInput{
		HsmArn: aws.String("HsmArn"), // Required
	}
	resp, err := svc.DeleteHsm(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_DeleteLunaClient() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.DeleteLunaClientInput{
		ClientArn: aws.String("ClientArn"), // Required
	}
	resp, err := svc.DeleteLunaClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_DescribeHapg() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.DescribeHapgInput{
		HapgArn: aws.String("HapgArn"), // Required
	}
	resp, err := svc.DescribeHapg(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_DescribeHsm() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.DescribeHsmInput{
		HsmArn:          aws.String("HsmArn"),
		HsmSerialNumber: aws.String("HsmSerialNumber"),
	}
	resp, err := svc.DescribeHsm(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_DescribeLunaClient() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.DescribeLunaClientInput{
		CertificateFingerprint: aws.String("CertificateFingerprint"),
		ClientArn:              aws.String("ClientArn"),
	}
	resp, err := svc.DescribeLunaClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_GetConfig() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.GetConfigInput{
		ClientArn:     aws.String("ClientArn"),     // Required
		ClientVersion: aws.String("ClientVersion"), // Required
		HapgList: []*string{ // Required
			aws.String("HapgArn"), // Required
			// More values...
		},
	}
	resp, err := svc.GetConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ListAvailableZones() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	var params *cloudhsm.ListAvailableZonesInput
	resp, err := svc.ListAvailableZones(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ListHapgs() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.ListHapgsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListHapgs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ListHsms() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.ListHsmsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListHsms(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ListLunaClients() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.ListLunaClientsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListLunaClients(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ListTagsForResource() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.ListTagsForResourceInput{
		ResourceArn: aws.String("String"), // Required
	}
	resp, err := svc.ListTagsForResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ModifyHapg() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.ModifyHapgInput{
		HapgArn: aws.String("HapgArn"), // Required
		Label:   aws.String("Label"),
		PartitionSerialList: []*string{
			aws.String("PartitionSerial"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyHapg(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ModifyHsm() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.ModifyHsmInput{
		HsmArn:     aws.String("HsmArn"), // Required
		EniIp:      aws.String("IpAddress"),
		ExternalId: aws.String("ExternalId"),
		IamRoleArn: aws.String("IamRoleArn"),
		SubnetId:   aws.String("SubnetId"),
		SyslogIp:   aws.String("IpAddress"),
	}
	resp, err := svc.ModifyHsm(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_ModifyLunaClient() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.ModifyLunaClientInput{
		Certificate: aws.String("Certificate"), // Required
		ClientArn:   aws.String("ClientArn"),   // Required
	}
	resp, err := svc.ModifyLunaClient(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudHSM_RemoveTagsFromResource() {
	sess := session.Must(session.NewSession())

	svc := cloudhsm.New(sess)

	params := &cloudhsm.RemoveTagsFromResourceInput{
		ResourceArn: aws.String("String"), // Required
		TagKeyList: []*string{ // Required
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
