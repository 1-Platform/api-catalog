# Overview

CLI is a core component of our whole architecture. It is what takes API Catalog closer to the API developer.

It mainly validates your API schema file, runs it against the rules you have specified, and then submits the final report to the server.

## CLI Architecture

![cli-architecture](/img/cli-arch.png)

### Config

The config provides options for teams to customize the rule working, disable some rules and add their own rules. It will soon give options for providing other information like API metadata, changelog files, etc.

The following formats as config file

1. YAML
2. JSON
3. TOML

<details>
<summary>A simple config file.</summary>

```yml
rules:
  url_plural_checker:
    options:
      base_urls:
        - /api/v1
  url_case_checker:
    options:
      base_urls:
        - /api/v1
      casing: kebabcase
plugins:
  rules:
    test_plugin:
      file: "./my-own-rule/index.js"
```

</details>

:::info

Rules section points towards builtin apic rules

Where as plugins sections points towards user defined rules

:::

### Rules

Rules are the validations executed over a schema file.

They can be as simple as URL length checks to API performance testing. Each rule will assign a score to one of the API Catalog measured properties and, if any validation fails, will be reported.

:::info

For more information on rules [head over to rules section.](/cli/rules/what-are-rules)

:::
