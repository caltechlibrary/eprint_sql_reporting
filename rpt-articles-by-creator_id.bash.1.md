% rpt-articles-by-creator_id.bash(1) user manual
% R. S. Doiel
% 2022-10-26

# NAME

rpt-articles-by-creator_id.bash

# SYNOPSIS

rpt-articles-by-creator_id.bash CSV_FILE COLUMN_NO_OF_AUTHOR_ID

# DESCRIPTION

This script uses both datatools and the cli from
eprint_sql_reporting to generate a report of journal articles
my the MySQL database for an EPrints 3.x repository.

# EXAMPLE

~~~
    rpt-articles-by-creator_id.bash "GPS_Faculty_ORCIDS - Sheet1.csv" 1         >GPS_Fuculty_Articles.csv
~~~


