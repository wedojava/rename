package main

import "testing"

func TestMatch(t *testing.T) {
	tcs := []struct {
		src, res string
	}{
		{"-01-2020-08-17-16-08-02.dat-UDP-000015-asdf1234.@asd123.@asdadf123.123.123.123@asdfasdf@C2DMXsafd.tif.bz2", "[-01-2020-08-17-16-08-02.dat][UDP-000015]@C2DMXsafd.tif.bz2"},
	}
	for _, tc := range tcs {
		got := match(tc.src)
		if tc.res != got {
			t.Errorf("\nwant: %v\ngot: %v\n", tc.res, got)
		}
	}
}
