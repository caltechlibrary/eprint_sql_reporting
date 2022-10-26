
eprint SQL reporting
====================

This repository holds a collection of Go cli (go programs you can run from the command line) and a set of Bash scripts for working with EPrints 3.x a the MySQL database level. In some cases the Bash scripts may include the use of tools from the [datatools project](https://github.com/caltechlibrary/datatools/latest/release). Bash scripts that interact with MySQL 8 do so via the `mysql` client.  The cli provided tend to generate SQL documents that can be run in the MySQL 8 client. The output from the client is then processed with common POSIX utilities (e.g. grep, sort, sed) and may service as input for the next stage of processing.

Requirements
------------

- go >= 1.19.2
- mysql 8 client
- make (e.g. GNU Make)
- Pandoc >= 2.19 (for generating HTML and man pages)
- bash

To run report examples you many also need to have the
[datatools](https://github.com/caltechlibrary/datatools/latest/release) cli installed.

More info
---------

Documentation provided in the [user manual](user-manual.html)

See [about](about.html) for credits

For questions see [GitHub Issues](https://github.com/caltechlibrary/eprint_sql_reporting/issues)

[License](LICENSE)

