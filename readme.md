# webshell-scanner-client

* [Introduction](#introduction)
* [CLI Usage](#cli-usage)
* [API Usage](#api-usage)
  * [Enqueue API](#enqueue-api)
  * [Result API](#result-api)

## Introduction

A Golang API client of [https://scanner.baidu.com](https://scanner.baidu.com). The API verifies if a sample is malicious, e.g `command execution backdoor`, `China Chopper`, and other `multi-functional webshell`.

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
# ./tmp/webdir-linux32.bin /tmp/a.zip /tmp/a.php
2018/03/08 14:16:33 Submitting /tmp/a.zip ..
2018/03/08 14:16:33 Success. MD5 is fb22bd7d65def142c2471ef30f88fdf8
2018/03/08 14:16:33 Task fb22bd7d65def142c2471ef30f88fdf8 pending
+----------+------------------------+
| FILENAME |         RESULT         |
+----------+------------------------+
| b.php    | BDS.WebShell.Exec.1    |
| a.php    | BDS.WebShell.Chopper.1 |
+----------+------------------------+
|                SCANNED:     2     |
|                DETECTED:    2     |
|                TOTAL FILES: 2     |
+----------+------------------------+
2018/03/08 14:16:35 Submitting /tmp/a.php ..
2018/03/08 14:16:35 Success. MD5 is 695b4b31951eaed322c9127ec0d53011
2018/03/08 14:16:35 Task 695b4b31951eaed322c9127ec0d53011 pending
+--------------------------------------+------------------------+
|               FILENAME               |         RESULT         |
+--------------------------------------+------------------------+
| 695b4b31951eaed322c9127ec0d53011.php | BDS.WebShell.Chopper.1 |
+--------------------------------------+------------------------+
|                                            SCANNED:     1     |
|                                            DETECTED:    1     |
|                                            TOTAL FILES: 1     |
+--------------------------------------+------------------------+
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

