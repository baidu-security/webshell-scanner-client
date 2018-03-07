# webshell-scanner-client

[![中文帮助](https://img.shields.io/badge/README-切换语言-yellow.svg)](readme-zh_CN.md)

## Introduction

A Golang API client of [https://scanner.baidu.com](https://scanner.baidu.com). 

## List of supported file formats

| Language   | File extensions            |
| ---------- | -------------------------- |
| PHP        | .php .phtml .php[345] .inc |
| ASP        | .asp .asa .cer .cdx .asa   |
| ASP.NET    | .aspx .asmx .ashx          |
| Java       | .jsp .jspx .war            |
| Archive    | .rar .zip .tar .xz .tbz .tgz .tbz2 .bz2 .gz |

## API Usage

The webshell scanner API runs asynchronously. The verdict can be retrieved later via the `result` API.

### Enqueue API

Submit a file/archive to scanning queue

```
curl https://scanner.baidu.com/enqueue -F archive=@web.zip
```

API response

```
{
  // API errors: usually none
  "error":  false,

  // status: always set to pending
  "status": "pending",

  // calculated sample hash (md5)
  "hash":   "de7128e140ac0add05d1728ab95d9d5d",

  // where to fetch scanning results
  "url":    "https://scanner.baidu.com/result/2744a8f7b85e799b79caf0ff014bfb3a"
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
    // total files scanned
    "cnt": 1,

    // sample hash
    "md5": "de7128e140ac0add05d1728ab95d9d5d",

    // number of detections
    "detected": 1,

    // status: pending/scanning/done
    "status": "done",

    // number of files scanned
    "scanned": 1,

    // results
    "data": [
    {
        // relative path (only useful for archives)
        "path": "2744a8f7b85e799b79caf0ff014bfb3a.php",

        // produced verdict
        "descr": "BDS.WebShell.Chopper.1"
    }
]
```

