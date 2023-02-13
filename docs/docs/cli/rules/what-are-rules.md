# What are rules

Rules define how you want to check an API schema. They are the validations to be executed against your API based on your schema.

Some of the builtin rules are

1. API URL length check
2. URL resources plural or singular validation
3. Casing checks

Rules are not limited to just quality checks.

Each one can give a validation on one of these categories.

1. Security
2. Quality
3. Performance

### So how do you write a rule?

Simple, it's just a JS file, as shown below.

```js
export default function (config, opts) {
  // config: object containing your schema file, schema type and some functions
  // opts: user passed options for rules customization
  // do the required checks
}
```

Check out [this section to know more about writing your rule.](/cli/rules/user-defined/overview)
