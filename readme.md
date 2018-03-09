# webshell-scanner-client

* [Introduction](#introduction)
* [CLI Usage](#cli-usage)
* [API Usage](#api-usage)
  * [Enqueue API](#enqueue-api)
  * [Result API](#result-api)

## Introduction

A Golang API client of [https://scanner.baidu.com](https://scanner.baidu.com). It determines whether a server side script is malicious or not.

### Supported file formats

Maximum allowed file size is currently `50 MB`, list of supported file formats are as below:

| Language   | File extensions                             |
| ---------- | ------------------------------------------- |
| PHP        | .php .phtml .php3 .php4 .php5 .inc          |
| ASP        | .asp .asa .cer .cdx .asa                    |
| ASP.NET    | .aspx .asmx .ashx                           |
| Java       | .jsp .jspx .war                             |
| Archive    | .rar .zip .tar .xz .tbz .tgz .tbz2 .bz2 .gz |

If you submit an unsupported file type it will be rejected.

### Detection capability

We do `sandbox analysis`, `code similarity analysis` and `static signature comparison`. Our current detection rate is `92%` among `6K` distinct webshell samples. 

Since we can't disclose our sample database, but here's a few interesting ones collected online:

* Obfuscated/Packed
  * [samples/conflg1.php](samples/conflg1.php)
  * [samples/chopper.php](samples/chopper.php)
* VBScript encoded
  * [samples/devshell.asp](samples/devshell.asp)
* Other tricks
  * [samples/unicode.php](samples/unicode.php)
  * [samples/plusplus.php](samples/plusplus.php)
* File droppers
  * [samples/crypt-dropper.php](samples/crypt-dropper.php)

Please see the [samples](samples) directory for more details.

## CLI usage

The binary release can be downloaded from the [Releases](https://github.com/baidu-security/webshell-scanner-client/releases) tab

Usage:

```
# ./tmp/webdir-linux32.bin
WebShell scanner API client - Copyright ©2017 Baidu Inc.
For more details visit: https://scanner.baidu.com

Usage: ./tmp/webdir-linux32.bin /tmp/a.php /tmp/b.php ...
```

Scanning multiple files:

```
# ./tmp/webdir-linux32.bin /tmp/obfuscation.tar /tmp/a.zip

2018/03/08 15:44:41 Submitting /tmp/obfuscation.tar ..
2018/03/08 15:44:41 Success. MD5 is b34d92050322d85e81f3c4d9358c2c92
2018/03/08 15:44:41 Task b34d92050322d85e81f3c4d9358c2c92 pending
2018/03/08 15:44:43 Task b34d92050322d85e81f3c4d9358c2c92 completed
+----------------------------------------+-------------------------+
|                FILENAME                |         RESULT          |
+----------------------------------------+-------------------------+
| obfuscation/unicode_var.php            | HEUR.WebShell.Chopper.X |
| obfuscation/chr.php                    | HEUR.WebShell.Chopper.X |
| obfuscation/comment.php                | HEUR.WebShell.Chopper.X |
| obfuscation/arithmetic.php             | -                       |
| obfuscation/code.php                   | HEUR.WebShell.Chopper.X |
| obfuscation/string_replace.php         | HEUR.WebShell.Exec.1    |
| obfuscation/string_return_function.php | HEUR.WebShell.Chopper.X |
| obfuscation/string_concat.php          | HEUR.WebShell.Chopper.X |
| obfuscation/slash_oct_string.php       | HEUR.WebShell.Chopper.X |
| obfuscation/__file__.php               | -                       |
| obfuscation/non-alphanumeric.php       | HEUR.WebShell.Various   |
| obfuscation/slash_hex_string.php       | HEUR.WebShell.Chopper.X |
+----------------------------------------+-------------------------+
|                                              SCANNED:     12     |
|                                              DETECTED:    10     |
|                                              TOTAL FILES: 12     |
+----------------------------------------+-------------------------+

2018/03/08 15:44:43 Submitting /tmp/a.zip ..
2018/03/08 15:44:43 Success. MD5 is 27b545a47aa420e0c0f30ac52eb3bbb8
2018/03/08 15:44:43 Task 27b545a47aa420e0c0f30ac52eb3bbb8 pending
2018/03/08 15:44:45 Task 27b545a47aa420e0c0f30ac52eb3bbb8 completed
+----------+-------------------------+
| FILENAME |         RESULT          |
+----------+-------------------------+
| b.php    | HEUR.WebShell.Exec.1    |
| a.php    | HEUR.WebShell.Chopper.X |
| a.js     | -                       |
+----------+-------------------------+
|                SCANNED:     3      |
|                DETECTED:    2      |
|                TOTAL FILES: 3      |
+----------+-------------------------+
```

## API Usage

The webshell scanner API runs asynchronously, all uploaded sample is queued for scanning. The result can be retrieved later via the `result` API.

### Enqueue API

Submit a sample to our scanning queue

```
curl https://scanner.baidu.com/enqueue -F archive=@web.zip
```

API response - on success

```
{
  // API return status and description
  "status": 0,
  "descr":  "Task enqueued",

  // sample hash
  "md5":    "de7128e140ac0add05d1728ab95d9d5d",

  // where to fetch scanning results
  "url":    "https://scanner.baidu.com/result/2744a8f7b85e799b79caf0ff014bfb3a",

  "data":   []
}
```

API response - on error

```
{
   // non-zero API status
   "status": 1

   // error description
   "descr":  "Extension 'xxx' is currently unsupported, list of supported file formats are: rar zip tar xz tbz tgz tbz2 bz2 gz php phtml inc php3 php4 php5 war jsp jspx asp aspx cer cdx asa ashx asmx cfm",

   "data" : [],   
}
```

## Result API

Retrieve the detection results, e.g

```
curl https://scanner.baidu.com/result/2744a8f7b85e799b79caf0ff014bfb3a
```

API response

```
[
  {
    // sample hash
    "md5": "de7128e140ac0add05d1728ab95d9d5d",

    // total files scanned
    "total": 1,

    // number of files scanned
    "scanned": 1,

    // number of detections
    "detected": 1,

    // status: pending/scanning/done
    "status": "done",

    // results
    "data": [
       {
          // relative path (only useful for archives)
          "path": "2744a8f7b85e799b79caf0ff014bfb3a.php",

          // produced verdict
          "descr": "BDS.WebShell.Chopper.1"
       }
    ]
  }
]
```

