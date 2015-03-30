// Copyright (C) 2013 Max Riveiro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package phash is a simple pHash wrapper library for the Go programming language.
package phash

/*
#cgo pkg-config: pHash
#include <stdlib.h>


typedef unsigned long long ulong64;
typedef unsigned char uint8_t;
typedef struct ph_digest {
    char *id;                   //hash id
    uint8_t *coeffs;            //the head of the digest integer coefficient array
    int size;                   //the size of the coeff array
} Digest;

extern int ph_compare_images_Wrapper(const char *file1, const char *file2, double *pcc, double sigma, double gamma, int N, double threshold);
extern ulong64 pc_dct_imagehash_Wrapper(const char *file);
extern int ph_image_digest_Wrapper(const char *file, double sigma, double gamma, Digest *digest, int N);
extern int ph_hamming_distance(ulong64 hasha, ulong64 hashb);
extern int ph_crosscorr_Wrapper(const Digest *x,const Digest *y,double *pcc,double threshold);
*/
import "C"

import "unsafe"
//import "fmt"

type Digest struct {
	DS *C.Digest
}

// ImageHash returns a DCT pHash for image with a given path.
func ImageHashDCT(file string) (uint64, error) {
	cs := C.CString(file)

	h, err := C.pc_dct_imagehash_Wrapper(cs)
	C.free(unsafe.Pointer(cs))

	return uint64(h), err
}

// HammingDistanceForHashes returns a Hamming Distance between two images' DCT pHashes.
func HammingDistanceForHashes(hasha uint64, hashb uint64) (int, error) {
	d, err := C.ph_hamming_distance(C.ulong64(hasha), C.ulong64(hashb))

	return int(d), err
}

// HammingDistanceForFiles returns a Hamming Distance between two images with a given paths.
func HammingDistanceForFiles(filea string, fileb string) (interface{}, error) {
	hasha, err := ImageHashDCT(filea)
	if err != nil {
		return nil, err
	}

	hashb, err := ImageHashDCT(fileb)
	if err != nil {
		return nil, err
	}

	return HammingDistanceForHashes(hasha, hashb)
}

func CompareImages(filea, fileb string) (res float64) {
	var n = C.int(180)
	var pcc, sigma, gamma, threshold C.double
	sigma = C.double(3.5)
	gamma = C.double(1.0)
	threshold = C.double(0.90)
	C.ph_compare_images_Wrapper(C.CString(filea), C.CString(fileb), &pcc, sigma, gamma, n, threshold)
	return float64(pcc)
}

func ImageDigest(file string) (d Digest) {
	var digest C.Digest
	var n = C.int(180)
	var sigma, gamma C.double
	sigma = C.double(3.5)
	gamma = C.double(1.0)
	C.ph_image_digest_Wrapper(C.CString(file), sigma, gamma, &digest, n)
	return Digest{DS :&digest}
} 

func ImageCrossCorr(d1,d2 Digest) (res float64) {
	var pcc, threshold C.double
	threshold = 0.90
	C.ph_crosscorr_Wrapper(d1.DS, d2.DS, &pcc, threshold)
	return float64(pcc)
}

func (d Digest) Bytes() (bytes []byte) {
	ulen := uintptr(d.DS.size)
	buf := make([]byte, ulen)

	uptr := uintptr(unsafe.Pointer(d.DS.coeffs))

	for i := uintptr(0); i < ulen; i++ {
		ptr := unsafe.Pointer(uptr + i)
		buf[i] = byte(*(*C.uint8_t)(ptr))
	}
	return buf
}

