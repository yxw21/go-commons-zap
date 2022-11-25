package go_common_aws_sqs

import (
	"fmt"
	"os"
	"testing"
)

func TestZincWriter(t *testing.T) {
	zincWriterInstance := NewZincWriter("test")

	// DocumentIndex
	resp, r, err := zincWriterInstance.DocumentIndex(map[string]interface{}{
		"name": "tom",
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error when calling `Document.Index``: %v\n", err)
		_, _ = fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		t.Error(err)
	}
	_, _ = fmt.Fprintf(os.Stdout, "Response from `Document.Index`: %v\n", resp.GetId())

	// DocumentBulkV2
	resp1, r1, err1 := zincWriterInstance.DocumentBulkV2([]map[string]interface{}{
		{
			"name": "jerry",
		},
	})
	if err1 != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error when calling `Document.Bulkv2``: %v\n", err1)
		_, _ = fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r1)
		t.Error(err1)
	}
	_, _ = fmt.Fprintf(os.Stdout, "Response from `Document.Bulkv2`: %v\n", resp1.GetRecordCount())
}

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.Info("test logger")
}

func TestLoggerWithZinc(t *testing.T) {
	logger := NewLoggerWithZinc(NewZincWriter("test"))
	logger.Debug("debug message")
	logger.Error("error message")
}
