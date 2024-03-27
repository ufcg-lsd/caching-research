#!/bin/bash

input_directory=""

for input_file in $input_directory/*input.csv; do
    output_file="${input_file%????}_reqs.csv"
    cut -d ',' -f 2 "$input_file" | tail -n +2 > "$output_file"
    echo "Processado $input_file => $output_file"
done