% creatorid_to_eprintid(1) user-manual
% R. S. Doiel
% 2022-10-26

# NAME

creatorid_to_eprintid

# SYNOPSIS

creatorid_to_eprintid CSV_FILENAME COL_NUMBER_OF_ID

# DESCRIPTION

The program will generate a SQL statements for matching a spreadsheet
with a "Name" column to the author id in CaltechAUTHOR's eprint_creators_id
table.  This can be then passed via mysql client to generate a tab
delimited file if eprintid that next be passeed through the Unix sort
command to produce a list of unique eprint id for a group of authors.

The CSV_FILENAME should be the name of the file containing the author
ids.  COL_NUMBER_OF_ID is the column number (1 is initial column)
holding the author_id.

# OPTIONS

-help
: display this help page.

# EXAMPLE

CSV file called "GPS_Faculty_ORCIDS - Sheet1.csv", the author id is
in the column called "Name". Here's the steps to produce the report.

~~~
creatorid_to_eprintid "GPS_Faculty_ORCIDS - Sheet1.csv" 1 \
    > creator_ids.sql

mysql caltechauthors --batch --skip-column-names < creator_ids.sql | \
   sort -u >eprint_ids.txt
wc -l eprint_ids.txt
~~~

This leaves you author a list of eprint ids related to all the authors
in the CSV file.  You need to filter these one step further if you
are interested in publication type.

~~~
eprintid_pub_types eprint_ids.txt >eprint_pub_types.sql
mysql caltechauthors --batch --skip-column-names <eprintid_pub_types.sql |\
   >eprintid_pub_types.tsv
~~

This last step will produce a CSV file with one eprint record per line
with the eprintid, article title, doi, publicatin type, publication date.

creatorid_to_eprintid 0.0.0
