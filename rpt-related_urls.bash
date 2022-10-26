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

${APP_NAME}

# DESCRIPTION

Generate a tab delimited file that reports all the eprintid and
their related url.

# EXAMPLE

~~~
    ${APP_NAME} > eprint_related_urls.tsv
~~~

EOT

}

function build_report() {
    SQL='SELECT eprint_related_url_url.eprintid AS eprintid, eprint_related_url_url.pos AS pos, related_url_url, related_url_type, related_url_description  FROM eprint_related_url_url JOIN (eprint_related_url_type, eprint_related_url_description) ON (eprint_related_url_url.eprintid = eprint_related_url_type.eprintid) AND (eprint_related_url_url.pos = eprint_related_url_type.pos) AND (eprint_related_url_url.eprintid = eprint_related_url_description.eprintid) AND (eprint_related_url_url.pos = eprint_related_url_description.pos) WHERE related_url_url IS NOT NULL AND CHAR_LENGTH(TRIM(related_url_url)) > 0'
    #echo "Generating a tab separated file"
    mysql caltechauthors --batch --execute "${SQL}"
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

build_report

