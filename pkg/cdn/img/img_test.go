package img

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/NoUmlautsAllowed/gocook/pkg/env"
	"github.com/NoUmlautsAllowed/gocook/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

var testData = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
	0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x30, 0x08, 0x02, 0x00, 0x00, 0x00, 0xd8, 0x60, 0x6e,
	0xd0, 0x00, 0x00, 0x07, 0x1d, 0x49, 0x44, 0x41, 0x54, 0x58, 0xc3, 0xed, 0x58, 0x7d, 0x6c, 0x53,
	0xd7, 0x15, 0xff, 0xdd, 0xf7, 0x9e, 0xfd, 0xec, 0x7c, 0xf8, 0x19, 0x13, 0x62, 0x12, 0x12, 0xc8,
	0x28, 0x4e, 0x42, 0x42, 0xd2, 0xc5, 0x24, 0xc1, 0x7c, 0xb4, 0x19, 0x4c, 0xb0, 0xd1, 0x86, 0x0e,
	0x13, 0xba, 0x31, 0x95, 0xb4, 0x2a, 0x52, 0xab, 0x4d, 0x9b, 0xd4, 0x3f, 0x46, 0xa5, 0xee, 0x83,
	0x51, 0x54, 0xc1, 0x98, 0xd2, 0x76, 0xa2, 0x48, 0x65, 0x10, 0x4d, 0x62, 0x52, 0xd7, 0x30, 0x69,
	0x44, 0x09, 0x5b, 0x02, 0x74, 0x09, 0xa5, 0x4d, 0x33, 0xad, 0xa5, 0x8c, 0x24, 0x86, 0x84, 0xc5,
	0x28, 0xe1, 0xc3, 0xa1, 0x24, 0x24, 0x24, 0xf8, 0x23, 0xb6, 0x9f, 0xed, 0x77, 0xf6, 0x87, 0x93,
	0x60, 0xc7, 0x71, 0x02, 0xa3, 0x93, 0x22, 0x2d, 0x57, 0x4f, 0x4f, 0xf6, 0x7d, 0xf7, 0x9e, 0xf3,
	0x3b, 0xe7, 0xfe, 0xce, 0x79, 0xe7, 0x3c, 0x46, 0x98, 0x5d, 0x83, 0xc3, 0x1c, 0xa0, 0x39, 0x40,
	0x73, 0x80, 0xfe, 0xb7, 0x43, 0x80, 0x30, 0xbb, 0x00, 0x31, 0x7a, 0x77, 0xb6, 0x79, 0x28, 0x71,
	0x8e, 0x43, 0x33, 0x78, 0xe8, 0xa1, 0x0e, 0xf6, 0xeb, 0x50, 0x45, 0x5f, 0x13, 0x20, 0x62, 0xa8,
	0xf9, 0x92, 0x7d, 0xd1, 0xfb, 0xb8, 0x78, 0x2a, 0x57, 0x93, 0x79, 0xd1, 0xcc, 0xa6, 0x09, 0x33,
	0xa2, 0x39, 0xdb, 0x05, 0x6e, 0xd9, 0x4b, 0xdf, 0x59, 0x95, 0xfa, 0x58, 0xde, 0x21, 0x74, 0x5d,
	0xb7, 0x29, 0x8e, 0xd3, 0xc5, 0x19, 0x33, 0x60, 0x9a, 0x0e, 0x50, 0x90, 0x70, 0xa5, 0x8f, 0xdd,
	0x50, 0x6d, 0xdd, 0xf8, 0xcd, 0x82, 0xa5, 0xe9, 0x29, 0x20, 0x90, 0x6b, 0x29, 0xe8, 0x11, 0x69,
	0x27, 0xf8, 0x58, 0xa2, 0x83, 0x88, 0x5a, 0x44, 0xf5, 0xb5, 0x4e, 0x45, 0x1a, 0x38, 0x9b, 0x9d,
	0x3a, 0x3d, 0xa0, 0xf8, 0x80, 0x7d, 0x41, 0x3c, 0xf7, 0x1e, 0x3b, 0x77, 0xf2, 0xe9, 0xa5, 0x0b,
	0x0d, 0x08, 0xf1, 0x08, 0x26, 0xe1, 0x46, 0x09, 0x42, 0xaa, 0xa8, 0x45, 0x0a, 0xc1, 0xef, 0x87,
	0x46, 0x03, 0x06, 0x10, 0x10, 0x0c, 0x42, 0x88, 0x96, 0x99, 0x38, 0x82, 0xac, 0xfb, 0x4c, 0x70,
	0x3f, 0x55, 0x64, 0xfa, 0x94, 0xe7, 0x8e, 0xff, 0xe9, 0xa3, 0xfd, 0x5b, 0x89, 0xd1, 0x34, 0x80,
	0x5e, 0x8c, 0x8f, 0xd6, 0x03, 0xfc, 0x06, 0x50, 0x08, 0x0a, 0x0f, 0x77, 0x26, 0xf5, 0xac, 0x81,
	0x1c, 0x82, 0xa0, 0x80, 0x8b, 0x70, 0x92, 0xdb, 0xcd, 0xea, 0x5b, 0xe9, 0xf9, 0xf5, 0xd0, 0x88,
	0x08, 0x04, 0xd8, 0x5f, 0xce, 0x91, 0xb5, 0x0c, 0x5a, 0x4d, 0x84, 0x10, 0x3d, 0x75, 0x3e, 0xcb,
	0x72, 0x3f, 0x82, 0x6a, 0x04, 0x44, 0xc8, 0x07, 0x2a, 0xa7, 0x0f, 0x7b, 0x11, 0x71, 0x2f, 0xf5,
	0x78, 0x7c, 0x8d, 0x64, 0x53, 0xcf, 0x3a, 0x10, 0x63, 0xb5, 0xe7, 0xd1, 0x7f, 0x37, 0x4a, 0x40,
	0x62, 0x02, 0x7d, 0x6f, 0x2d, 0xd4, 0x6a, 0x00, 0x10, 0x78, 0xda, 0xb8, 0x12, 0x6a, 0x15, 0xd8,
	0xf8, 0xc6, 0xf0, 0x9d, 0x04, 0xea, 0x7a, 0x06, 0xee, 0xc5, 0x00, 0xc0, 0x03, 0xea, 0x31, 0xf9,
	0xa4, 0x1e, 0xbb, 0x22, 0x95, 0x3e, 0x0a, 0x21, 0x18, 0xa3, 0x0d, 0x45, 0x98, 0xa7, 0x8f, 0x9a,
	0xe4, 0x79, 0xe8, 0x92, 0xc1, 0x31, 0x00, 0xe0, 0x38, 0x5a, 0x30, 0xff, 0x54, 0x5b, 0x83, 0xf5,
	0x9d, 0x6d, 0xd6, 0x77, 0xb6, 0xb5, 0xf5, 0xb5, 0xa3, 0xe5, 0x22, 0x06, 0x06, 0xa7, 0xe4, 0xb8,
	0xdd, 0x0e, 0xab, 0x15, 0x56, 0x2b, 0x8e, 0x1c, 0x01, 0x51, 0x0c, 0xa9, 0x89, 0x62, 0xb5, 0x4f,
	0x85, 0x29, 0x35, 0x65, 0xda, 0x44, 0x43, 0x27, 0x2f, 0x9e, 0xec, 0x0c, 0x75, 0x95, 0x56, 0x94,
	0x02, 0xf8, 0xab, 0xfd, 0x94, 0x6f, 0xd4, 0x64, 0x11, 0x4d, 0xb1, 0x2b, 0x3b, 0x3a, 0x50, 0x57,
	0x97, 0x6f, 0xb1, 0xec, 0x54, 0x14, 0xa5, 0xbf, 0xdf, 0xf1, 0xc1, 0x07, 0x47, 0x2a, 0x2b, 0x23,
	0x00, 0x29, 0x0a, 0xea, 0xea, 0xa0, 0x28, 0x0f, 0x36, 0x64, 0x65, 0xa1, 0xb8, 0xf8, 0x21, 0x7c,
	0xe6, 0xf1, 0x20, 0x18, 0x82, 0xa4, 0x9b, 0xc8, 0x11, 0xad, 0x97, 0x5b, 0x75, 0x16, 0xc9, 0x5c,
	0x60, 0x66, 0x8c, 0x75, 0x6b, 0xba, 0x9b, 0xaf, 0xf7, 0xdc, 0xbf, 0xf5, 0xcf, 0x4d, 0xd2, 0x46,
	0x46, 0x63, 0xf6, 0xb5, 0xb5, 0xe1, 0xf8, 0x71, 0xdc, 0xb9, 0xb3, 0x45, 0x14, 0xd7, 0x14, 0x14,
	0x14, 0x10, 0x91, 0xdd, 0xae, 0x3d, 0x73, 0x86, 0x55, 0x56, 0xd2, 0x03, 0x40, 0x44, 0x78, 0xed,
	0x35, 0x94, 0x94, 0x6c, 0x4f, 0x4e, 0xd6, 0x01, 0xe4, 0x70, 0x38, 0xd6, 0xae, 0xfd, 0x7b, 0x72,
	0x32, 0xd2, 0xd2, 0x66, 0x02, 0x74, 0xbb, 0x9f, 0x0d, 0xbb, 0xa8, 0xf4, 0xc9, 0x29, 0x1f, 0x9a,
	0x4c, 0xa6, 0x6e, 0xea, 0x6e, 0xee, 0x68, 0xd6, 0x68, 0x45, 0x10, 0x83, 0xff, 0x4a, 0xfb, 0x6d,
	0xc7, 0xe9, 0xd3, 0x94, 0x92, 0x52, 0xb9, 0x7c, 0xf9, 0xf2, 0x15, 0x2b, 0x56, 0x10, 0x51, 0xdc,
	0x3c, 0x44, 0x84, 0xb2, 0xa7, 0xd7, 0x1b, 0x8d, 0x0b, 0x79, 0x9e, 0xbf, 0x72, 0xc5, 0xf6, 0x61,
	0xcd, 0xcd, 0x8e, 0x8e, 0x7f, 0x1f, 0x3e, 0x1c, 0xe5, 0xb6, 0xf1, 0x38, 0x57, 0xc2, 0x74, 0x01,
	0x80, 0x65, 0xdf, 0xa0, 0xe9, 0x28, 0xc7, 0x72, 0x72, 0x72, 0x00, 0xec, 0xae, 0x79, 0x7d, 0x62,
	0xb2, 0xa8, 0xa8, 0xa8, 0xbc, 0xbc, 0x3c, 0x21, 0x21, 0x21, 0x6e, 0xd8, 0x87, 0x80, 0x10, 0xc0,
	0x78, 0x1c, 0x38, 0xf8, 0x96, 0x75, 0x6b, 0xc5, 0xb7, 0x37, 0x6c, 0xca, 0xcb, 0x5b, 0xb1, 0xeb,
	0xe5, 0x57, 0x7e, 0xbd, 0xf7, 0xe7, 0xf5, 0xf5, 0x01, 0x81, 0x8f, 0x51, 0xd3, 0x71, 0x15, 0x6a,
	0x15, 0xe5, 0x99, 0x62, 0x89, 0xe6, 0xf6, 0xb9, 0x65, 0x25, 0xc0, 0xa2, 0x27, 0xb3, 0xb3, 0xb3,
	0xf7, 0xec, 0xd9, 0x13, 0x1d, 0x06, 0xfc, 0x24, 0xdc, 0xc1, 0x60, 0x82, 0xd3, 0xe9, 0x49, 0x48,
	0x02, 0xe3, 0xc0, 0x3b, 0xde, 0x44, 0x03, 0xc3, 0xbd, 0xe7, 0xe0, 0x22, 0xf7, 0xbd, 0x26, 0xdf,
	0x57, 0xb7, 0xfb, 0xb2, 0xb3, 0x73, 0x25, 0x49, 0x6f, 0x59, 0xb5, 0xe6, 0xe3, 0xf3, 0xe7, 0x40,
	0xca, 0x4f, 0x77, 0x6e, 0x32, 0x08, 0x8b, 0x71, 0x3f, 0x7d, 0x4c, 0x80, 0x56, 0x84, 0x94, 0x0c,
	0x51, 0x1c, 0xfb, 0x1b, 0x52, 0x58, 0xcb, 0x45, 0x2c, 0xd0, 0xcb, 0x02, 0xed, 0x7c, 0xaf, 0xd2,
	0xbc, 0xdd, 0x9c, 0x9b, 0x9b, 0x1b, 0x56, 0x19, 0x46, 0x46, 0x44, 0x3c, 0xcf, 0x73, 0x11, 0x63,
	0x02, 0x4a, 0x28, 0x14, 0xe2, 0x38, 0x4e, 0x92, 0xa4, 0x79, 0xf3, 0x56, 0x7d, 0xff, 0x85, 0xc6,
	0xcb, 0x2f, 0xa3, 0x51, 0x0b, 0xae, 0x1d, 0x68, 0x67, 0x50, 0x3f, 0x81, 0xf9, 0xcf, 0x62, 0xa8,
	0xc2, 0xd6, 0xd4, 0x7c, 0xf6, 0xf2, 0x65, 0x9b, 0xd7, 0x3b, 0x6a, 0x34, 0x2e, 0x7c, 0xe3, 0x8d,
	0x5f, 0x65, 0x2e, 0xce, 0x1a, 0xdb, 0x2d, 0xcb, 0xac, 0xd3, 0x0e, 0x59, 0x86, 0xa4, 0x83, 0x2e,
	0x39, 0xb2, 0x10, 0x20, 0xa3, 0x04, 0x9e, 0x07, 0x10, 0x0c, 0x05, 0x44, 0x51, 0x54, 0xa9, 0xc6,
	0x52, 0x79, 0x61, 0x61, 0xa1, 0xd1, 0x68, 0xac, 0xae, 0xae, 0x66, 0x31, 0x11, 0x9b, 0x99, 0x99,
	0x69, 0xb1, 0x58, 0xf6, 0xed, 0xdb, 0xe7, 0x72, 0xb9, 0x78, 0x9e, 0x17, 0x45, 0xf1, 0xce, 0x10,
	0x6c, 0x40, 0x7b, 0x64, 0x3d, 0x94, 0xb4, 0x16, 0x86, 0x0a, 0x88, 0xaf, 0xba, 0xaa, 0xaa, 0xf7,
	0xdb, 0xed, 0xdd, 0x23, 0x23, 0xc3, 0x0b, 0x52, 0x16, 0x88, 0xea, 0x09, 0x4f, 0x84, 0x30, 0x30,
	0x3c, 0x05, 0xa7, 0x38, 0x0e, 0xb9, 0x26, 0x68, 0xc4, 0x58, 0x36, 0x34, 0x36, 0x36, 0xd6, 0xd4,
	0xd4, 0x0c, 0x0d, 0x0d, 0x45, 0xba, 0xa4, 0xbe, 0xbe, 0xde, 0xe9, 0x74, 0xda, 0x6c, 0xb6, 0x63,
	0xc7, 0x8e, 0x5d, 0xba, 0x74, 0x29, 0x96, 0xd7, 0x42, 0xa4, 0xad, 0xda, 0x42, 0x2c, 0xd9, 0x0f,
	0x4f, 0x3b, 0xbd, 0xdd, 0x74, 0x40, 0xf1, 0x01, 0x5d, 0x10, 0xb8, 0x71, 0xe3, 0xb4, 0x5a, 0x5a,
	0x5f, 0x1a, 0x55, 0xd3, 0x0c, 0x0d, 0x23, 0x51, 0x0b, 0x8d, 0x66, 0x92, 0x44, 0x15, 0x51, 0x9e,
	0xcf, 0x0f, 0xc0, 0xe7, 0x19, 0xd5, 0xe9, 0x74, 0xe5, 0xe5, 0xe5, 0x91, 0x5a, 0xf5, 0x7a, 0x3d,
	0xc7, 0x71, 0x81, 0x40, 0x40, 0x96, 0xe5, 0x37, 0xf7, 0xee, 0x2d, 0x51, 0xa0, 0xf1, 0xfa, 0x54,
	0x81, 0x60, 0xfc, 0xb7, 0xbd, 0x88, 0xe5, 0x7f, 0x03, 0x00, 0xff, 0x75, 0x74, 0x59, 0x00, 0x17,
	0x7b, 0x50, 0x5b, 0x51, 0x44, 0x88, 0x91, 0xc2, 0x3e, 0xeb, 0xa0, 0x95, 0xd9, 0xc8, 0x98, 0x9c,
	0x1b, 0x78, 0xa2, 0x0c, 0x7f, 0x00, 0x80, 0x6b, 0xc9, 0x92, 0x80, 0xec, 0xf7, 0x7a, 0xbd, 0xe1,
	0x79, 0xa7, 0xd3, 0xe9, 0xf5, 0x7a, 0xcb, 0xca, 0xca, 0xc2, 0xf4, 0x4a, 0x4f, 0x4f, 0x4f, 0x4f,
	0x4b, 0xcb, 0x72, 0x7a, 0x04, 0x7f, 0xa0, 0x3f, 0x18, 0x9a, 0xae, 0x84, 0x65, 0x22, 0x98, 0x08,
	0x8d, 0x09, 0x85, 0x76, 0x70, 0x61, 0x75, 0x52, 0x0f, 0x5b, 0x72, 0x01, 0x00, 0x6b, 0xeb, 0x64,
	0x1d, 0x57, 0x01, 0x80, 0x71, 0xb4, 0x65, 0x1d, 0x16, 0x2d, 0x8c, 0xdc, 0x37, 0x3f, 0x39, 0xc5,
	0xe3, 0xf2, 0xf8, 0x7c, 0xbe, 0x20, 0x43, 0x90, 0x31, 0x05, 0xe8, 0xed, 0xed, 0x6d, 0x68, 0x68,
	0x90, 0x65, 0xd9, 0xef, 0xf7, 0xdb, 0x6c, 0xb6, 0x13, 0x27, 0x4e, 0x04, 0x02, 0x01, 0xbf, 0xdf,
	0xdf, 0xdb, 0xdb, 0x7b, 0xe8, 0xd0, 0x21, 0x00, 0x41, 0x06, 0x59, 0x51, 0x46, 0xfd, 0x5e, 0x55,
	0xf1, 0xd8, 0x8b, 0x2f, 0x6e, 0x3d, 0x14, 0x18, 0xc0, 0xdd, 0x3f, 0x80, 0xee, 0x83, 0x78, 0x0e,
	0x82, 0x0c, 0xdd, 0x75, 0x66, 0x72, 0x93, 0xab, 0x00, 0xca, 0x78, 0xf9, 0xc1, 0x45, 0x45, 0xaf,
	0x9a, 0x53, 0x1d, 0xfd, 0xd1, 0xef, 0x77, 0xff, 0x71, 0xb7, 0x6b, 0xb5, 0x8b, 0x95, 0x94, 0x80,
	0xb1, 0xf0, 0x49, 0xd5, 0xd6, 0xd6, 0xd6, 0xd6, 0xd6, 0x4e, 0x2c, 0x6b, 0x6a, 0x6a, 0x0a, 0xff,
	0xc8, 0xcf, 0xcf, 0x27, 0xe0, 0x9c, 0x2e, 0xe9, 0x9a, 0xdd, 0xfe, 0xcb, 0x2f, 0x5f, 0x2f, 0x68,
	0x02, 0x13, 0x00, 0x80, 0xad, 0x0c, 0x6a, 0x63, 0x5c, 0x94, 0xe0, 0xbf, 0xa9, 0xdc, 0x3d, 0x31,
	0x32, 0xf0, 0x0b, 0x10, 0xd1, 0xe6, 0x6f, 0x3d, 0x79, 0xf0, 0x67, 0x3b, 0x0a, 0x96, 0xa6, 0x03,
	0x0c, 0x43, 0x79, 0x20, 0x21, 0x5e, 0x81, 0x4c, 0xa0, 0xaf, 0x3c, 0x8e, 0xba, 0x96, 0xf3, 0x97,
	0xee, 0x8e, 0x58, 0xad, 0x56, 0x00, 0x5e, 0xaf, 0x77, 0x74, 0x74, 0x34, 0x5e, 0xda, 0x94, 0x24,
	0xa9, 0xaa, 0xaa, 0x6a, 0xf0, 0xa9, 0x3e, 0xd5, 0x8f, 0xef, 0xa8, 0xc7, 0xb3, 0x0a, 0x4b, 0xfc,
	0x73, 0xff, 0xe4, 0xb5, 0x37, 0x1b, 0x36, 0x0c, 0xb6, 0xac, 0x4a, 0xc9, 0x06, 0x41, 0xa3, 0xd1,
	0x70, 0x1c, 0xa7, 0xf1, 0xf7, 0x2b, 0xce, 0xdb, 0xf7, 0xdc, 0xa3, 0x61, 0xb7, 0x32, 0xc6, 0x52,
	0xa5, 0xa4, 0x78, 0x95, 0xdd, 0x17, 0x6d, 0xf6, 0x0b, 0xdd, 0xc3, 0x66, 0xb3, 0xf9, 0x21, 0xea,
	0x5a, 0xe5, 0x94, 0xb4, 0xcd, 0x69, 0xcc, 0xd4, 0x2c, 0xcb, 0x78, 0x40, 0x6a, 0xcf, 0x48, 0xea,
	0x64, 0x3b, 0xbb, 0x78, 0xa3, 0xa8, 0xca, 0x59, 0x97, 0x4b, 0x44, 0x8c, 0xb1, 0x7f, 0xb4, 0x7c,
	0x76, 0xb1, 0xdf, 0xf0, 0xe9, 0xf0, 0x32, 0x43, 0xeb, 0xbb, 0xbb, 0x5e, 0x7d, 0x85, 0x31, 0x26,
	0xcb, 0xf2, 0xc1, 0xdf, 0x1e, 0xf5, 0x7c, 0xf7, 0x2d, 0x68, 0xa5, 0x29, 0xf4, 0x68, 0x57, 0x63,
	0x51, 0x8f, 0xab, 0xa5, 0xe6, 0xf9, 0x1d, 0x3f, 0x20, 0x22, 0x8f, 0xc7, 0x73, 0xf4, 0xfd, 0xea,
	0x50, 0xc5, 0xdb, 0xb8, 0xfa, 0xc9, 0xe6, 0x0c, 0x4f, 0x7e, 0x61, 0x01, 0x03, 0xeb, 0xeb, 0xeb,
	0xfb, 0xb0, 0xb9, 0x13, 0x1b, 0x7e, 0x82, 0x27, 0xca, 0xa0, 0xd1, 0x7b, 0x46, 0x66, 0xaa, 0xa9,
	0x1d, 0xb7, 0x1c, 0xb6, 0xf6, 0x0e, 0x22, 0xe2, 0x38, 0xee, 0xe3, 0xc1, 0xcc, 0xb6, 0x9c, 0x2d,
	0x08, 0x8e, 0xaa, 0x3e, 0x3f, 0x64, 0x30, 0x18, 0x18, 0xc7, 0xfc, 0x72, 0x50, 0xd8, 0x76, 0x00,
	0xe6, 0x17, 0xa1, 0xd5, 0x4f, 0x6d, 0xfb, 0x8d, 0x56, 0xf1, 0x46, 0xad, 0x61, 0xbe, 0x81, 0x88,
	0xd4, 0x5a, 0x1d, 0xfb, 0x61, 0x15, 0x4a, 0x5f, 0x42, 0x9a, 0x29, 0x91, 0xda, 0x0d, 0x06, 0x30,
	0x06, 0x37, 0xe9, 0xf1, 0xcc, 0x26, 0x94, 0xec, 0x84, 0x12, 0xdb, 0x71, 0x1d, 0xa5, 0xc9, 0x1e,
	0xba, 0xf6, 0x09, 0x7a, 0x3e, 0x9f, 0x38, 0x10, 0x2a, 0xde, 0x01, 0xc3, 0x62, 0x0c, 0xda, 0x51,
	0x7f, 0x80, 0x85, 0xcf, 0x49, 0xd4, 0xd2, 0x0b, 0xbf, 0x03, 0x13, 0xe3, 0x1e, 0x46, 0xdf, 0xbf,
	0x70, 0xe6, 0xf0, 0x98, 0x04, 0x43, 0x3a, 0x59, 0xf7, 0x01, 0x02, 0x88, 0x70, 0xeb, 0x02, 0x6b,
	0x3a, 0x02, 0x02, 0x8a, 0x36, 0x93, 0x79, 0xfb, 0xd4, 0x31, 0x1e, 0x03, 0x28, 0xcc, 0x13, 0x8a,
	0x6a, 0x85, 0xa6, 0x68, 0xf5, 0xd8, 0xa3, 0x34, 0x85, 0x8f, 0x20, 0x41, 0x88, 0x23, 0x8d, 0x3d,
	0x5e, 0x03, 0xcb, 0xfe, 0xeb, 0x16, 0x78, 0x16, 0xf6, 0xf6, 0xb3, 0xec, 0xbb, 0x30, 0xdb, 0xb5,
	0x7e, 0x76, 0x21, 0x12, 0x04, 0xe7, 0xdc, 0xf7, 0xa1, 0x39, 0x40, 0x73, 0x80, 0xfe, 0xbf, 0x00,
	0xfd, 0x07, 0x65, 0x17, 0xdc, 0x9f, 0x66, 0x66, 0x1c, 0xd8, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45,
	0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}

