package gomlib_test

import (
	"testing"

	"github.com/prife/gomlib"
)

func TestDevice(t *testing.T) {
	t.Log(gomlib.Android.String())
	t.Log(gomlib.OSAndroid.String())
}
