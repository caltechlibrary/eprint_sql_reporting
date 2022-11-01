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
    SQL='SELECT eprintid, IFNULL(title, "") AS title, IFNULL(doi, "") AS doi, username, email, CONCAT(name_family, ", ", name_given) AS name, eprint_status FROM eprint JOIN user ON (eprint.userid = user.userid) ORDER BY eprint_status, eprint.userid, eprintid'
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