func TestNewImageCdn(t *testing.T) {
	c := NewImageCdn(env.NewEnv())

	if c.defaultClient.Timeout == 0 {
		t.Error("no timeout set for default client; Timeout should be set, see documentation")
	}
}

func TestImageCdn_GetRawImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := ImageCdn{
		cdnURL:        s.URL + "/cdn",
		defaultClient: http.Client{},
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		e := json.NewEncoder(w)
		err := e.Encode(api.Recipe{})
		if err != nil {
			t.Error("expected no error")
		}
		if r.URL.Path != "/cdn/123456" {
			t.Error("expected 123456")
		}
	})

	r, err := a.GetRawImage(http.MethodGet, "123456")
	if err != nil {
		t.Error("did not expect error")
	}

	if r == nil {
		t.Error("recipe expected")
	}

	s.Close()
}

func TestImageCdn_GetRawImage2(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := ImageCdn{
		cdnURL:        s.URL + "/cdn",
		defaultClient: http.Client{},
	}

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotModified)
		if r.URL.Path != "/cdn/123456" {
			t.Error("expected 123456")
		}
		if r.Method != http.MethodHead {
			t.Error("expected HEAD method")
		}
	})

	r, err := a.GetRawImage(http.MethodHead, "123456")
	if err != nil {
		t.Error("did not expect error")
	}

	if len(r) > 0 {
		t.Error("no image expected with head method")
	}

	s.Close()
}

