# APIC Run

To run APIC cli validation against the api schema file

## Options

| Name    | Flag | Desc                                             | Required | Type            | Default |
| ------- | ---- | ------------------------------------------------ | -------- | --------------- | ------- |
| apiType | -a   | API schema type                                  | true     | Enum[ openapi ] | -       |
| schema  |      | URL or local file containing the API schema file | true     | String          | -       |
| config  |      | path to apic config file                         | false    | String          | ./apic  |
| export  |      | export the final report to the given directory   | false    | String          | -       |
