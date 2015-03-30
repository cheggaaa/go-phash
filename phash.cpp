// Copyright (C) 2013 Max Riveiro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

#include <pHash.h>

#ifdef __cplusplus
extern "C" {
#endif

ulong64 pc_dct_imagehash_Wrapper(const char *file) {
    cimg::exception_mode(0);
    ulong64 hash;

    if (ph_dct_imagehash(file, hash) == 0)
        errno = 0;

    return hash;
}

int ph_compare_images_Wrapper(const char *file1, const char *file2,double &pcc, double sigma, double gamma, int N,double threshold) {
	cimg::exception_mode(0);
	return ph_compare_images(file1,file2,pcc,sigma,gamma,N,threshold);
}

int ph_image_digest_Wrapper(const char *file, double sigma, double gamma, Digest &digest, int N){
	cimg::exception_mode(0);
	return ph_image_digest(file,sigma,gamma,digest,N);
}


int ph_crosscorr_Wrapper(const Digest &x,const Digest &y,double &pcc,double threshold){
	cimg::exception_mode(0);
	return ph_crosscorr(x,y,pcc,threshold);
}

#ifdef __cplusplus
}
#endif
