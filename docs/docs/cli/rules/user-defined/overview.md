# Writing your own rule

APIC is quite extensible. Teams can write their own rule and load them on APIC.

These scripts could be on the web or in your local directory. Let's see how to write your own rule.

## Scenario

Consider your team uses some CLI for testing security over your API. Let's see how to integrate that into APIC.

## Rule File

Rules are defined using a **Javascript** file. They contain the schema validation logic.

APIC looks for a default exported function, and it gets executed. To execute system commands from the rule file, APIC provides a module called `apic/exec`.

```js title="my-rule.js"
import cmd from "apic/exec";

export default function (config, opts) {
  const output = cmd(`cli commands ....`); // output is the stdout of the CLI
  // If this is JSON stdout
  // output will be object with data and error as key
  const data = JSON.parse(output.data);
  data.forEach((d) => {
    // check for condition
    // if it fails report it using config.report({});
  });

  // finally set a score out of 100 based on validations that are checked
  config.setScore("security", 100);
}
```

### Some important notes

1. opts argument contains whatever user wants to pass to the rule
2. config contains a couple of things

| Properties | Desc                                                  | Type                                                                                                           |
| ---------- | ----------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| type       | type of api schema file                               | Enum[openapi]                                                                                                  |
| schema     | api schema file                                       | object                                                                                                         |
| setScore   | function to set score by the rule                     | (category, value) => void                                                                                      |
| report     | function to say what went wrong on validation failure | ({message:string, <br/>path: string, <br/>method:string, <br/> metadata:{key:string, value:string}[], })=>void |

## Connect new rule

To connect your new rule with APIC, specify the rule name and the file location in the config file.

The file can be absolute, relative or a web URL. If it's relative URL, remember that it's relative to the point of execution.

```yaml title="apic.yaml"
plugins:
  rules:
    # your rule name
    my_security_cli:
      file: "<rule dir/url>"
      # where you can pass some options to rule
      options:
        option1: something
```

:::caution

All rule names must be in snakecase. This is for rendering purpose of UI in future.

:::

## Run it

Now just run apic, if you have named the config file as apic and placed it in same directory then it will be automatically loaded. If not use `--config` option.

```bash
apic run -a openapi --schema https://petstore3.swagger.io/api/v3/openapi.json --config .
```

#### 🎉🎉🎉🎉 Congrats that's your own first rule.
