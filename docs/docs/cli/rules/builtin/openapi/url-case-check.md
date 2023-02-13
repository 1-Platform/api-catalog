# URL Casing Check

## Desc

1. Rule name: url_case_checker
2. Category: quality

Checks for URL's casing.

## Options

| Name            | Desc                                         | Required | Type                                              | Default   |
| --------------- | -------------------------------------------- | -------- | ------------------------------------------------- | --------- |
| casing          | casing system to be followed                 | false    | Enum[kebabcase, snakecase, camelcase, pascalcase] | kebabcase |
| base_urls       | base url followed by the schema for all URLs | false    | String[]                                          | [ ]       |
| blacklist_paths | URLs to be ignored                           | false    | String[]                                          | [ ]       |
