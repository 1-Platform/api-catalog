# Status Code Checks

## Desc

1. Rule Name: status_code_check
2. Category: quality

Checks whether there is any invalid status code. Also can be used only allow particular set of status code in an API.

## Options

| Name                 | Desc                                     | Required | Type     | Default |
| -------------------- | ---------------------------------------- | -------- | -------- | ------- |
| allowed_status_codes | status code that are only allowed in api | false    | String[] | [ ]     |
