#!/bin/bash

function usage() {
    APP_NAME=$(basename "$0")

cat <<EOT
% ${APP_NAME}(1) user manual
% R. S. Doiel
% 2022-11-10

# NAME

${APP_NAME}

# SYNOPSIS

${APP_NAME} EPRINT_REPO_ID

# DESCRIPTION

Generate a table of eprint record with eprintid, item type, journal
or book title, isbn, and issn. 

# EXAMPLE

~~~
    ${APP_NAME} EPRINT_REPO_ID > eprint_isbn_issn.tsv
~~~

EOT

}

function build_report() {
    SQL='SELECT eprintid, IFNULL(eprint_status, "") AS eprint_status, IFNULL(type, "") AS eprint_type, IFNULL(title, "") AS title, IFNULL(publication, "") AS journal, IFNULL(publisher, "") AS publisher, IFNULL(book_title, "") book_title, IFNULL(issn, "") AS issn, IFNULL(isbn, "") AS isbn FROM eprint ORDER BY eprint_status, eprint_type, publisher, publication, title'
    #echo "Generating a tab separated file"
    mysql "$1" --batch --execute "${SQL}"
}

#
# Main processing
#
for ARG in "$@"; do
    case ARG in
        -h|-help|--help)
        usage
        exit 0
        ;;
		"$0")
		# Ignore the app name
		;;
		*)
		REPO_ID="$ARG"
		;;
    esac
done
if [ "$REPO_ID" = "" ]; then
	usage
	echo "expected EPrints REPO_ID"
	exit 1
fi
build_report "$REPO_ID"

