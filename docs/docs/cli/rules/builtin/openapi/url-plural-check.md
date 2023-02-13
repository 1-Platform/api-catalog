# URL Plural Check

## Desc

1. Rule Name: url_plural_checker
2. Category: quality

Checks whether resources in URL are in singular or plural as specified in config.

Ex:

`/property/:id` would fail for plural check as it should be `/properties/:id`

| Name            | Desc                                         | Required | Type                     | Default  |
| --------------- | -------------------------------------------- | -------- | ------------------------ | -------- |
| type            | Is it singular or plural                     | false    | Enum[ singular, plural ] | singular |
| base_urls       | base url followed by the schema for all URLs | false    | String[]                 | [ ]      |
| blacklist_paths | URLs to be ignored                           | false    | String[]                 | [ ]      |
