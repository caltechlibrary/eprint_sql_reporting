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

Generate a table of eprintid, eprint_status and date eprint was
created.

# EXAMPLE

~~~
    ${APP_NAME} EPRINT_REPO_ID > eprintids_by_created.tsv
~~~

EOT

}

function build_report() {
    SQL='SELECT eprintid, eprint_status, CONCAT(LPAD(datestamp_year, 4, "0"), "-", LPAD(datestamp_month, 2, "0"), "-", LPAD(datestamp_day, 2, "0")) AS created FROM eprint ORDER BY eprintid;'
    #echo "Generating a tab separated file"
    mysql "$1" --batch --execute "${SQL}"
}

#
# Main processing
#
for ARG in "$@"; do
    case $ARG in
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
	echo "expected a EPrints REPO_ID"
	exit 1
fi
build_report "$REPO_ID"

