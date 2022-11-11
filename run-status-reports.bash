#!/bin/bash

#
# Run the status reports for all the EPrints repositories.
#
for REPO_ID in caltechauthors caltechcampuspubs caltechconf calteches caltechln caltechoh caltechthesis; do
    ./rpt-eprint-status.bash $REPO_ID \
        >"$REPO_ID-eprint-status.tsv"
    ./rpt-status-counts-by-user.py \
        "$REPO_ID-eprint-status.tsv" \
        "$REPO_ID-eprint-status-by-user.csv"
    ./rpt-status-counts-by-reviewer.py \
        "$REPO_ID-eprint-status.tsv" \
        "$REPO_ID-eprint-status-by-reviewer.csv"
done
