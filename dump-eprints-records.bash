#!/bin/bash

function usage() {
    APP_NAME=$(basename "$0")
    cat <<EOF
% ${APP_NAME}(1) user manual
% R. S. Doiel
% 2022-10-26

# NAME

${APP_NAME}

# SYNOPSIS

${APP_NAME} REPOSITORY_ID

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
    ${APP_NAME} caltechauthors
~~~

This would generate a dump for development use as

~~~
    sql-dumps/caltechchauthors-2021-09-29.sql.gz
~~~

EOF

    exit 1
}

#
# Main Processing
#

# Handle command line parameter
case "$1" in
"" | "-h" | "--help")
    usage
    ;;
*)
    REPO_ID="$1"
    ;;
esac

#
# dump-caltechauthors.bash dumps the tables that are most useful as test data.
# notably absent are all access*, cache*, and counter tables.
#

EXT="_$(date +"%Y-%m-%d").sql"

if [ ! -d sql-dumps ]; then
    mkdir sql-dumps
fi
FNAME="sql-dumps/${REPO_ID}-dump${EXT}"

TABLE_NAMES=$(mysql --batch "${REPO_ID}" --execute "SHOW TABLES" | grep -v -E "Tables|triple|access|irstat|sort|__|index|cache|shelf|shelve|request")
#TABLE_NAMES="${TABLE_NAMES} eprint__rindex"

if [[ -f "${FNAME}" ]]; then
    rm "${FNAME}"
fi
touch "${FNAME}"
for T in ${TABLE_NAMES}; do
    echo "Dumping table ${T} to ${FNAME}"
    mysqldump --add-drop-table "${REPO_ID}" "${T}" >>"${FNAME}"
done
