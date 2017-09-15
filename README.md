Nginx Proxying
==============

Proxying to multiple app servers based on `Location`.

Boot containers and tail logs:

```
% make
```

Make HTTP requests to paths:

```
% curl -H "Host: myhost" http://localhost:8888/
% curl -H "Host: myhost" http://localhost:8888/foo
% curl -H "Host: myhost" http://localhost:8888/foo/bar
% curl -H "Host: myhost" http://localhost:8888/foo/bar/hoge
```

Stop containers:

```
% make stop
```

Clean containers:

```
% make clean
```
