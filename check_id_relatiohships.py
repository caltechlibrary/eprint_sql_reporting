#!/usr/bin/env python3

import sys
import os
import csv

# Scan and record all the cl_people_id
# For each EPrint alternative people id (e.g. authors_id, thesis_id, and
# advisor_id) confirm they do don't conflict with cl_people_id unless they
# appropraite match (i.e. on the same row)

def find_key(l, k):
    for i, item in enumerate(l):
        if item == k:
            return i
    return -1                    

# cl_people_ids are unique in their own column we need to maintain
# the column/row relationship and use this list to check for collisions
# in mapping.
cl_people_ids = []
# The following dictionaries hold an id string and point to a row number
authors_creator_ids = {}
thesis_creator_ids = {}
thesis_advisor_ids = {}
with open("people.csv", newline = "") as csvfile:
    reader = csv.DictReader(csvfile)
    for i, row in  enumerate(reader):
        cl_people_ids.append(row['cl_people_id'])
        if row['authors_id'] != '':
            # Save the authors string and row value
            authors_creator_ids[row['authors_id']] = i
        if row['thesis_id'] != '':
            # Save the thesis string and row value
            thesis_creator_ids[row['thesis_id']] = i
        if row['advisor_id'] != '':
            thesis_advisor_ids[row['advisor_id']] = i

# Now see if the various ids collide with cl_people_id
# inappropriately

def check_maps(id_map, label):
    for key in id_map:
        i = id_map[key]
        pos = find_key(cl_people_ids, key)
        if (pos >= 0) and (pos != i):
            print(f'{pos+2}, {cl_people_ids[pos]}, {label}, {key},  {i+2},  {cl_people_ids[i]}')

print(f'row no., cl_people_id, conflicts with, value, conflicting row no., conflicting cl_people_id')
check_maps(authors_creator_ids, "authors_id")
check_maps(thesis_creator_ids, "thesis_id")
check_maps(thesis_advisor_ids, "advisor_id")

