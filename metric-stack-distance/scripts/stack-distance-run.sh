#!/bin/bash

input_directory=""
program="./stack-distance"

for input_file in $input_directory/*_reqs.csv; do
    output_file="${input_file%????}_output.csv"
    $program < "$input_file" > "$output_file"
    echo "Processed $input_file => $output_file"
done
