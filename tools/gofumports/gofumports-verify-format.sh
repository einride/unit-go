#!/bin/bash

go_files=$(git ls-files --exclude-standard --cached --others '*.go')

[[ -z $go_files ]] && {
	exit 0
}

non_generated_go_files=$(xargs grep --files-without-match -e "// Code generated" -e "// Generated code" <<<"$go_files")

[[ -z $non_generated_go_files ]] && {
	exit 0
}

not_formatted=$(xargs gofumports -l <<<"$non_generated_go_files")

[[ -z $not_formatted ]] && {
	exit 0
}

echo 'Files not gofumports-ed:'
echo "$not_formatted"
exit 1
