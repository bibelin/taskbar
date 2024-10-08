package taskbar_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/bibelin/taskbar"
)

const desktop = "my.app.desktop"

func TestLibunity(t *testing.T) {
	os.Setenv("GO_TASKBAR_BACKEND", "libunity")
	if _, res := os.LookupEnv("GITHUB_ACTIONS"); res {
		os.Setenv("XDG_SESSION_TYPE", "wayland")
	}
	tb, err := taskbar.Connect(desktop, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer tb.Disconnect()

	if err = tb.SetProgress(50); err != nil {
		t.Fatal(err)
	}
	if err = tb.SetCount(1); err != nil {
		t.Fatal(err)
	}
	if err = tb.SetPulse(true); err != nil {
		t.Fatal(err)
	}
}

func TestXapp(t *testing.T) {
	if _, res := os.LookupEnv("GITHUB_ACTIONS"); res {
		t.Skip("Running in Github workflow, skipping the test.")
	}
	os.Setenv("GO_TASKBAR_BACKEND", "xapp")
	xid, err := strconv.ParseInt(os.Getenv("GO_TASKBAR_TEST_XID"), 10, 0)
	if err != nil {
		t.Fatalf("Error getting GO_TASKBAR_TEST_XID: %v\n You need to set it to valid xid to pass the test.", err)
	}

	tb, err := taskbar.Connect(desktop, int(xid))
	if err != nil {
		t.Fatal(err)
	}
	defer tb.Disconnect()

	if err = tb.SetProgress(50); err != nil {
		t.Fatal(err)
	}
	if err = tb.SetCount(1); err != nil {
		t.Fatal(err)
	}
	if err = tb.SetPulse(true); err != nil {
		t.Fatal(err)
	}
}

func TestBackendsFail(t *testing.T) {
	os.Setenv("GO_TASKBAR_BACKEND", "libunity")
	tb, err := taskbar.Connect("", 123)
	if err == nil {
		t.Fail()
	}
	errs := tb.Disconnect()
	if errs == nil {
		t.Fail()
	}

	os.Setenv("GO_TASKBAR_BACKEND", "xapp")
	tb, err = taskbar.Connect(desktop, 0)
	if err == nil {
		t.Fail()
	}
	errs = tb.Disconnect()
	if errs == nil {
		t.Fail()
	}
}
