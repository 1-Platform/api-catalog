# URL Similarity Checks

## Desc

1. Rule Name: url_similarity_check
2. Category: Quality

Checks for very similar URLs in a schema

## Options

| Name            | Desc                                                                              | Required | Type     | Default |
| --------------- | --------------------------------------------------------------------------------- | -------- | -------- | ------- |
| weight          | Weight define how much similar are two URLs. 0 < similarity <1 where 1 being same | false    | Number   | 0.9     |
| base_urls       | base url followed by the schema for all URLs                                      | false    | String[] | [ ]     |
| blacklist_paths | URLs to be ignored                                                                | false    | String[] | [ ]     |
