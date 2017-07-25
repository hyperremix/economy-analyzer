#!/bin/bash

mongoimport --drop -d ea -c classifications --file data/classifications.json
mongoimport --drop -d ea -c transactions --file data/transactions.json