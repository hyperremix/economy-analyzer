#!/bin/bash

mongoexport --pretty -d ea -c classifications -o data/classifications.json
mongoexport --pretty -d ea -c transactions -o data/transactions.json