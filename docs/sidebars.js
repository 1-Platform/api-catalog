/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  tutorialSidebar: [
    "introduction",
    "architecture",
    {
      type: "category",
      label: "CLI",
      items: [
        "cli/overview",
        "cli/installation",
        {
          type: "category",
          label: "Rules",
          items: ["cli/rules/what-are-rules", "cli/rules/customizing"],
        },
        {
          type: "category",
          label: "Builtin Rules",
          items: [
            "cli/rules/builtin/overview",
            {
              type: "category",
              label: "OpenAPI",
              items: [
                {
                  type: "autogenerated",
                  dirName: "cli/rules/builtin/openapi", // Generate sidebar slice from docs/tutorials/easy
                },
              ],
            },
          ],
        },
        "cli/rules/user-defined/overview",
        {
          type: "category",
          label: "Modules",
          items: [
            {
              type: "autogenerated",
              dirName: "cli/modules", // Generate sidebar slice from docs/tutorials/easy
            },
          ],
        },
        {
          type: "category",
          label: "CLI Commands",
          items: ["cli/cli-commands/apic-run", "cli/cli-commands/apic-help"],
        },
      ],
    },
    "kudos",
  ],
};

module.exports = sidebars;
