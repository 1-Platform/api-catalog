# Schema Case Check

## Desc

1. Rule Name: schema_case_checker
2. Category: Quality

Checks for casing over a request body and query params.

## Options

| Name            | Desc                                            | Required | Type                                              | Default   |
| --------------- | ----------------------------------------------- | -------- | ------------------------------------------------- | --------- |
| req_body_casing | casing system to be followed for request bodies | false    | Enum[kebabcase, snakecase, camelcase, pascalcase] | camelcase |
| params_casing   | casing system to be followed for query params   | false    | Enum[kebabcase, snakecase, camelcase, pascalcase] | camelcase |
