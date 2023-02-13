---
sidebar_position: 1
---

# Overview

APIC comes with some JS modules to help you write your own rule.

You can import those packages under the `apic/` namespace.

List of APIC modules

| Name                                   | Package      | Desc                                 |
| -------------------------------------- | ------------ | ------------------------------------ |
| [Exec Module](/cli/modules/exec)       | apic/exec    | To run system commands               |
| [Env Module](/cli/modules/env)         | apic/env     | To set and get environment variables |
| [Strings Module](/cli/modules/strings) | apic/strings | To check string casing               |

:::caution

Right now apic doen't support importing anything else, even relative JS file.

We are still working on ways to package complex plugins.

Soon will add support to apic to support complex plugins.
:::
