
Overview
========

This repository holds a collection of command line programs (written in Go) and a set of Bash scripts for working with EPrints 3.x MySQL database tables. In some cases the Bash scripts may include the use of tools from the [datatools project](https://github.com/caltechlibrary/datatools/latest/release). Bash scripts that interact with MySQL 8 do so via the `mysql` client.  The cli provided tend to generate SQL documents that can be run in the MySQL 8 client. The output from the client is then processed with common POSIX utilities (e.g. grep, sort, sed) and may service as input for the next stage of processing.

