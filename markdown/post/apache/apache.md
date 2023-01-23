---
date: "209-12-08"
draft: false
lastmod: "2019-12-08"
publishdate: "2019-12-08"
tags:
- linux
- System-Management-2
- OYK2018
- web-server
title: Apache
---

# APACHE


In the last days of the course, I will try to convey the web topic told by our climax teacher.

Doruk Hodja explained the subject on the slide. He then gave us some tasks and asked us to do them. I plan to tell what I have left in my mind from my own notes and teacher. If you see anything wrong, please do not forget to report it.


+ One centos minimal virtual machine is required in the course.
+ To change the errors in the narration: [Gitlab](https://gitlab.com/rection/lyk18-SistemYonetim-2-duzey.git)
+ For the slide used in the lesson: [Slide](http://topluluk.ozguryazilim.com.tr/wp-content/sunumlar/sistemyonetim-sunumlar/apache_web_sunücü.html#1)

## Web Architecture


{{<figure src = "images / apache / 0.png" title = "image title">}}


Clients in the photo are web browsers. Web browsers are an interface to access a page published on the Internet. If we give examples of these, such as firefoxi chrome, safari are vehicles. When you search for any address from the browser, you go to the server where the address is located via networks. The server must also be a web server to request a page from the server.

The web server is to check whether the page is requested and give the page. If the page was not found, the error code is returned. If the page is found, if there is a different operation requested (registration, data processing ...), it is directed to the application server. The application server is the mechanism that directs and controls the incoming request to perform the operations in it. Examples of the application server are tools such as zend (php), J2EE (java). If the database is used in the system, the application server sends and processes requests to the database. Communication between database and application server is done with http protocol.


HTTP: Hyper Text Transfer Protocol, port 80 is used.
HTTPS: HTTP Secure, port 443 is used.
URL: The part of the location of a web server that is read by people.

According to the number given by Netcraft (September 2018) web server shares:

![Alt Img](https://raw.githubusercontent.com/rection/LYK18-GNULinuxSistemYonetim-Duzey2/resimduzenleme/upgrade/katkida-bulunanlar/safa-bayar/apache/images/1.png)


According to the number provided by StatCounter (August 2018) client shares:

![Alt Img](https://raw.githubusercontent.com/rection/LYK18-GNULinuxSistemYonetim-Duzey2/resimduzenleme/upgrade/katkida-bulunanlar/safa-bayar/apache/images/2.png)


### Apache History

1991 - HTTP's 0.9 was announced.
1994 - NCSA HTTPd development stopped.
1995 - "A Patchy Server" based on NCSA.
1996 - It was named Apache in time and became the most popular server in the world.
1999 - Apache Software Foundation was established.
2018 - Apache is the world's most popular server with 350+ Apache projects.

### Apache Setup

Installing Apache is quite simple. It is explained through the distribution of Centos 7. In Debian distributions, the location of commands and configuration files change.

> `# yum install httpd`

It downloads and installs the httpd package.

> `# service httpd start`

We run the service after installation.

> `# # chkconfig httpd on`

We apply this command to avoid rebooting at every boot.

> `# sudo service firewalld start`

We start the firewalld service to allow external connection.

> `# sudo firewall-cmd --permanent --add-port = 80 / tcp`

We open 80th port for tcp protocol via firewall.


The settings files are as follows;

``> /etc/httpd/conf/httpd.conf ``
``> /etc/httpd/conf.d / *. conf ``
``> / etc / sysconfig / httpd ''

### Working Models

MPM (Multi-Processing Modules): Used to change the basic functions of the web server. This is thanks to the modular design of apache. It is used in all http operations. There are three selection stages to use these stages. These are Prefork, Worker and Event.

> ** Prefork (default): ** Apache starts with sub processes when it is first run. These are the services that work to meet the default requests. Prefork allows this to be checked.

> ** Worker (Threaded): ** Each subprocess creates new queues when there is a connection. A different subprocess is created when the difference from Prefork is the link here.

> ** Event (Threaded - Apache 2.4 +): ** Allows more requests to be served at the same time than other processes. Here, the client can keep the connection open after completing the first request. In this way, it can send more requests using the same socket and reduce the overload of the connection.

By default, Nginx is a Worker (Threaded) running service. Apache runs prefork by default. We can pass both web servers to one of the stages described.

### Modular Structure:
Many additional functions are added to the httpd core. Httpd core enables to manage the modular structure. The best feature of Apache is that it is in module format. Because the modules must be read each time, the less modules are added, the more they work. The modules can be loaded dynamically as well as embedded statically. Static is considered to be an advantage, but when any changes are made, it needs to be recompiled. This causes the system to slow down.

** `# # http -l` ** command is used to see the static compiled modules.

Adding / Removing Modules:

There is `LoadModule auth_basic_module modules / mod_auth_basic.so` in line 51 in /etc/httpd/conf/httpd.conf. 'modules' specifies the / etc / httpd / Modules directory. The module you want to add must be put in the modules directory.

Some of Apache's own modules are:

+ mod_dav: With HTTP protocol, remote server enables operations such as creating, moving and deleting resources.
+ mod_deflate: It provides compression on the server side.
+ mod_ldap: It is the module used to improve the performance of the LDAP service running at the back.

Some of Apache's common additional modules are:

+ mod_pagespeed: Automatically provides web optimization.
+ mod_security: It has a firewall system for itself. It can also be used for vulnerabilities such as sql injection.
+ mod_wsgi: Used in the settings and presentation of the application written in Python language.

## Basic Settings:


The /etc/httpd/conf/httpd.conf file has the ** ServerRoot ** setting on line 30. Indicates where the configuration files are located.

In the /etc/httpd/conf/httpd.conf file, there is a ** DocumentRoot ** setting on line 109. Specifies the location of the file to be presented. It is usually / var / www / html.

In the /etc/httpd/conf/httpd.conf file, there is a ** ServerName ** setting on line 88. This is where the private hostname and port number are entered. There is no obligation to be a hostname. There may also be an address registered on the domain name service, such as www.exapmle.com.


In the /etc/httpd/conf/httpd.conf file, there is a ** Listen ** setting on line 40. It specifies which ip address and port to serve. supports ipv6.

** ServerAlias ​​**, which is used when defining Virtualhost, is to define more than one domain name.

## Directory Specific Settings:

First, I will talk about the use of single directories.

> `` <Directory /> `` ''
> `` Options FollowSymLinks ''
> `` AllowOverride None`` ''
> `` </ Directory> ```


** Options FollowSymLinks: ** It comes by default. It means that the server follows symbolic links in the directory. For more information [click](https://httpd.apache.org/docs/2.4/mod/core.html#Options%20Directive).

** AllowOverride None: ** Refers to refusing to read files in the directory system. It is important for *** .htaccess *** files. Other methods that can be applied are authentication (AuthConfig), File information (FileInfo) ... [click] for details (https://httpd.apache.org/docs/2.4/mod/core.html#allowoverride).

Multiple directories must also be set as follows for use.

> `` <DirectoryMatch / [a-d]. *> ``
> `` Options -FollowSymLinks ''
> `` AllowOverride All``
> `` </ DirectoryMatch> ```

It means to apply these rules for the letters between a and d that are intended to be explained in the '/uthentica-d...*' section defined in DirectoryMatch, and for subsequent directory mappings.
** - '' option in Options -FollowSymLinks ** means that this feature is turned off. + Comes by default.
** AllowOverride All ** indicates that it allows reading files.

Similarly, there are *** file, url, proxy and virtualhost *** specific settings. There are some important settings.

** Order allow, deny: ** *** Indicates *** Reject rules *** are processed before the allow rules ***. If the client does not match *** Permission rules *** or if there is a matching *** Deny rule *** then the client will deny access.

** Order deny, allow: ** *** Specifies that *** allow rules *** will be processed before Deny rules ***. If the client does not match the *** Deny rule ***, or if there is a matching *** Permission rule ***, the client will allow access.

It is obligatory to write Deny and Allow lines under this line. For more information [click](http://www.maxi-pedia.com/Order+allow+deny).

## DirectoryIndex
It also enables access to directories and error codes to be adjusted accordingly. If the defined file does not match, the Indexes property is examined. If there is an Indexes authorization, the file list is displayed. Returns "403 Forbidden" if Indexes is not authorized.

** ServerTokens: ** Information that is carried at the beginning of the server at the beginning of the response returned by the server. It shows the information given according to the version of Apache and Operating system. For more information [click](https://httpd.apache.org/docs/2.4/mod/core.html#servertokens).

** StartServers 10: ** By default, open 10 processes and wait for the request to arrive. [Click] for more information (https://httpd.apache.org/docs/2.4/mod/mpm_common.html#startservers).

** MinSpareServers 5: ** Determines the minimum number of idle subprocesses. The idle process is a process that does not provide a service at the moment. If MinSpareServers has fewer idle processes than the number of processes, the main process will create new child processes to complete the number. [Click] to learn more (https://httpd.apache.org/docs/2.4/en/mod/prefork.html).

** MaxSpareServers 20: ** Determines the maximum number of idle subprocesses. The idle process is a process that does not provide a service at the moment. If MaxSpareServers has more than one idle process, the main process will kill those surpluses. [Click] to learn more (https://httpd.apache.org/docs/2.4/en/mod/prefork.html).

** MaxRequestsPerChild 4000: ** Specifies how many transactions will kill the transaction. [Click] for more information (https://httpd.apache.org/docs/2.4/mod/mpm_common.html#MaxRequestsPerChild).

** MaxClients 150: ** It can be adjusted according to the maximum amount of clients according to the amount of ram. [Click] for more information (https://2bits.com/articles/tuning-the-apache-maxclients-parameter.html).

** ServerLimit 256: ** Determines the maximum amount of ram Apache processes can spend. It should be adjusted proportionally with MaxRequestWorkers. [Click] to learn more (https://httpd.apache.org/docs/current/en/mod/mpm_common.html).

** KeepAlive On: ** Allows to receive requests in a row. [Click] for more information (https://httpd.apache.org/docs/2.4/mod/core.html#KeepAlive).

** KeepAliveTimeout 15: ** Specifies how many seconds to wait for the next request before the server closes a permanent connection. The speed of the users must be checked and adjusted. [Click] to learn more (https://httpd.apache.org/docs/2.4/en/mod/core.html).

** MaxKeepAliveRequests 4000: ** Limits the number of requests allowed per connection. To keep server performance high, it should be a high value. For more information [click](https://httpd.apache.org/docs/2.4/mod/core.html#maxkeepaliverequests).

## .htaccess
.htaccesss is a tool that allows you to make directory-specific settings where the data to be presented is located. The .htaccess file should be used when the host server configuration file cannot be accessed or changed. It is also used to redirect from old site to new site or from old pages to new pages.

The AccessFileName setting must be enabled in + httpd.conf.
> AccessFileName .htaccess
+ Settings made in httpd.conf change based on directory.
+ Apache first looks for the .htaccess file in each directory.
These settings apply to the + O and below directories.

#### Advantages:
It does not require authorization to edit + httpd.conf.
+ Settings take effect immediately.
+ Settings are also moved when moving the directory.

#### Disadvantages:
+ Lowers performance. When a document is requested, it should look up to the root directory of each directory and it reads the first .htaccess file again at the stage of changing each directory. This causes it to slow down.
+ Reduces security. Incorrect configuration of the directives in their files, and documents in the directory and subdirectories can also cause problems.

#### Alias
It makes it possible to store the directory where the files will be presented in a different place. It provides mapping of url path and directory path.

#### Redirect
Used to redirect an old url. It is a method used to force https. While regex cannot be used in Redirect, it can be used with * RedirectMatch *.

> `` Redirect 301 / old_place.html http: // www.ozguryazilim.com.tr / new_place.php`` ''
> `` RedirectMatch ^ / images /(.*)$ http: // images.ozguryazilim.com.tr / $ 1`` ''

* Our teacher explained the rewrite topic on the slide. Anyway, the slide is sufficiently illustrative about this topic. *

## Encryption with SSL:

Its expansion is Secure Sockets Layer. Its purpose is to establish a secure connection between a website and the visitor's browser. In other words, it is to verify that the other party is the place we want to reach. For SSL domain name. Not for the server. [Click] to learn more (http://www.networksolutions.com/education/what-is-an-ssl-certificate/).

We are not talking about the SSL setup phase because it is one of the tasks given in the last part of the lesson.

If the file is to be uploaded, it is first downloaded under / tmp / directory. It is then transferred. Their definition is used in php provisioning.

## Proxy:

With a simple definition, users are intermediate servers during internet access. It can only be adjusted via the browser. In a connection, your request is first opened to the proxy server and then to the internet. Its purpose is to enter banned sites, but it is not like vpn. The connection between you and the proxy server is not encrypted and the proxy server may not support https. It has two types.

![Proxy](https://raw.githubusercontent.com/rection/lyk18-gnulinuxsistemyonetimi-duzey2/resimduzenleme/upgrade/katkida-bulunanlar/safa-bayar/apache/images/3.png)

#### Forward Proxy (Redirected proxy):
The subject discussed above is forward proxy.

#### Reverse Proxy:
It is the reverse of the forward proxy. It meets the requests received by the client and directs it to the server located behind.

The difference between the two is: While the forward proxy hides the client's identity, Reverse allows the proxy server to hide its identity.

With the reverse proxy, we can remove the SSL load on the server side. In this way, the load of the server is reduced. To debug and capture, server side http should be used.

## Apache Logs

By default, log records are kept under / var / log / httpd /. It is possible to keep it under the name of the client upon request. Logs must be compressed and signed at the end of the day. Logs must be kept for 2 years in order to avoid any problem in the face of any legal situation. In case of not keeping the logs, the culprit is the person who should keep a record. For detailed information [click](http://www.mevzuat.gov.tr/MevzuatMetin/1.5.5651.pdf).

Some tools were mentioned to analyze the logs. Of these, piwik is used the most.


Homework given by #### Doruk Hoca:
1) SSL
2-) It is requested to refer to Reverse Proxy -----> hürriyet.com.tr.
3-) Virtualhost -------> test1.linux.gen.tr
               |
               -------> is test2.linux.gen.t
4-) It is requested to create a password area.
5-) PHP-fpm setup -----> reverse proxy ----> phpinfo ()
6-) IP setup

The results of the homework will be announced in the following days. ***The under construction***
