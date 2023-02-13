# Customizing

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

So let's take the URL case checker.

By default, the **URL case checker** rule validates whether your URL is a kebab case.

#### But wait!

What if your team prefers a snake case or camel case?

APIC has got your team.

Write your customization in the APIC config file and pass it to APIC. APIC will handle the rest for you.

So how does a config file looks? Here you go.

<Tabs>
<TabItem value="yaml" label="apic.yaml" default>

```yaml
rules:
  url_case_checker:
    options:
      casing: snakecase
```

</TabItem>
<TabItem value="json" label="apic.json">

```json
{
  "rules": {
    "url_case_checker": {
      "options": {
        "casing": "snakecase"
      }
    }
  }
}
```

</TabItem>
<TabItem value="toml" label="apic.toml">

```toml
[rules.url_case_checker.options]
casing = "snakecase"
```

</TabItem>
</Tabs>

:::info

By default apic looks for a apic file in the directory your executing the cli.

You can change that by using the `--config` option in the cli

:::

## APIC Config file Definition

```yaml
# builtin rules are accessed with rules key
rules:
  rule_name:
    disable: false # to disable a rule
    options:
      option1: value
# user defined rules are in plugins
plugins:
  rule_name:
    file: ./some/dir
    disable: false
    options:
      option1: value
```
