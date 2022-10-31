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

${APP_NAME} EPRINT_REPO_ID

# DESCRIPTION

Generate a small table of counts for eprint_status for a given repository
database.

# EXAMPLE

~~~
    ${APP_NAME} EPRINT_REPO_ID > eprint_status_counts.tsv
~~~

EOT

}

function build_report() {
    SQL='SELECT eprint_status, COUNT(eprint_status) AS counts  FROM eprint GROUP BY eprint_status ORDER BY eprint_status'
    #echo "Generating a tab separated file"
    mysql caltechauthors --batch --execute "${SQL}"
}

#
# Main processing
#
if [ "$#" != "1" ]; then
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


build_report "$1"

