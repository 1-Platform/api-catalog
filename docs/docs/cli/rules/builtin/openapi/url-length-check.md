# URL Length Check

## Desc

1. Rule Name: url_length
2. Category: quality

Checks whether API's URL is too long.

Weights are assigned to all dynamic path params like `/url/{dynamic}` that estimate how long are they generally.

## Options

| Name            | Desc                                                                                | Required | Type     | Default |
| --------------- | ----------------------------------------------------------------------------------- | -------- | -------- | ------- |
| weight          | Dynamic path params average length. Ex: `/url/{id}` the average length id could be. | false    | Number   | 5       |
| max_url_length  | maximum allowed url length                                                          | false    | Number   | 75      |
| blacklist_paths | URLs to be ignored                                                                  | false    | String[] | [ ]     |
