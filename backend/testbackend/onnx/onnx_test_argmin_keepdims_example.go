package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ArgMin", "TestArgminKeepdimsExample", NewTestArgminKeepdimsExample)
}

// NewTestArgminKeepdimsExample version: 3.
func NewTestArgminKeepdimsExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ArgMin",
		Title:  "TestArgminKeepdimsExample",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x86, 0x1, 0xa, 0x34, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x6, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6, 0x41, 0x72, 0x67, 0x4d, 0x69, 0x6e, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x2a, 0xf, 0xa, 0x8, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x1c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x6d, 0x69, 0x6e, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0x16, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x18, 0xa, 0x6, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0xe, 0xa, 0xc, 0x8, 0x7, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x1, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data"},
		     Output:    []string{"result"},
		     Name:      "",
		     OpType:    "ArgMin",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000434000)(name:"axis" type:INT i:1 ),
		    (*ir.AttributeProto)(0xc000434100)(name:"keepdims" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 2),
				tensor.WithBacking([]float32{2, 1, 3, 10}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 1),
				tensor.WithBacking([]int64{1, 0}),
			),
		},
	}
}
