
Overview
========

This repository holds a collection of Go cli (go programs you can run from the command line) and a set of Bash scripts. In some cases the Bash scripts may include the use of tools from the [datatools project](https://github.com/caltechlibrary/datatools/latest/release). Bash scripts that interact with MySQL 8 do so via the `mysql` client.  The cli provided tend to generate SQL documents that can be run in the MySQL 8 client. The output from the client is then processed with common POSIX utilities (e.g. grep, sort, sed) and may service as input for the next stage of processing.

