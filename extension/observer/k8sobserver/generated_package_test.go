// Code generated by mdatagen. DO NOT EDIT.

package k8sobserver

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("k8s.io/apimachinery/pkg/watch.(*Broadcaster).loop"))
}