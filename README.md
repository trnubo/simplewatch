## simplewatch

Similar to the linux watch command, but for doing checks in the background.
Simply calls the defined command each x seconds and if the command exits
non-zero it exits with the same exit code.

Altough this could have been done in a couple of lines with bash I thought I'd
give golang a try.

### Usage

```
# ./simplewatch -h
Usage of ./simplewatch:
  -firstwait=30: wait in seconds before first check
  -wait=30: wait in seconds between checks
```

### Example

```
root@62fc98952d7b:/# ./simplewatch -firstwait 2 -wait 10 /usr/lib/nagios/plugins/check_http -H 127.0.0.1 -p 8000
2015/06/15 01:55:26 Starting simplewatch. First Wait:  2  Wait:  10  CMD:  [/usr/lib/nagios/plugins/check_http -H 127.0.0.1 -p 8000]
2015/06/15 01:55:28 ==> Output:  HTTP OK: HTTP/1.0 200 OK - 965 bytes in 0.002 second response time |time=0.001914s;;;0.000000 size=965B;;;0
2015/06/15 01:55:38 ==> Output:  HTTP OK: HTTP/1.0 200 OK - 965 bytes in 0.002 second response time |time=0.001632s;;;0.000000 size=965B;;;0
2015/06/15 01:55:48 ==> Output:  HTTP OK: HTTP/1.0 200 OK - 965 bytes in 0.001 second response time |time=0.001440s;;;0.000000 size=965B;;;0
2015/06/15 01:55:58 ==> Output:  HTTP OK: HTTP/1.0 200 OK - 965 bytes in 0.002 second response time |time=0.001582s;;;0.000000 size=965B;;;0
2015/06/15 01:56:08 ==> Output:  HTTP OK: HTTP/1.0 200 OK - 965 bytes in 0.001 second response time |time=0.001457s;;;0.000000 size=965B;;;0
2015/06/15 01:56:18 ==> Output:  HTTP OK: HTTP/1.0 200 OK - 965 bytes in 0.002 second response time |time=0.001529s;;;0.000000 size=965B;;;0
2015/06/15 01:56:28 ==> Output:  Connection refused
HTTP CRITICAL - Unable to open TCP socket
2015/06/15 01:56:28 Exit Status: 2
2015/06/15 01:56:28 Done. Exit:  2
root@62fc98952d7b:/# echo $?
2
```

### License

The MIT License (MIT)

Copyright Tim Robinson <tim@voltgrid.com>

Some code for getting exit codes is from http://stackoverflow.com/a/10385867
