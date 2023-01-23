---
date: "209-12-05"
draft: false
lastmod: "2019-12-05"
publishdate: "2019-12-05"
tags:
- linux
- find
- prune
title: Prune Parameter in Find Command
---

# PRUNE

Find command is typically used to search for a directory or file but i want to search except a directory. This solve different ways with find command but i did this way. I'll tell prune parameter.

``` 
# find / -path /run/media -prune -print -o -print 
```

You are know the structure of the find command.

+ / : Root directory.
+ -path : I want search  outside directory. This place want directory name.
+ -prune : Prune parameter should make search outside previous directory name. we have already purposed it.
+ -print : Display search result
+ -o : It's 'or', logical expression. Man page say:  
        ``` 
		expr1 -o expr2
             	Or; expr2 is not evaluated if expr1 is true.
        ```

If you want search a file or directory name, you must use grep command. Grep command  search a word and you are may use regex. I used this type:

``` 
# find / -path /run/media -prune -print -o -print  | grep -i 'meminfo' 
```
