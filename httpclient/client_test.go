package httpclient

import "testing"

func TestDoHTTP(t *testing.T) {
	resp, err := DoHTTP(Request{
		TimeoutSec:30,
		Log:true,
	}, "https://a-very-bad-domain-name-nonnnnnnbreeiunsdvc.rich")
	if err == nil {
		t.Fatal("did not error")
	}
	if resp.Non2xxToError() == nil {
		t.Fatal("did not error")
	}

	resp, err = DoHTTP(Request{
		TimeoutSec:30,
		Log:true,
	}, "https://github.com")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Non2xxToError() != nil {
		t.Fatal(err)
	}
}
