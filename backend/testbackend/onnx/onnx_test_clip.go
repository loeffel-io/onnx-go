package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Clip", "TestClip", NewTestClip)
}

// NewTestClip version: 3.
func NewTestClip() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Clip",
		Title:  "TestClip",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x69, 0xa, 0x2a, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x4, 0x43, 0x6c, 0x69, 0x70, 0x2a, 0xd, 0xa, 0x3, 0x6d, 0x61, 0x78, 0x15, 0x0, 0x0, 0x80, 0x3f, 0xa0, 0x1, 0x1, 0x2a, 0xd, 0xa, 0x3, 0x6d, 0x69, 0x6e, 0x15, 0x0, 0x0, 0x80, 0xbf, 0xa0, 0x1, 0x1, 0x12, 0x9, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x70, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Clip",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000127300)(name:"max" type:FLOAT f:1 ),
		    (*ir.AttributeProto)(0xc000127400)(name:"min" type:FLOAT f:-1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1, 0.4001572, 0.978738, 1, 1, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1, -0.20515826, 0.3130677, -0.85409576, -1, 0.6536186, 0.8644362, -0.742165, 1, -1, 0.045758516, -0.18718386, 1, 1, 0.15494743, 0.37816253, -0.88778573, -1, -0.34791216, 0.15634897, 1, 1, -0.3873268, -0.30230275, -1, -1, -1, 1, -0.5096522, -0.4380743, -1, 0.7774904, -1, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117}),
			),
		},
	}
}
