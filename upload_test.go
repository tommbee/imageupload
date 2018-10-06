package imageupload

import (
	"bytes"
	"strings"
	"testing"
)

func init() {
	initExtMap()
	fs = dummyFS{}
}

func TestUnknownFormat(t *testing.T) {
	_, err := saveFile(strings.NewReader("nop"), "/", "testID", "unknown", 0)
	if err != ErrFileNotSupported {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestJPGDecodeFail(t *testing.T) {
	_, err := saveFile(strings.NewReader("nop"), "/", "testID", "jpg", 0)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPNGDecodeFail(t *testing.T) {
	_, err := saveFile(strings.NewReader("nop"), "/", "testID", "png", 0)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGIFDecodeFail(t *testing.T) {
	_, err := saveFile(strings.NewReader("nop"), "/", "testID", "gif", 0)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestJPGHappyPath(t *testing.T) {
	p, err := saveFile(bytes.NewReader(testJPGImage), "/", "testID", "jpg", 0)
	if err != nil {
		t.Fatal(err)
	}
	if p != "/testID.jpg" {
		t.Errorf("invalid file name, expected: testID.jpg, got: %s", p)
	}
}

func TestGetExt(t *testing.T) {
	testTable := []struct{
		Input string
		Output string
	}{
		{"img.jpg", "jpg"},
		{"imgjpg", "imgjpg"},
		{"img.gif.jpg", "jpg"},
	}

	for _, tt := range testTable {
		if tt.Output != getExt(tt.Input) {
			t.Errorf("unexpected output, expeceted: %s, got: %s", tt.Output, getExt(tt.Input))
		}
	}
}

type dummyFile struct{}

func (dummyFile) Close() error { return nil }
func (dummyFile) Write(p []byte) (n int, err error) { return len(p), nil}

type dummyFS struct{
	createFile string
}

func (dummyFS) Create(name string) (file, error) { return dummyFile{}, nil }

var testJPGImage = []byte{
	0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10, 0x4a, 0x46, 0x49, 0x46, 0x00, 0x01, 0x01, 0x01, 0x00, 0x60,
	0x00, 0x60, 0x00, 0x00, 0xff, 0xdb, 0x00, 0x43, 0x00, 0x06, 0x04, 0x05, 0x06, 0x05, 0x04, 0x06,
	0x06, 0x05, 0x06, 0x07, 0x07, 0x06, 0x08, 0x0a, 0x10, 0x0a, 0x0a, 0x09, 0x09, 0x0a, 0x14, 0x0e,
	0x0f, 0x0c, 0x10, 0x17, 0x14, 0x18, 0x18, 0x17, 0x14, 0x16, 0x16, 0x1a, 0x1d, 0x25, 0x1f, 0x1a,
	0x1b, 0x23, 0x1c, 0x16, 0x16, 0x20, 0x2c, 0x20, 0x23, 0x26, 0x27, 0x29, 0x2a, 0x29, 0x19, 0x1f,
	0x2d, 0x30, 0x2d, 0x28, 0x30, 0x25, 0x28, 0x29, 0x28, 0xff, 0xdb, 0x00, 0x43, 0x01, 0x07, 0x07,
	0x07, 0x0a, 0x08, 0x0a, 0x13, 0x0a, 0x0a, 0x13, 0x28, 0x1a, 0x16, 0x1a, 0x28, 0x28, 0x28, 0x28,
	0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
	0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0xff, 0xc0,
	0x00, 0x11, 0x08, 0x00, 0x77, 0x01, 0x40, 0x03, 0x01, 0x22, 0x00, 0x02, 0x11, 0x01, 0x03, 0x11,
	0x01, 0xff, 0xc4, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x02, 0x02, 0x03, 0x01, 0x01, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x08, 0x01, 0x06, 0x03, 0x05, 0x09, 0x02, 0x04,
	0xff, 0xc4, 0x00, 0x44, 0x10, 0x00, 0x01, 0x03, 0x02, 0x02, 0x05, 0x0a, 0x01, 0x07, 0x0b, 0x04,
	0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x11, 0x07, 0x08, 0x12,
	0x21, 0x31, 0x13, 0x18, 0x22, 0x41, 0x51, 0x56, 0x61, 0x71, 0x81, 0x93, 0xb1, 0x14, 0x32, 0x62,
	0x82, 0x91, 0xa1, 0xd2, 0x15, 0x17, 0x23, 0x33, 0x42, 0x52, 0x55, 0x72, 0x94, 0xa2, 0xc1, 0x16,
	0x73, 0x92, 0xb2, 0x25, 0xc2, 0xe1, 0xff, 0xc4, 0x00, 0x16, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0xff, 0xc4,
	0x00, 0x17, 0x11, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x11, 0x01, 0x21, 0xff, 0xda, 0x00, 0x0c, 0x03, 0x01, 0x00, 0x02, 0x11,
	0x03, 0x11, 0x00, 0x3f, 0x00, 0xb2, 0x98, 0x8e, 0xba, 0xe5, 0x6f, 0xa1, 0x59, 0xed, 0x56, 0x97,
	0x5d, 0x64, 0x6f, 0xce, 0x81, 0x95, 0x0d, 0x89, 0xea, 0x9f, 0x47, 0x6b, 0x72, 0xaf, 0x9a, 0xa1,
	0x0d, 0xde, 0xb5, 0x89, 0x8a, 0xc7, 0x5e, 0xea, 0x2b, 0xc6, 0x0d, 0xbd, 0xd0, 0xd5, 0xb7, 0x8c,
	0x55, 0x0f, 0x8d, 0x8e, 0x5f, 0x14, 0xcf, 0x8a, 0x78, 0xa6, 0xe2, 0x79, 0x3a, 0xac, 0x47, 0x87,
	0x6d, 0x18, 0x96, 0xde, 0xea, 0x2b, 0xed, 0xba, 0x9a, 0xba, 0x99, 0x7f, 0x62, 0x66, 0x67, 0xb2,
	0xbd, 0xad, 0x5e, 0x2d, 0x5f, 0x14, 0x54, 0x50, 0x20, 0xde, 0x74, 0x16, 0x9e, 0xed, 0xdc, 0xbd,
	0xf8, 0x87, 0x3a, 0x0b, 0x4f, 0x76, 0xee, 0x5e, 0xfc, 0x47, 0x4b, 0xa4, 0x8d, 0x5c, 0x26, 0x81,
	0xb2, 0xd7, 0x60, 0x5a, 0x97, 0x4e, 0xc4, 0xe9, 0x2d, 0xba, 0xa9, 0xfd, 0x3f, 0x28, 0xe4, 0x5e,
	0x3e, 0x4e, 0xfb, 0x4a, 0xf1, 0x70, 0xa2, 0xaa, 0xb6, 0xd6, 0xcd, 0x47, 0x70, 0xa6, 0x9a, 0x96,
	0xae, 0x17, 0x6c, 0xc9, 0x0c, 0xcc, 0x56, 0x3d, 0x8b, 0xd8, 0xa8, 0xa4, 0xeb, 0x53, 0x16, 0x87,
	0x9d, 0x05, 0xa7, 0xbb, 0x77, 0x2f, 0x7e, 0x21, 0xce, 0x82, 0xd3, 0xdd, 0xbb, 0x97, 0xbf, 0x11,
	0x55, 0x81, 0x29, 0x31, 0x6a, 0x79, 0xd0, 0x5a, 0x7b, 0xb7, 0x72, 0xf7, 0xe2, 0x1c, 0xe8, 0x2d,
	0x3d, 0xdb, 0xb9, 0x7b, 0xf1, 0x15, 0x58, 0x0a, 0x4c, 0x5a, 0x9e, 0x74, 0x16, 0x9e, 0xed, 0xdc,
	0xbd, 0xf8, 0x87, 0x3a, 0x0b, 0x4f, 0x76, 0xee, 0x5e, 0xfc, 0x45, 0x56, 0x02, 0x93, 0x16, 0xa7,
	0x9d, 0x05, 0xa7, 0xbb, 0x77, 0x2f, 0x7e, 0x21, 0xce, 0x82, 0xd3, 0xdd, 0xbb, 0x97, 0xbf, 0x11,
	0x55, 0x80, 0xa4, 0xc5, 0xa9, 0xe7, 0x41, 0x69, 0xee, 0xdd, 0xcb, 0xdf, 0x88, 0x90, 0x74, 0x49,
	0xa5, 0x4a, 0x1d, 0x23, 0xc9, 0x73, 0x8e, 0x92, 0xdf, 0x51, 0x43, 0x25, 0x0a, 0x46, 0xe5, 0x6c,
	0xf2, 0x35, 0xca, 0xf4, 0x7e, 0xd6, 0xf4, 0xd9, 0xec, 0xd9, 0xfb, 0xca, 0x2a, 0x4a, 0xda, 0xb4,
	0x62, 0x36, 0x58, 0x34, 0xa3, 0x49, 0x0c, 0xef, 0x46, 0x53, 0x5d, 0x22, 0x75, 0x0b, 0x95, 0x78,
	0x23, 0xd5, 0x51, 0xd1, 0xff, 0x00, 0x73, 0x72, 0xfa, 0xc5, 0xa6, 0xe2, 0xed, 0x00, 0x0a, 0xc8,
	0x00, 0x00, 0x08, 0xff, 0x00, 0x4e, 0x38, 0x96, 0x5c, 0x29, 0x81, 0xbf, 0x29, 0xd3, 0xb9, 0x5a,
	0xf6, 0x57, 0xd2, 0x22, 0xaa, 0x2e, 0x59, 0xb7, 0x97, 0x6b, 0x9e, 0x9e, 0x4a, 0xd6, 0xb9, 0x17,
	0xc1, 0x4d, 0xfa, 0x29, 0x19, 0x2c, 0x4c, 0x92, 0x37, 0x23, 0xd8, 0xf4, 0x47, 0x35, 0xc9, 0xc1,
	0x51, 0x78, 0x28, 0x1f, 0x40, 0xc3, 0xdc, 0x8d, 0x6a, 0xb9, 0xca, 0x88, 0xd4, 0xde, 0xaa, 0xbc,
	0x11, 0x08, 0xfb, 0x41, 0xf8, 0x9e, 0x5c, 0x5b, 0x84, 0xab, 0x6e, 0x53, 0x39, 0xce, 0xff, 0x00,
	0xca, 0xd6, 0x31, 0x99, 0xae, 0x79, 0x47, 0xca, 0xab, 0x98, 0x9e, 0x8d, 0x72, 0x27, 0xa0, 0x12,
	0x10, 0x00, 0x00, 0x00, 0x08, 0xf3, 0x4b, 0x7a, 0x51, 0xa1, 0xd1, 0xba, 0x5a, 0xd2, 0xb2, 0xdf,
	0x51, 0x5c, 0xfa, 0xe5, 0x93, 0x65, 0xb0, 0x3d, 0xad, 0xd8, 0x46, 0x6c, 0xe6, 0xab, 0xb5, 0xfc,
	0xc8, 0x47, 0x7c, 0xe8, 0x2d, 0x3d, 0xdb, 0xb9, 0x7b, 0xf1, 0x11, 0x96, 0xb3, 0xd8, 0x8d, 0x97,
	0xdd, 0x27, 0x4d, 0x4b, 0x4e, 0xf4, 0x75, 0x3d, 0xa6, 0x14, 0xa3, 0x45, 0x4e, 0x0b, 0x26, 0x7b,
	0x52, 0x7d, 0x8a, 0xa8, 0xdf, 0xaa, 0x44, 0x64, 0xad, 0x66, 0x2d, 0x4f, 0x3a, 0x0b, 0x4f, 0x76,
	0xee, 0x5e, 0xfc, 0x43, 0x9d, 0x05, 0xa7, 0xbb, 0x77, 0x2f, 0x7e, 0x22, 0xab, 0x02, 0x52, 0x62,
	0xd4, 0xf3, 0xa0, 0xb4, 0xf7, 0x6e, 0xe5, 0xef, 0xc4, 0x39, 0xd0, 0x5a, 0x7b, 0xb7, 0x72, 0xf7,
	0xe2, 0x2a, 0xb0, 0x14, 0x98, 0xb5, 0x3c, 0xe8, 0x2d, 0x3d, 0xdb, 0xb9, 0x7b, 0xf1, 0x0e, 0x74,
	0x16, 0x9e, 0xed, 0xdc, 0xbd, 0xf8, 0x8a, 0xac, 0x05, 0x26, 0x2d, 0x4f, 0x3a, 0x0b, 0x4f, 0x76,
	0xee, 0x5e, 0xfc, 0x43, 0x9d, 0x05, 0xa7, 0xbb, 0x77, 0x2f, 0x7e, 0x22, 0xab, 0x01, 0x49, 0x8b,
	0x53, 0xce, 0x82, 0xd3, 0xdd, 0xbb, 0x97, 0xbf, 0x10, 0xe7, 0x41, 0x69, 0xee, 0xdd, 0xcb, 0xdf,
	0x88, 0xab, 0xb4, 0x34, 0x75, 0x35, 0xf5, 0x90, 0xd2, 0x50, 0xd3, 0xcb, 0x53, 0x55, 0x33, 0xb6,
	0x63, 0x86, 0x16, 0x2b, 0xde, 0xf5, 0xec, 0x44, 0x4d, 0xea, 0x58, 0x5d, 0x1b, 0xea, 0xe3, 0x53,
	0x54, 0xd8, 0xab, 0xb1, 0xcd, 0x4b, 0xa9, 0x62, 0x5e, 0x92, 0x5b, 0xa9, 0x9e, 0x9c, 0xa2, 0xf8,
	0x49, 0x27, 0x06, 0xf9, 0x37, 0x35, 0xf1, 0x42, 0xf4, 0x98, 0xda, 0xac, 0xfa, 0xc5, 0xc1, 0x7a,
	0xae, 0x65, 0x15, 0xa3, 0x07, 0x5e, 0xeb, 0x6a, 0xdd, 0xc2, 0x2a, 0x77, 0xb1, 0xee, 0xf3, 0x5c,
	0xb8, 0x27, 0x8a, 0xee, 0x26, 0x3c, 0x37, 0x70, 0xb9, 0xdc, 0xa8, 0xb9, 0x7b, 0xb5, 0x9d, 0xf6,
	0x97, 0xbb, 0x25, 0x6c, 0x12, 0x54, 0x32, 0x57, 0xe5, 0xf4, 0xb6, 0x37, 0x27, 0xda, 0xa6, 0x70,
	0xd6, 0x1b, 0xb3, 0xe1, 0x9b, 0x7b, 0x68, 0xac, 0x36, 0xea, 0x6a, 0x1a, 0x64, 0xe2, 0xd8, 0x59,
	0x92, 0xbb, 0xc5, 0xcb, 0xc5, 0xcb, 0xe2, 0xaa, 0xaa, 0x76, 0xc5, 0x64, 0x00, 0x00, 0x34, 0xad,
	0x25, 0x68, 0xda, 0xc3, 0x8f, 0xed, 0xfc, 0x9d, 0xd2, 0x0e, 0x4a, 0xba, 0x36, 0xaa, 0x53, 0xd7,
	0x42, 0x88, 0x92, 0xc4, 0xbd, 0x99, 0xfe, 0xd3, 0x7e, 0x8a, 0xee, 0xf2, 0x5d, 0xe6, 0xea, 0x00,
	0xf3, 0xef, 0x48, 0xb8, 0x0e, 0xf3, 0x80, 0xaf, 0x5f, 0x21, 0xbc, 0x44, 0x8e, 0x8a, 0x4c, 0xd6,
	0x9a, 0xae, 0x34, 0x5e, 0x4a, 0xa1, 0xa9, 0xd6, 0x9d, 0x8a, 0x9d, 0x6d, 0x5d, 0xe9, 0xe2, 0x9b,
	0xcd, 0x4c, 0xf4, 0x53, 0x1a, 0x61, 0x6b, 0x5e, 0x30, 0xc3, 0xf5, 0x16, 0x8b, 0xd4, 0x09, 0x2d,
	0x34, 0xa9, 0x9b, 0x5c, 0x9b, 0x9f, 0x13, 0xd3, 0x83, 0xd8, 0xbd, 0x4e, 0x4f, 0xfe, 0x2e, 0xe5,
	0x28, 0x96, 0x91, 0x30, 0x75, 0xc3, 0x03, 0x62, 0x8a, 0x8b, 0x3d, 0xcd, 0x36, 0xb6, 0x7f, 0x49,
	0x04, 0xe8, 0x99, 0x36, 0x78, 0x95, 0x7a, 0x2f, 0x4f, 0x82, 0xa7, 0x52, 0xa2, 0xa1, 0x9d, 0xc6,
	0xb3, 0x6b, 0x58, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xb8, 0xa4, 0x7c, 0x52, 0x32, 0x48,
	0x9e, 0xe6, 0x48, 0xc7, 0x23, 0x9a, 0xf6, 0xae, 0x4a, 0xd5, 0x45, 0xcd, 0x15, 0x17, 0xb5, 0x14,
	0xf8, 0x00, 0x5e, 0x0d, 0x08, 0x69, 0x42, 0x8f, 0x1d, 0xd8, 0xe3, 0xa6, 0xac, 0x96, 0x38, 0xb1,
	0x15, 0x2b, 0x11, 0x2a, 0x60, 0x55, 0xcb, 0x95, 0x44, 0xdd, 0xca, 0xb1, 0x3a, 0xd1, 0x7a, 0xd1,
	0x38, 0x2f, 0x86, 0x44, 0xa0, 0x79, 0xab, 0x47, 0x55, 0x51, 0x45, 0x55, 0x15, 0x55, 0x14, 0xf2,
	0xd3, 0xd4, 0xc4, 0xed, 0xb8, 0xe5, 0x89, 0xea, 0xc7, 0xb1, 0x7b, 0x51, 0x53, 0x7a, 0x29, 0x30,
	0x61, 0xad, 0x62, 0x71, 0x8d, 0xaa, 0x16, 0xc3, 0x71, 0x6d, 0x0d, 0xe2, 0x36, 0xa6, 0x48, 0xfa,
	0x86, 0x2c, 0x72, 0xfa, 0xb9, 0x9b, 0x97, 0xd5, 0x0b, 0x53, 0x71, 0x72, 0x81, 0x5d, 0x2d, 0x7a,
	0xd0, 0x5b, 0x5e, 0xad, 0x4b, 0xae, 0x1b, 0xad, 0x81, 0x32, 0xe9, 0x3a, 0x9a, 0xa1, 0x92, 0xef,
	0xf2, 0x72, 0x34, 0xd8, 0xa3, 0xd6, 0x47, 0x04, 0x39, 0x3a, 0x51, 0x5e, 0x59, 0xfc, 0xd4, 0x88,
	0xbf, 0x07, 0x28, 0xa9, 0x35, 0xf3, 0xad, 0xbb, 0xd1, 0xba, 0x2d, 0x81, 0x8a, 0xa9, 0xfa, 0x4b,
	0x94, 0x2d, 0xcb, 0xb5, 0x36, 0x5e, 0xbf, 0xe0, 0x8d, 0xac, 0x1a, 0xc8, 0x5c, 0x6d, 0x18, 0x72,
	0xd9, 0x6c, 0xfc, 0x81, 0x05, 0x54, 0xd4, 0x74, 0xd1, 0xc0, 0xea, 0x89, 0x6a, 0xd5, 0xbc, 0xa2,
	0xb5, 0x32, 0xda, 0xd9, 0x46, 0x6e, 0xdc, 0x89, 0xd6, 0x75, 0x5a, 0x7f, 0xd2, 0xc5, 0xbb, 0x48,
	0x14, 0xb6, 0xbb, 0x7d, 0x8a, 0x9a, 0xb2, 0x2a, 0x3a, 0x59, 0x5d, 0x51, 0x24, 0x95, 0x2d, 0x46,
	0x2c, 0x8f, 0x56, 0xec, 0xa2, 0x23, 0x51, 0x57, 0x72, 0x22, 0xae, 0xf5, 0x5e, 0xbe, 0x04, 0x30,
	0x45, 0xcc, 0x4f, 0xd7, 0x9d, 0x64, 0xee, 0x37, 0x4b, 0x05, 0xc6, 0xdf, 0xfe, 0x9f, 0xa7, 0xa6,
	0x9a, 0xaa, 0x9e, 0x48, 0x19, 0x51, 0x1d, 0x5b, 0x9d, 0xc9, 0xab, 0x9b, 0x96, 0xd6, 0xca, 0xb7,
	0x7e, 0x59, 0xaf, 0x59, 0x21, 0x6a, 0x87, 0x2b, 0x1d, 0xa3, 0x2a, 0xc8, 0x5a, 0xbd, 0x28, 0xae,
	0x72, 0xa2, 0xa7, 0x66, 0x6c, 0x8d, 0x50, 0xa7, 0xe4, 0xc7, 0xab, 0xfe, 0x95, 0x6d, 0xda, 0x3e,
	0x65, 0xd6, 0x8e, 0xf7, 0x4d, 0x57, 0x2d, 0x15, 0x63, 0xd9, 0x33, 0x24, 0xa6, 0x44, 0x7a, 0xc6,
	0xf6, 0xb7, 0x65, 0x51, 0x5a, 0xaa, 0x9b, 0x95, 0x32, 0xde, 0x8b, 0xd4, 0x0d, 0xc5, 0xcd, 0x04,
	0x2c, 0xfd, 0x64, 0x70, 0x43, 0x53, 0xa3, 0x15, 0xe5, 0xfe, 0x0d, 0xa4, 0x44, 0xf8, 0xb8, 0xd7,
	0xae, 0x9a, 0xd0, 0x5b, 0x23, 0x73, 0x92, 0xd5, 0x86, 0xeb, 0xa7, 0x4c, 0xba, 0x2e, 0xa9, 0xa8,
	0x64, 0x5b, 0xfc, 0x9b, 0xb4, 0x5a, 0x91, 0x62, 0x88, 0xd3, 0x4d, 0x5a, 0x4f, 0xa2, 0xc0, 0x36,
	0x39, 0x22, 0xa7, 0x96, 0x39, 0xb1, 0x05, 0x4b, 0x15, 0x29, 0x29, 0xb3, 0xcd, 0x59, 0x9e, 0xee,
	0x55, 0xe9, 0xd4, 0xd4, 0xfb, 0xd7, 0x72, 0x75, 0xe5, 0x01, 0x62, 0x4d, 0x62, 0xb1, 0x8d, 0xd2,
	0x17, 0x43, 0x6d, 0x8e, 0x82, 0xce, 0xc7, 0x22, 0xa2, 0xbe, 0x06, 0x2c, 0xb2, 0xfa, 0x39, 0xfb,
	0x93, 0xfe, 0x24, 0x3d, 0x5d, 0x59, 0x53, 0x5f, 0x59, 0x2d, 0x5d, 0x75, 0x44, 0xd5, 0x35, 0x53,
	0x3b, 0x6a, 0x49, 0xa6, 0x7a, 0xbd, 0xef, 0x5e, 0xd5, 0x55, 0xde, 0xa2, 0xae, 0x63, 0xe2, 0x79,
	0x64, 0xa8, 0x9e, 0x49, 0xa7, 0x91, 0xd2, 0x4d, 0x23, 0x95, 0xef, 0x7b, 0x97, 0x35, 0x73, 0x95,
	0x73, 0x55, 0x5f, 0x15, 0x55, 0xcc, 0xe3, 0x00, 0x8a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1b,
	0x4e, 0x8f, 0x70, 0x35, 0xe7, 0x1d, 0xde, 0x92, 0xdf, 0x65, 0x85, 0x36, 0x59, 0x93, 0xaa, 0x2a,
	0x64, 0x45, 0xe4, 0xa9, 0xda, 0xbd, 0x6e, 0x5e, 0xde, 0xc6, 0xa6, 0xf5, 0xfb, 0x54, 0xe1, 0xd1,
	0xfe, 0x11, 0xb8, 0x63, 0x7c, 0x4f, 0x4b, 0x66, 0xb5, 0xb7, 0x27, 0xc9, 0xd3, 0x9a, 0x65, 0x4c,
	0xdb, 0x04, 0x49, 0xf3, 0x9e, 0xef, 0x2e, 0xa4, 0xeb, 0x55, 0x44, 0x2f, 0x76, 0x07, 0xc2, 0x76,
	0xbc, 0x19, 0x87, 0xa9, 0xed, 0x16, 0x68, 0x76, 0x21, 0x8f, 0xa4, 0xf9, 0x1d, 0xbd, 0xf3, 0x3d,
	0x78, 0xbd, 0xeb, 0xd6, 0xab, 0xf7, 0x70, 0x4d, 0xc8, 0x33, 0x13, 0x75, 0xd3, 0x68, 0xcf, 0x46,
	0x76, 0x1c, 0x01, 0x43, 0xb3, 0x6d, 0x87, 0x97, 0xb8, 0xc8, 0xd4, 0x49, 0xeb, 0xe6, 0x44, 0xe5,
	0x64, 0xf0, 0x4f, 0xdd, 0x6f, 0xd1, 0x4f, 0x5c, 0xd7, 0x79, 0xbc, 0x00, 0x69, 0x90, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x8b, 0xb5, 0x84, 0xc0, 0xac, 0xc6, 0x38, 0x1a, 0x79, 0xa9, 0x62, 0x45,
	0xbb, 0xdb, 0x1a, 0xea, 0x9a, 0x55, 0x4e, 0x2f, 0x44, 0x4e, 0x9c, 0x7f, 0x59, 0x13, 0x77, 0x8a,
	0x21, 0x28, 0x85, 0xe0, 0x07, 0x99, 0xe8, 0xb9, 0xa2, 0x2a, 0x70, 0x50, 0x6d, 0xda, 0x59, 0xb0,
	0xb7, 0x0d, 0x69, 0x1f, 0x10, 0x5a, 0xe2, 0x6a, 0x32, 0x08, 0xea, 0x96, 0x48, 0x5a, 0x9d, 0x51,
	0xc8, 0x88, 0xf6, 0xa7, 0xa2, 0x3b, 0x2f, 0x43, 0x51, 0x32, 0xd8, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x17, 0x72, 0x66, 0xbc, 0x10, 0x1b, 0x46, 0x8b, 0xec, 0x49, 0x89, 0x74, 0x83,
	0x60, 0xb4, 0xc8, 0xd5, 0x74, 0x33, 0xd5, 0x35, 0x66, 0x4e, 0xd8, 0xd9, 0xd3, 0x7f, 0xf6, 0xb5,
	0x53, 0xd4, 0x0b, 0x63, 0xab, 0x9e, 0x05, 0x6e, 0x11, 0xc0, 0xf0, 0xd6, 0x55, 0xc4, 0x8d, 0xbb,
	0xdd, 0x5a, 0xda, 0x99, 0xd5, 0x53, 0xa5, 0x1b, 0x15, 0x33, 0x8e, 0x3f, 0x44, 0x5c, 0xd7, 0xc5,
	0x54, 0x95, 0xcc, 0x35, 0x11, 0x1a, 0x88, 0x89, 0x92, 0x27, 0x51, 0x93, 0x4c, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x29, 0x9e, 0xb5, 0xd0, 0x32, 0x1d, 0x2b, 0xab, 0xd8,
	0x99, 0x3a, 0x6b, 0x7c, 0x0f, 0x7f, 0x8a, 0xa2, 0xbd, 0xbf, 0x06, 0xa1, 0x0d, 0x92, 0xf6, 0xb5,
	0x15, 0x8d, 0xa9, 0xd2, 0xdd, 0x44, 0x4d, 0xe3, 0x4b, 0x45, 0x04, 0x2e, 0xf3, 0x54, 0x73, 0xfe,
	0x0f, 0x42, 0x21, 0x32, 0xde, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xbd, 0xaa,
	0xc5, 0x33, 0x67, 0xd2, 0xdd, 0x3c, 0x8e, 0x4c, 0xd6, 0x9e, 0x8a, 0xa2, 0x56, 0xf9, 0xe4, 0xd6,
	0xfc, 0x1c, 0xa4, 0x42, 0x4b, 0x3a, 0xaf, 0x55, 0xa5, 0x36, 0x97, 0xe8, 0x23, 0x5e, 0x15, 0x54,
	0xd5, 0x10, 0xfa, 0xec, 0xed, 0xff, 0x00, 0xe8, 0x0d, 0x5d, 0x70, 0x01, 0xa6, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x56, 0x27, 0xd0, 0x66, 0x17, 0xc4, 0xd8, 0x96,
	0xbe, 0xf7, 0x77, 0xaa, 0xbb, 0xcb, 0x57, 0x58, 0xf4, 0x7b, 0xda, 0xca, 0x86, 0xb1, 0x8d, 0xc9,
	0xa8, 0xd4, 0x44, 0x44, 0x6e, 0x79, 0x22, 0x22, 0x75, 0x9c, 0x10, 0x6a, 0xf7, 0x80, 0x63, 0x4e,
	0x9d, 0x0d, 0x74, 0xbf, 0xcf, 0x5d, 0x27, 0xf8, 0x54, 0x25, 0xc0, 0x16, 0xa2, 0xaf, 0xcc, 0x0e,
	0x8f, 0x7f, 0x84, 0x54, 0x7f, 0x5d, 0x37, 0xe2, 0x1f, 0x98, 0x1d, 0x1e, 0xff, 0x00, 0x09, 0xa8,
	0xfe, 0xba, 0x6f, 0xc4, 0x4a, 0xa0, 0x15, 0x0f, 0x54, 0xea, 0xed, 0x80, 0xe6, 0xfd, 0x5c, 0x37,
	0x38, 0x3f, 0xdb, 0xad, 0x72, 0xff, 0x00, 0xdb, 0x33, 0xa7, 0xad, 0xd5, 0x97, 0x0b, 0xc8, 0xc5,
	0xf9, 0x1d, 0xde, 0xf5, 0x03, 0xfa, 0x95, 0xee, 0x8a, 0x44, 0xfb, 0x36, 0x13, 0xe2, 0x4f, 0x20,
	0x42, 0xea, 0xae, 0xdd, 0x35, 0x5e, 0xad, 0x63, 0x55, 0x6d, 0x38, 0x9e, 0x9e, 0x57, 0x75, 0x36,
	0xaa, 0x91, 0xcc, 0xfb, 0xda, 0xe5, 0xf8, 0x1a, 0x3d, 0xf7, 0x40, 0x78, 0xf2, 0xd5, 0xb6, 0xe8,
	0x6d, 0xf4, 0xd7, 0x28, 0x9a, 0xbf, 0x3a, 0x8a, 0xa1, 0xaa, 0xab, 0xf5, 0x5f, 0xb2, 0xa5, 0xda,
	0x04, 0x85, 0x79, 0xbb, 0x77, 0xb4, 0x5c, 0x6c, 0xb5, 0x0b, 0x4f, 0x78, 0xa0, 0xab, 0xa0, 0x99,
	0x17, 0x2d, 0x8a, 0x98, 0x5d, 0x1a, 0xaf, 0x96, 0x69, 0xbf, 0xd0, 0xfc, 0x27, 0xa4, 0xf7, 0x0b,
	0x7d, 0x1d, 0xca, 0x99, 0xd4, 0xd7, 0x1a, 0x58, 0x2a, 0xa9, 0xdf, 0xf3, 0xa2, 0x9e, 0x34, 0x91,
	0xab, 0xe6, 0x8b, 0xb8, 0x87, 0xf1, 0xbe, 0xaf, 0x38, 0x5a, 0xf6, 0xc9, 0x26, 0xb0, 0xac, 0x96,
	0x2a, 0xd5, 0xde, 0x9c, 0x8f, 0x4e, 0x07, 0x2f, 0x8c, 0x6a, 0xbb, 0xbe, 0xaa, 0xa0, 0x8b, 0x54,
	0xe0, 0x1b, 0xb6, 0x90, 0xb4, 0x65, 0x89, 0x70, 0x24, 0xaa, 0xeb, 0xc5, 0x1f, 0x2b, 0x40, 0xab,
	0x93, 0x2b, 0xe9, 0xb3, 0x7c, 0x2b, 0xd9, 0x9a, 0xf1, 0x62, 0xf8, 0x39, 0x13, 0xc3, 0x33, 0x49,
	0x22, 0x80, 0x00, 0x00, 0x00, 0x00, 0x1b, 0x9e, 0x8f, 0xf4, 0x6b, 0x89, 0x71, 0xd4, 0xc8, 0xb6,
	0x5a, 0x2d, 0x8a, 0x14, 0x5c, 0x9f, 0x5d, 0x51, 0x9b, 0x20, 0x6f, 0x6e, 0x4b, 0xc5, 0xcb, 0xe0,
	0xd4, 0x5f, 0x1c, 0x80, 0xd3, 0x0f, 0xdb, 0x69, 0xb5, 0x5c, 0x2f, 0x15, 0x29, 0x4f, 0x68, 0xa1,
	0xaa, 0xae, 0x9d, 0x57, 0x2d, 0x8a, 0x68, 0x5d, 0x22, 0xfa, 0xe4, 0x9b, 0xbd, 0x4b, 0x6f, 0x82,
	0x35, 0x77, 0xc3, 0x16, 0x66, 0xc7, 0x36, 0x20, 0x74, 0x97, 0xca, 0xc4, 0xde, 0xa9, 0x2e, 0x71,
	0xc0, 0xd5, 0xf0, 0x8d, 0x17, 0x7f, 0xd6, 0x55, 0xf2, 0x26, 0x2b, 0x6d, 0xba, 0x8a, 0xd7, 0x4a,
	0xda, 0x6b, 0x6d, 0x25, 0x3d, 0x25, 0x3b, 0x7e, 0x6c, 0x50, 0x46, 0x91, 0xb5, 0x3d, 0x13, 0x71,
	0x62, 0x55, 0x2d, 0xb1, 0x68, 0x17, 0x1e, 0xdd, 0x51, 0x8e, 0x96, 0xdd, 0x4f, 0x6d, 0x89, 0xcb,
	0xf3, 0xab, 0x6a, 0x1a, 0xd5, 0x4f, 0x1d, 0x96, 0xed, 0x29, 0xbc, 0xda, 0xb5, 0x5e, 0xae, 0x7b,
	0x73, 0xbb, 0x62, 0x6a, 0x68, 0x5d, 0xfb, 0xb4, 0xb4, 0x8e, 0x93, 0xef, 0x73, 0x93, 0xe0, 0x5a,
	0x30, 0x22, 0x5d, 0x40, 0xb4, 0x3a, 0xb2, 0x61, 0x88, 0xd8, 0x9f, 0x2d, 0xbc, 0x5e, 0x67, 0x7f,
	0x5a, 0xb1, 0xd1, 0x46, 0x9f, 0x66, 0xc2, 0xfc, 0x4e, 0xe2, 0x9b, 0x57, 0x5c, 0x09, 0x17, 0xeb,
	0x21, 0xb9, 0xcf, 0xfe, 0xe5, 0x6b, 0x93, 0xfe, 0xa8, 0x84, 0xc4, 0x0b, 0x0a, 0x8a, 0x93, 0x40,
	0x3a, 0x3e, 0xfe, 0x13, 0x51, 0xfd, 0x74, 0xdf, 0x88, 0x7e, 0x60, 0x74, 0x7b, 0xfc, 0x22, 0xa3,
	0xfa, 0xd9, 0xbf, 0x11, 0x2a, 0x80, 0x95, 0x11, 0xcf, 0xab, 0xde, 0x00, 0x91, 0x3a, 0x14, 0x35,
	0xd1, 0x7f, 0x25, 0x74, 0x9f, 0xe5, 0x54, 0xe4, 0xc3, 0xba, 0x08, 0xc2, 0xd8, 0x77, 0x11, 0xdb,
	0xef, 0x56, 0x9a, 0xab, 0xc4, 0x55, 0x74, 0x52, 0xa4, 0xac, 0x6b, 0xaa, 0x5a, 0xf6, 0x3b, 0x72,
	0xa2, 0xa2, 0xa2, 0xb7, 0x3c, 0x95, 0x15, 0x53, 0x89, 0x2c, 0x00, 0x50, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x07, 0x1d, 0x4c, 0x11, 0x54, 0xc1, 0x24, 0x35, 0x31, 0x32, 0x58, 0x64, 0x6a, 0xb5, 0xf1,
	0xc8, 0xd4, 0x73, 0x5c, 0x8b, 0xc5, 0x15, 0x17, 0x72, 0xa1, 0x59, 0xf4, 0xcf, 0xa0, 0x36, 0xc1,
	0x14, 0xf7, 0xbc, 0x07, 0x0b, 0xb6, 0x1a, 0x8a, 0xf9, 0xed, 0x4d, 0xdf, 0xbb, 0xad, 0x61, 0xfc,
	0x1f, 0xf1, 0xec, 0x2c, 0xe0, 0x05, 0x79, 0xa0, 0xa8, 0xa8, 0xaa, 0x8b, 0xb9, 0x53, 0x71, 0x82,
	0xce, 0xeb, 0x2d, 0xa2, 0x86, 0x3e, 0x1a, 0x9c, 0x61, 0x87, 0x29, 0xf6, 0x65, 0x67, 0x4e, 0xe5,
	0x4d, 0x1b, 0x7e, 0x7a, 0x75, 0xcc, 0xd4, 0x4e, 0xb4, 0xfd, 0xae, 0xd4, 0xdf, 0xc5, 0x17, 0x3a,
	0xc4, 0x65, 0xbc, 0x0c, 0x98, 0x2c, 0xbe, 0xad, 0x3a, 0x28, 0x64, 0x91, 0xd3, 0xe3, 0x0c, 0x49,
	0x4e, 0x8e, 0xcf, 0xa7, 0x6d, 0xa6, 0x91, 0xbb, 0xb2, 0xea, 0x9d, 0xc8, 0xbf, 0xda, 0x9f, 0x5b,
	0xb0, 0x1b, 0xaf, 0xcf, 0xa1, 0x7d, 0x02, 0x7c, 0xb2, 0x28, 0x2f, 0x78, 0xee, 0x17, 0xb6, 0x07,
	0x22, 0x3e, 0x0b, 0x5a, 0xe6, 0xd5, 0x72, 0x75, 0x3a, 0x6e, 0xb4, 0x4f, 0xa1, 0xf6, 0xf6, 0x16,
	0x72, 0x96, 0x9a, 0x0a, 0x4a, 0x78, 0xe9, 0xe9, 0x61, 0x8e, 0x18, 0x23, 0x6a, 0x35, 0x91, 0xc6,
	0xd4, 0x6b, 0x5a, 0x89, 0xc1, 0x11, 0x13, 0x72, 0x21, 0xca, 0x0d, 0x30, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc3, 0xd8,
	0xd9, 0x18, 0xe6, 0x3d, 0xa8, 0xe6, 0xb9, 0x32, 0x54, 0x54, 0xcd, 0x15, 0x0a, 0x2d, 0xa7, 0x6c,
	0x0a, 0x98, 0x1b, 0x1c, 0x4d, 0x05, 0x23, 0x15, 0xb6, 0x9a, 0xe4, 0x5a, 0x9a, 0x2e, 0xc6, 0xb5,
	0x57, 0xa5, 0x1f, 0xd5, 0x5f, 0xb9, 0x5a, 0x5e, 0xa2, 0x29, 0xd6, 0x4f, 0x0a, 0x26, 0x24, 0xd1,
	0xb5, 0x5d, 0x4c, 0x31, 0xed, 0x57, 0x5a, 0x57, 0xe5, 0xb0, 0xaa, 0x26, 0xf5, 0x6a, 0x27, 0xe9,
	0x1b, 0xea, 0xcc, 0xd7, 0xcd, 0xa8, 0x4d, 0x5c, 0x56, 0x5d, 0x08, 0x60, 0x65, 0xc7, 0x78, 0xe2,
	0x9e, 0x8e, 0xa5, 0x8a, 0xb6, 0xaa, 0x54, 0x4a, 0x9a, 0xd5, 0xea, 0x56, 0x22, 0xee, 0x67, 0x9b,
	0x97, 0x77, 0x96, 0xd7, 0x61, 0x7b, 0xa2, 0x8d, 0x91, 0x44, 0xc8, 0xe2, 0x63, 0x59, 0x1b, 0x11,
	0x1a, 0xd6, 0xb5, 0x32, 0x44, 0x44, 0xe0, 0x88, 0x84, 0x4b, 0xab, 0x26, 0x14, 0x6e, 0x1e, 0xd1,
	0xc4, 0x15, 0xf3, 0x47, 0xb3, 0x5d, 0x78, 0x5f, 0x95, 0xc8, 0xaa, 0x99, 0x2a, 0x47, 0xc2, 0x26,
	0xf9, 0x6c, 0xf4, 0xbe, 0xb2, 0x92, 0xe0, 0xc3, 0x40, 0x01, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38, 0xea, 0x60,
	0x8a, 0xa6, 0x9e, 0x58, 0x27, 0x62, 0x49, 0x0c, 0xad, 0x56, 0x3d, 0x8e, 0xe0, 0xe6, 0xaa, 0x64,
	0xa8, 0xbe, 0x80, 0x01, 0x8a, 0x3a, 0x68, 0x68, 0xe9, 0x20, 0xa5, 0xa5, 0x8d, 0xb1, 0xc1, 0x0b,
	0x1b, 0x1c, 0x6c, 0x6f, 0x06, 0xb5, 0xa9, 0x92, 0x22, 0x7a, 0x21, 0xca, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f,
	0xff,0xd9,
}