func TestImageCdn_GetRawImage3(t *testing.T) {
	// mal crafted url

	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	// only way to produce url join error is to put some weird control character into the base url
	a := ImageCdn{
		cdnURL:        s.URL + "\x01/cdn",
		defaultClient: http.Client{},
	}

	r, err := a.GetRawImage(http.MethodGet, "abcdefg/\x01")
	if err == nil {
		t.Error("did expect error")
	}

	if r != nil {
		t.Error("no image expected with error")
	}

	s.Close()
}

func TestImageCdn_GetRawImage4(t *testing.T) {
	// produce cdn timeout request

	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 20)
	})

	a := ImageCdn{
		cdnURL: s.URL + "/cdn",
		defaultClient: http.Client{
			Timeout: time.Millisecond * 10,
		},
	}

	r, err := a.GetRawImage(http.MethodGet, "123456")
	if err == nil {
		t.Error("did expect error")
	}

	if r != nil {
		t.Error("no image expected with request error")
	}

	s.Close()
}

func TestImageCdn_GetImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := ImageCdn{
		cdnURL:        s.URL + "/cdn",
		defaultClient: http.Client{},
	}

	u, _ := url.Parse(s.URL + "/cdn/123456")

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(testData)
		if r.URL.Path != "/cdn/123456" {
			t.Error("expected 123456")
		}
		if r.Method != http.MethodGet {
			t.Error("expected GET method")
		}
	})

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Method: http.MethodGet,
		URL:    u,
	}
	ctx.Params = gin.Params{
		{Key: "path", Value: "123456"},
	}

	a.GetImage(ctx)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("expected", http.StatusOK)
	}
}

