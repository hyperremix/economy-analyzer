#!/bin/bash

SOURCE="$( dirname "${BASH_SOURCE[0]}" )"

mongoimport --drop -d ea -c classifications --file $SOURCE/data/classifications.json
mongoimport --drop -d ea -c transactions --file $SOURCE/data/transactions.json