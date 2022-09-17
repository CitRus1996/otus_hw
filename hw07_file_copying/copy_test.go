package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	type input struct {
		name             string
		sourcePath       string
		targetPath       string
		offset           int64
		limit            int64
		expectedErr      error
		expectedFilePath string
	}
	cases := []input{
		{
			"offset0_limit0",
			"./testdata/input.txt",
			"./testdata/out_test1.txt",
			0,
			0,
			nil,
			"./testdata/out_offset0_limit0.txt",
		},
		{
			"offset0_limit10",
			"./testdata/input.txt",
			"./testdata/out_test2.txt",
			0,
			10,
			nil,
			"./testdata/out_offset0_limit10.txt",
		},
		{
			"offset0_limit1000",
			"./testdata/input.txt",
			"./testdata/out_test3.txt",
			0,
			1000,
			nil,
			"./testdata/out_offset0_limit1000.txt",
		},
		{
			"offset0_limit10000",
			"./testdata/input.txt",
			"./testdata/out_test4.txt",
			0,
			10000,
			nil,
			"./testdata/out_offset0_limit10000.txt",
		},
		{
			"offset100_limit1000",
			"./testdata/input.txt",
			"./testdata/out_test5.txt",
			100,
			1000,
			nil,
			"./testdata/out_offset100_limit1000.txt",
		},
		{
			"offset6000_limit1000",
			"./testdata/input.txt",
			"./testdata/out_test6.txt",
			6000,
			1000,
			nil,
			"./testdata/out_offset6000_limit1000.txt",
		},
		//{
		//	"offset0_limit10",
		//	"./testdata/input.txt",
		//	"./testdata/out_test2.txt",
		//	0,
		//	10,
		//	nil,
		//	"./testdata/out_offset0_limit10.txt",
		//},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			err := Copy(
				testCase.sourcePath,
				testCase.targetPath,
				testCase.offset,
				testCase.limit,
			)
			expectedFileData, err := os.ReadFile(testCase.expectedFilePath)
			if err != nil {
				t.Fatalf("can't read expected file: %s", err)
			}

			resultFileData, err := os.ReadFile(testCase.targetPath)
			if err != nil {
				t.Fatalf("can't read result file: %s", err)
			}

			os.Remove(testCase.targetPath)

			require.ErrorIs(t, err, testCase.expectedErr)
			require.Equal(t, expectedFileData, resultFileData)
		})

	}
}
