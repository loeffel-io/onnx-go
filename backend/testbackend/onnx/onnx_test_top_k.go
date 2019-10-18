package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("TopK", "TestTopK", NewTestTopK)
}

// NewTestTopK version: 4.
func NewTestTopK() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "TopK",
		Title:  "TestTopK",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x86, 0x1, 0xa, 0x1d, 0xa, 0x1, 0x78, 0xa, 0x1, 0x6b, 0x12, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0x4, 0x54, 0x6f, 0x70, 0x4b, 0x12, 0xa, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x5a, 0xf, 0xa, 0x1, 0x6b, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x1, 0x62, 0x18, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x62, 0x19, 0xa, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x7, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "k"},
		     Output:    []string{"values", "indices"},
		     Name:      "",
		     OpType:    "TopK",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{3}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 3),
				tensor.WithBacking([]float32{3, 2, 1, 7, 6, 5, 11, 10, 9}),
			),

			tensor.New(
				tensor.WithShape(3, 3),
				tensor.WithBacking([]int64{3, 2, 1, 3, 2, 1, 3, 2, 1}),
			),
		},
	}
}
