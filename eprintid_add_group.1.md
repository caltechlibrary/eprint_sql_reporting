% eprintid_add_group(1) user-manual
% R. S. Doiel
% 2022-10-26

# NAME

eprintid_add_group

# SYNOPSIS

eprintid_add_group EPRINT_ID_LIST GROUP_NAME

# DESCRIPTION

The program will generate a SQL statements to add a local_group
for a set of eprintid.

The EPRINT_ID_LIST should have only eprintid one per line. The
local group name applied is provided in the command line (i.e.
GROU_NAME above)

# OPTIONS

-help
: display this help page.

-license
: display software license

-version
: display version


# EXAMPLE

Process the eprint id list and the group name "Astronomy Department"
updated the eprint table appropriately.

~~~
eprintid_add_group eprint_ids.txt "Astronomy Department" \
    > eprintid_add_group.sql

mysql caltechauthors < eprintid_add_group.sql
~~~


eprintid_add_group 0.0.0
