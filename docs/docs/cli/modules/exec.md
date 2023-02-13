# Exec Module

Exec module can be imported from `apic/exec` to execute system commands.

## Functions

### cmd

```js
import cmd from "apic/exec";

const { data, error } = cmd("....");
```

This will execute a system command.

Data key contains system output and error key contains the stderr
