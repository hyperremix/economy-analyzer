#!/bin/bash

SOURCE="$( dirname "${BASH_SOURCE[0]}" )"

mongoexport --pretty -d ea -c classifications -o $SOURCE/data/classifications.json
mongoexport --pretty -d ea -c transactions -o $SOURCE/data/transactions.json