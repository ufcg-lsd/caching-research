#!/bin/bash

input_directory=""

for file in "$input_directory"/*_output.csv; do
    file_name=$(basename "$file")
    ordered_name="$input_directory/ordered_$file_name"

    sort -nr "$file" > "$ordered_name"

    echo "File $file sorted and saved as $ordered_name"
done