func TestImageCdn_GetImage2(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := ImageCdn{
		cdnURL:        s.URL + "/cdn",
		defaultClient: http.Client{},
	}

	u, _ := url.Parse(s.URL + "/cdn/123456")

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.URL.Path != "/cdn/123456" {
			t.Error("expected 123456")
		}
		if r.Method != http.MethodHead {
			t.Error("expected HEAD method")
		}
	})

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Method: http.MethodHead,
		URL:    u,
	}
	ctx.Params = gin.Params{
		{Key: "path", Value: "123456"},
	}

	a.GetImage(ctx)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("expected", http.StatusOK)
	}
}

func TestImageCdn_GetImage3(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := ImageCdn{
		cdnURL:        s.URL + "/cdn",
		defaultClient: http.Client{},
	}

	u, _ := url.Parse(s.URL + "/cdn/123456")

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.URL.Path != "/cdn/123456" {
			t.Error("expected 123456")
		}
		if r.Method != http.MethodGet {
			t.Error("expected HEAD method")
		}
	})

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Method: http.MethodGet,
		URL:    u,
	}
	ctx.Params = gin.Params{
		{Key: "path", Value: "123456"},
	}

	a.GetImage(ctx)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Error("expected", http.StatusInternalServerError)
	}
}

func TestImageCdn_GetImage4(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := ImageCdn{
		cdnURL:        s.URL + "/cdn",
		defaultClient: http.Client{},
	}

	u, _ := url.Parse(s.URL + "/cdn/123456")

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.URL.Path != "/cdn/123456" {
			t.Error("expected 123456")
		}
		if r.Method != http.MethodHead {
			t.Error("expected HEAD method")
		}
	})

	responseWriter := utils.NewMockResponseWriter(ctrl)
	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = &http.Request{
		Method: http.MethodHead,
		URL:    u,
	}
	ctx.Params = gin.Params{
		{Key: "path", Value: "123456"},
	}

	responseWriter.EXPECT().WriteHeader(http.StatusOK)
	responseWriter.EXPECT().Write([]uint8{}).Return(0, errors.New("writer error"))
	responseWriter.EXPECT().Header().Return(http.Header{})
	responseWriter.EXPECT().Write([]byte{123, 34, 101, 114, 114, 111, 114, 34, 58, 34, 119, 114, 105, 116, 101, 114, 32, 101, 114, 114, 111, 114, 34, 125})

	a.GetImage(ctx)
}

