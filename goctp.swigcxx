/* Copyright 2011 The Go Authors.  All rights reserved.
   Use of this source code is governed by a BSD-style
   license that can be found in the LICENSE file.  */

/* An example of writing a C++ virtual function in Go.  */

%module(directors="1") goctp

%init %{ 
  //printf("Initialization goctp done.\n");
%}

%typemap(gotype) (char **ppInstrumentID, int nCount) "[]string"

%typemap(in) (char **ppInstrumentID, int nCount)
%{
  {
    int i;
    _gostring_* a;

    $2 = $input.len;
    a = (_gostring_*) $input.array;
    $1 = (char **) malloc (($2 + 1) * sizeof (char *));
    for (i = 0; i < $2; i++) {
      
      /* Not work */
      //_gostring_ *ps = &a[i];
      //$1[i] = (char *) ps->p;
      //$1[i][ps->n] = '\0';

      /*Work well*/
      _gostring_ *ps = &a[i];
      $1[i] = (char*) malloc(ps->n + 1);
      memcpy($1[i], ps->p, ps->n);
      $1[i][ps->n] = '\0';
    }
    $1[i] = NULL;
  }
%}

%typemap(argout) (char **ppInstrumentID, int nCount) "" /* override char *[] default */

%typemap(freearg) (char **ppInstrumentID, int nCount)
%{
  {
    int i;
    for (i = 0; i < $2; i++)
    {
      free ($1[i]);
    }
    free($1);
  }
%}
 

/* 在复杂的企业环境中，可能有一些 C/C++ 头文件定义了您希望向脚本框架公开的全局变量和常量。
 * 在接口文件中使用 %include <header.h> 和 %{ #include <header.h> %}，可解决在头文件中重复所有元素的声明的问题。
 */

/* Includes the header files in the wrapper code */
%header %{
#if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__) || defined(WIN32)

#include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcUserApiDataType.h" 
#include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcUserApiStruct.h"
#include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcMdApi.h"
#include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcTraderApi.h"

#else

#include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcUserApiDataType.h" 
#include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcUserApiStruct.h"
#include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcMdApi.h"
#include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcTraderApi.h"

#endif
%}

/* Parse the header files to generate wrappers */
%include "std_string.i"

%feature("director") CThostFtdcMdSpi;
%feature("director") CThostFtdcTraderSpi;

#if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__) || defined(WIN32)

%include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcUserApiDataType.h" 
%include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcUserApiStruct.h"
%include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcMdApi.h"
%include "./api/ThostTraderApi_v6.3.6_20160606_win32/ThostFtdcTraderApi.h"

#else

%include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcUserApiDataType.h" 
%include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcUserApiStruct.h"
%include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcMdApi.h"
%include "./api/ThostTraderApi_v6.3.6_20160606_linux64/ThostFtdcTraderApi.h"

#endif

