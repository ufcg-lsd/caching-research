#!/bin/bash

input_directory=""
output_directory=""
for input_file in "$input_directory"/ordered_*.csv; do
    n_lines=$(wc -l < "$input_file")
    file_name=$(basename "$input_file")
    go run parse-stack-distance-hitratio.go -file_path "$input_file" -n_lines "$n_lines" > "$output_directory"/"${file_name%????}"_hitratio.csv
    echo "File $input_file has been converted"
done
