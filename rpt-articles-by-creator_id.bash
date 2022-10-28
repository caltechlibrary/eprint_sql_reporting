#!/bin/bash

function usage() {
    APP_NAME=$(basename "$0")

cat <<EOT
% ${APP_NAME}(1) user manual
% R. S. Doiel
% 2022-10-26

# NAME

${APP_NAME}

# SYNOPSIS

${APP_NAME} CSV_FILE COLUMN_NO_OF_AUTHOR_ID

# DESCRIPTION

This script uses both datatools and the cli from
eprint_sql_reporting to generate a report of journal articles
my the MySQL database for an EPrints 3.x repository.

# EXAMPLE

~~~
    ${APP_NAME} "GPS_Faculty_ORCIDS - Sheet1.csv" 1
~~~

The resulting report is "eprintid_pub_type.csv" and
"eprintid_pub_type.tsv".

EOT

}

function build_report() {
    CSV_FILE="$1"
    COLUMN_NO="$2"
    if [ ! -f eprint_ids.txt ]; then
        echo "Building creator_ids.sql file"
        creatorid_to_eprintid "$CSV_FILE" "$COLUMN_NO" \
            >creator_ids.sql
        echo "Generating a sorted unique eprint_ids.txt"
        mysql caltechauthors --batch --skip-column-names \
            < creator_ids.sql | sort -n -u > eprint_ids.txt
    fi
    echo "Generating eprintid_pub_type.sql"
    eprintid_pub_type eprint_ids.txt >eprintid_pub_type.sql
    echo "Generating csv of eprintid and publication type"
    printf '"eprintid","title","doi","publication_type","date_type","date"\n' >eprintid_pub_type.csv
    mysql caltechauthors --batch --skip-column-names \
            < eprintid_pub_type.sql | \
			sed -E "s/'/\'/;s/\t/\",\"/g;s/^/\"/;s/$/\"/;s/\n//g" \
			>> eprintid_pub_type.csv
    grep -E "2022|2021|2020|2019|2018|2017|2016|2015|2014|2013|2012" \
		eprintid_pub_type.csv | grep ',"article",' | wc -l

    printf "eprintid\ttitle\tdoi\tpublication_type\tdate_type\tdate\n" >eprintid_pub_type.tsv
    mysql caltechauthors --batch --skip-column-names \
            < eprintid_pub_type.sql >> eprintid_pub_type.tsv
    grep -E "2022|2021|2020|2019|2018|2017|2016|2015|2014|2013|2012" eprintid_pub_type.tsv | grep "\tarticle\t" | wc -l
}

#
# Main processing
#
if [ "$#" = "1" ]; then
    usage
    exit 1
fi
for ARG in "$@"; do
    case ARG in
        -h|-help|--help)
        usage
        exit 0
        ;;
    esac
done

if [ "$#" != 2 ]; then
    echo "Missing the CSV filename or column number holding the author id"
    exit 1
fi

build_report "$1" "$2"

