# webshell-scanner

[![Build Status](https://img.shields.io/badge/README-切换语言-yellow.svg)](readme-zh_CN.md)

## Introduction

API client of [https://scanner.baidu.com](https://scanner.baidu.com) implemented in Golang.

## List of supported file formats

| Language   | File extensions            |
| ---------- | -------------------------- |
| PHP        | .php, .phtml, .php[345], .inc |
| ASP        | .asp, .asa, .cer, .cdx, .asa   |
| ASPX       | .aspx, .asmx, .ashx          |
| JSP        | .jsp, .jspx .war            |
| Archive    | .rar, .zip, .tar, .xz, .tbz, .tgz, .tbz2, .bz2, .gz   |

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
  "error":  false,
  "status": "pending",
  "hash":   "de7128e140ac0add05d1728ab95d9d5d",
  "url":    "https://scanner.baidu.com/result/de7128e140ac0add05d1728ab95d9d5d"
}
```

## Result API

Fetch the detection result from previous submits

```
curl https://scanner.baidu.com/result/de7128e140ac0add05d1728ab95d9d5d
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
        // relative path (useful for archives)
        "path": "/de7128e140ac0add05d1728ab95d9d5d.php",

        // produced verdict
        "descr": "BDS.WebShell.Chopper.1"
    }
]
```

