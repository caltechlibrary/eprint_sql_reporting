INSTALL
=======

This is a collection of experimental cli. It is only distributed in source code. If you wish to compiler and test them you need the required development environment and follow the steps listed below in "Compiling from Source".

Requirements
------------

Bottler is currently implemented as a Go application. This may change in the future.

- Git to clone the repository
- Compiling the cli
    - [Golang](https://golang.org) 1.19.2 or better
    - GNU Make
    - Pandoc 2.10 or better (to build documentation)

Compiling from Source
---------------------

1. clone the repository
2. change into the cloned directory
3. run "make", "make install" to install the cli, Bash scripts and man pages

Here's the steps I take on my macOS box or Linux box.

~~~
git clone git@github.com:caltechlibrary/eprint_sql_reports.git
cd eprint_sql_reports
make
make install
~~~

