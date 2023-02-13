# Strings Module

To check for string casing and some more operations not covered by the JS strings.

## Functions

### isCasing

```js
import { isCasing } from "apic/strings";

const value = "hello_world";
const check = isCasing(value, "snakecase");
```

To check whether given value is a valid provided casing.

The casing can be snakecase, camelcase, pascalcase and kebabcase.

### isPlural

```js
import { isPlural } from "apic/strings";

const value = "properties";
const check = isPlural(value);
```

### isSingular

```js
import { isSingular } from "apic/strings";

const value = "property";
const check = isSingular(value);
```

### pluralize

```js
import { pluralize } from "apic/strings";

const value = pluralize("property");
```

### singular

```js
import { singular } from "apic/strings";

const value = singular("properties");
```
