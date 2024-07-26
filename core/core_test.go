package core

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
)

func TestParseAndPrintURL(t *testing.T) {
	testURL := "https://www.google.com/maps/place/Tsue+Chong+-+Retail+Store/@47.5979442,-122.3223302,0a,75y,90t/data=!3m7!1e2!3m5!1sAF1QipP9YKxn3p9QqdMxWHag6EN5jtmVnEpQ4eCGynZI!2e10!6shttps:%2F%2Flh5.googleusercontent.com%2Fp%2FAF1QipP9YKxn3p9QqdMxWHag6EN5jtmVnEpQ4eCGynZI%3Dw150-h150-k-no-p!7i3024!8i4032!4m8!3m7!1s0x54906abdca40b047:0x11ea097e0297d297!8m2!3d47.5979442!4d-122.3223302!9m1!1b1!16s%2Fg%2F1xgzbg_m?entry=ttu"

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ctx := logr.NewContext(context.Background(), testr.New(t))

	err := ParseAndPrintURL(ctx, testURL)
	if err != nil {
		t.Fatalf("ParseAndPrintURL failed: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("Failed to copy output: %v", err)
	}
	output := buf.String()

	expectedOutputs := []string{
		"Scheme: https",
		"Host: www.google.com",
		"Path: /maps/place/Tsue+Chong+-+Retail+Store/@47.5979442,-122.3223302,0a,75y,90t/data=!3m7!1e2!3m5!1sAF1QipP9YKxn3p9QqdMxWHag6EN5jtmVnEpQ4eCGynZI!2e10!6shttps://lh5.googleusercontent.com/p/AF1QipP9YKxn3p9QqdMxWHag6EN5jtmVnEpQ4eCGynZI=w150-h150-k-no-p!7i3024!8i4032!4m8!3m7!1s0x54906abdca40b047:0x11ea097e0297d297!8m2!3d47.5979442!4d-122.3223302!9m1!1b1!16s/g/1xgzbg_m",
		"Path Component 1: maps",
		"Path Component 2: place",
		"Path Component 3: Tsue+Chong+-+Retail+Store",
		"Path Component 4: @47.5979442,-122.3223302,0a,75y,90t",
		"Path Component 5: data=!3m7!1e2!3m5!1sAF1QipP9YKxn3p9QqdMxWHag6EN5jtmVnEpQ4eCGynZI!2e10!6shttps:",
		"Path Component 6: ",
		"Path Component 7: lh5.googleusercontent.com",
		"Path Component 8: p",
		"Path Component 9: AF1QipP9YKxn3p9QqdMxWHag6EN5jtmVnEpQ4eCGynZI=w150-h150-k-no-p!7i3024!8i4032!4m8!3m7!1s0x54906abdca40b047:0x11ea097e0297d297!8m2!3d47.5979442!4d-122.3223302!9m1!1b1!16s",
		"Path Component 10: g",
		"Path Component 11: 1xgzbg_m",
		"Query Param: entry = ttu",
	}

	for _, expected := range expectedOutputs {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain '%s', but it didn't.\nGot:\n%s", expected, output)
		}
	}
}
