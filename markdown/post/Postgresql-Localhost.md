---
date: "209-12-05"
draft: false
lastmod: "2019-12-05"
publishdate: "2019-12-05"
tags:
- linux
- postgresql
- postgres
- postgresql-ident
title: Postgresql Ident Authentication Failed
---

System: GNU/Linux  
Distribution: Centos 7  
Process: Install redmine  


I had installing redmine. Redmine want access postgresql database. I applied this command :  
``` 
RAILS_ENV=production bundle exec rake db:migrate
```  

but it's give a error:  
``` 
FATAL: Ident authentication failed for user "<database-name>"
```  

I checked /var/lib/pgsql/11/data/pg_hba.conf, i corrected these lines:
```
# "local" is for Unix domain socket connections only
local   all             all                                     trust
# IPv4 local connections:
host    all             all             127.0.0.1/32           	trust
```
but it give same error thereafter i search. I found some think, postgresql used ipv6 when work on localhost. I correct this line:
```
# IPv6 local connections:
host    all             all             ::1/128                 trust
```
and I took back last changes. Consequently it worked at this stage, so when postgresql works on localhost it used ipv6.
