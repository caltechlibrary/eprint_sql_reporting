% dump-eprints-records.bash(1) user manual
% R. S. Doiel
% 2022-10-26

# NAME

dump-eprints-records.bash

# SYNOPSIS

dump-eprints-records.bash REPOSITORY_ID

# DESCRIPTION

This script dumps an EPrints repository including only the tables useful
for development externally to EPrints. It explicitly excludes tables only
needed for EPrints internal functionality, e.g. any table name containing
triple, access, irstat, sort, "__", index, cache, shelf, shelve, saved
or request. These dumps are NOT BACKUPS!

It saves the files in a "sql-dumps" directory with the name form of
REPOSITORY_ID dash dump dash date string period "sql".

# EXAMPLE

Running the command on Sept. 9, 2021

~~~
	dump-eprints-records.bash caltechauthors
~~~

This would generate a dump for development use as

~~~
    sql-dumps/caltechchauthors-2021-09-29.sql.gz
~~~

