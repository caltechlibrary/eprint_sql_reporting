% eprintid_pub_types(1) user-manual
% R. S. Doiel
% 2022-10-26

# NAME

eprintid_pub_types

# SYNOPSIS

eprintid_pub_types EPRINT_ID_LIST

# DESCRIPTION

The program will generate a SQL statements to produce a list of
eprint records containing eprintid, title, doi, publication type
and pub date. It is suitable to process via the mysql client and
rendering a tab delimited file result.

The EPRINT_ID_LIST should have only eprintid one per line.

# OPTIONS

-help
: display this help page.

# EXAMPLE

Process the eprint id list and produce a report of eprintid containing
their title, doi, publication type and publication date.

~~~
eprintid_pub_types eprint_id_list.txt \
    > eprintid_report.sql

mysql caltechauthors --batch < eprintid_pub_type_report.sql | \
   > eprintid_pub_type_report.tsv
grep '\tarticle\t' <eprint_pub_type_report.tsv| wc -l
~~~


