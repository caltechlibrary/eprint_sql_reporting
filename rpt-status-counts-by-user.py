#!/usr/bin/env python3
import os
import sys
import csv

def usage(app_name):
    print(f'''% {app_name}(1) user manual
% R. S. Doiel
% 2022-11-01

# NAME

{app_name}

# SYNOPSIS

{app_name} TSV_INPUT_FILENAME CSV_REPORT_FILENAME

# DESCRIPTION

This takes the output in tab separated values form the
rpt-eprint-status.bash and processes the resulting
table reporting on counts for all the different
status values by user.

TSV_INPUT_FILENAME is the  name of the tab separated value
file generated from rpt-eprint-status.bash.

CSV_REPORT_FILENAME is the name given to the generated
report based on processinf the TSV_INPUT_FILENAME.

# OPTIONS

-h, -help, --help
: display this help page

# EXAMPLE

```
rpt-eprint-status.bash caltechthesis \\
    >caltechthesis-eprint-status.tsv
{app_name} caltechthesis-eprint-status.tsv \\
    caltechthesis-eprint-status-counts-by-user.csv
```

''')


# process_tsv is an example of iterating through the tab
# separated value file one row at a time and then counting
# things. I am using a dictionary to hold the users and then
# where the value itself is another map of username, email,
# given_name, family_name, eprint_status. The eprint_status
# element is itself another dictionary holding the status value
# (e.g. inbox, archive, deletion) pointing at the eprintid having that
# status. The procedure returns a touple of the dictionary holding
# the aggragation and an error value (string). If an error occurs then
# the second element of the returned touple is a non-empty string.
#
# ```
#     data, err = process_tsv('caltechauthors-eprint-status.tsv')
#     if err != '':
#         printf(f'ERROR: {err}', file=os.stderr)
#          sys.exit(1)
# ```
# The "report" then just counts things up and writes them to the screen.
#
def process_tsv(tsv_filename):
    '''process the tsv file and return a dict of dict with aggregated data'''
    data = {}
    with open(tsv_filename, newline = '', ) as csvfile:
        reader = csv.DictReader(csvfile, dialect = 'excel-tab')
        for i, row in enumerate(reader):
            if 'username' in row:
                username = row['username']
                if not username in data:
                    data[username] = { 'email': row['email'], 'name': row['name'] }
                    data[username]['eprint'] = []
                data[username]['eprint'].append({'eprintid': row['eprintid'], 'eprint_status': row['eprint_status']})
            else:
                return None
    return data

# report takes our dictionary of the aggregated information in your tsv file and
# renders the report
def report(csv_filename, data):
    fieldnames =  ["username", "email", "name", "eprint_status", "count"]
    with open(csv_filename, 'w', newline = '') as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames = fieldnames)
        writer.writeheader()
        for key, val in data.items():
            # Setup identifying elements in the row
            obj = {'username': key, 'email': val['email'], 'name': val['name']}
            # Now run out aggregation
            for status in [ 'NULL', 'archive', 'buffer', 'deletion', 'inbox' ]:
                # Count each status in the 'items' array
                cnt = 0
                for item in val['eprint']:
                    if item['eprint_status'] == status:
                        cnt += 1
                # Now er can output a row in the spreadsheet for that user with the status.
                obj['eprint_status'] = status
                obj['count'] = cnt
                writer.writerow(obj)

#
# Process the command line and run the report.
#
if __name__ == '__main__':
    app_name = os.path.basename(sys.argv[0])
    options = []
    args = []
    for arg in sys.argv[1:]:
        if arg.startswith('-'):
            options.append(arg)
        else:
            args.append(arg)
    if ('-h' in options) or ('-help' in options) or ('--help' in options):
        usage(app_name)
        sys.exit(0)
    if len(args) == 0:
        usage(app_name)
        sys.exit(1)
    if len(args) != 2:
        print(f'expected tsv_input_filename and tsv_output_filename, got {" ".join(sys.argv)}', file=os.stderr)
        sys.exit(1)

    data = process_tsv(args[0])
    report(args[1], data)