func TestImageCdn_PostImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := ImageCdn{
		cdnURL:        s.URL + "/cdn",
		defaultClient: http.Client{},
	}

	u, _ := url.Parse(s.URL + "/cdn/123456")

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Method: http.MethodPost,
		URL:    u,
	}
	ctx.Params = gin.Params{
		{Key: "path", Value: "123456"},
	}

	a.GetImage(ctx)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Error("expected", http.StatusMethodNotAllowed)
	}
}

func TestImageCdn_UserAgent(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := utils.NewMockHandler(ctrl)

	s := httptest.NewServer(m)
	defer s.Close()

	a := NewImageCdn(env.NewEnv())
	a.cdnURL = s.URL + "/cdn"
	a.defaultClient = http.Client{}

	u, _ := url.Parse(s.URL + "/cdn/123456")

	m.EXPECT().ServeHTTP(gomock.Any(), gomock.Any()).Do(func(w http.ResponseWriter, r *http.Request) {
		expectedUserAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:108.0) Gecko/20100101 Firefox/108.0"
		if r.UserAgent() != expectedUserAgent {
			t.Error("expected user agent '" + expectedUserAgent + "', got '" + r.UserAgent() + "'")
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(testData)
		if r.URL.Path != "/cdn/123456" {
			t.Error("expected 123456")
		}
		if r.Method != http.MethodGet {
			t.Error("expected GET method")
		}
	})

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Method: http.MethodGet,
		URL:    u,
	}
	ctx.Params = gin.Params{
		{Key: "path", Value: "123456"},
	}

	a.GetImage(ctx)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("expected", http.StatusOK)
	}
}
