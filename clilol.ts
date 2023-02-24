const completionSpec: Fig.Spec = {
  name: "clilol",
  description: "a cli for omg.lol",
  subcommands: [
    {
      name: ["create"],
      description: "Create things",
      subcommands: [
        {
          name: ["dns"],
          description: "Create a DNS record",
          args: [
            { name: "name", description: "name of the DNS record" },
            { name: "type", description: "type of the DNS record" },
            { name: "data", description: "data of the DNS record" },
          ],
          options: [
            {
              name: ["--priority", "-p"],
              description: "priority of the DNS record",
              args: [{ name: "priority", default: "0" }],
            },
            {
              name: ["--ttl", "-T"],
              description: "time to live of the DNS record",
              args: [{ name: "ttl", default: "3600" }],
            },
          ],
        },
        {
          name: ["paste"],
          description: "Create or update a paste",
          args: [{ name: "title", description: "title of the paste" }],

          options: [
            {
              name: ["--filename", "-f"],
              description: "file to read paste from (default stdin)",
              args: [{ name: "filename" }],
            },
            {
              name: ["--listed", "-l"],
              description: "create paste as listed (default false)",
            },
          ],
        },
        {
          name: ["purl"],
          description: "Create a PURL",
          args: [
            { name: "name", description: "name of the PURL" },
            { name: "url", description: "URL that the PURL redirects to" },
          ],

          options: [
            {
              name: ["--listed", "-l"],
              description: "create as listed (default false)",
            },
          ],
        },
        {
          name: ["status"],
          description: "Create a status",
          args: [{ name: "text", description: "text of the status" }],
          options: [
            {
              name: ["--emoji", "-e"],
              description: "emoji to add to status (default sparkles)",
              args: [{ name: "emoji" }],
            },
            {
              name: ["--skip-mastodon-post"],
              description: "do not cross-post to Mastodon",
            },
          ],
        },
        {
          name: ["weblog"],
          description: "Create a weblog entry",
          options: [
            {
              name: ["--filename", "-f"],
              description: "file to read entry from (default stdin)",
              args: [{ name: "filename" }],
            },
          ],
        },
      ],
    },
    {
      name: ["delete"],
      description: "Delete things",
      subcommands: [
        {
          name: ["account"],
          description: "Delete information about your account",
          subcommands: [
            {
              name: ["session"],
              description: "Delete a session",
              args: [
                {
                  name: "id",
                  description: "ID of the session to delete",
                  isDangerous: true,
                },
              ],
            },
          ],
        },
        {
          name: ["dns"],
          description: "Delete a DNS record",
          args: [
            {
              name: "id",
              description: "ID of the record to delete",
              isDangerous: true,
            },
          ],
        },
        {
          name: ["paste"],
          description: "Delete a paste",
          args: [
            {
              name: "id",
              description: "ID of the paste to delete",
              isDangerous: true,
            },
          ],
        },
        {
          name: ["purl"],
          description: "Delete a PURL",
          args: [
            {
              name: "id",
              description: "ID of the PURL to delete",
              isDangerous: true,
            },
          ],
        },
        {
          name: ["weblog"],
          description: "Delete a weblog entry",
          args: [
            {
              name: "id",
              description: "ID of the weblog entry to delete",
              isDangerous: true,
            },
          ],
        },
      ],
    },
    {
      name: ["get"],
      description: "Get things",
      subcommands: [
        {
          name: ["account"],
          description: "Get information about your account",
          subcommands: [
            { name: ["info"], description: "Get info about your account" },
            { name: ["name"], description: "Get your account name" },
            { name: ["settings"], description: "Get your account settings" },
          ],
        },
        {
          name: ["address"],
          description: "Get information about an address",
          subcommands: [
            {
              name: ["availability"],
              description: "Get address availability",
              args: [{ name: "address", description: "address to get" }],
            },
            {
              name: ["expiration"],
              description: "Get address expiration",
              args: [{ name: "address", description: "address to get" }],
            },
            {
              name: ["info"],
              description: "Get information about an address",
              subcommands: [
                {
                  name: ["private"],
                  description: "Get private information about an address",
                  args: [{ name: "address", description: "address to get" }],
                },
                {
                  name: ["public"],
                  description: "Get public information about an address",
                  args: [{ name: "address", description: "address to get" }],
                },
              ],
            },
          ],
        },
        { name: ["email"], description: "Get email forwarding address(es)" },
        {
          name: ["now"],
          description: "Get a Now page",
          options: [
            {
              name: ["--address", "-a"],
              description: "address whose Now page to get",
              args: [{ name: "address" }],
            },
            {
              name: ["--filename", "-f"],
              description: "file to write Now page to (default stdout)",
              args: [{ name: "filename" }],
            },
          ],
        },
        {
          name: ["paste"],
          description: "Get a paste",
          args: [{ name: "title", description: "title of the paste" }],

          options: [
            {
              name: ["--address", "-a"],
              description: "address whose paste to get",
              args: [{ name: "address" }],
            },
            {
              name: ["--filename", "-f"],
              description: "file to write paste to (default stdout)",
              args: [{ name: "filename" }],
            },
          ],
        },
        {
          name: ["purl"],
          description: "Get a PURL",
          args: [{ name: "name", description: "name of the PURL" }],

          options: [
            {
              name: ["--address", "-a"],
              description: "address whose PURL to get",
              args: [{ name: "address" }],
            },
          ],
        },
        { name: ["service"], description: "Get service stats" },
        {
          name: ["status"],
          description: "Get status",
          args: [{ name: "id", description: "ID of the status" }],

          options: [
            {
              name: ["--address", "-a"],
              description: "address whose status to get",
              args: [{ name: "address" }],
            },
          ],
        },
        {
          name: ["status-bio"],
          description: "Get status bio",
          options: [
            {
              name: ["--address", "-a"],
              description: "address whose status bio to get",
              args: [{ name: "address" }],
            },
          ],
        },
        {
          name: ["theme"],
          description: "Get theme information",
          args: [{ name: "name", description: "name of the theme" }],

          subcommands: [
            {
              name: ["preview"],
              description: "Get theme preview",
              args: [{ name: "name", description: "name of the theme" }],

              options: [
                {
                  name: ["--filename", "-f"],
                  description: "file to write preview to (default stdout)",
                  args: [{ name: "filename" }],
                },
              ],
            },
          ],
        },
        {
          name: ["web"],
          description: "Get your webpage content",
          options: [
            {
              name: ["--filename", "-f"],
              description: "file to write webpage to (default stdout)",
              args: [{ name: "filename" }],
            },
          ],
        },
        {
          name: ["weblog"],
          description: "Get a weblog entry",
          args: [{ name: "id", description: "ID of the weblog entry" }],

          subcommands: [
            {
              name: ["config"],
              description: "Get your weblog config",
              options: [
                {
                  name: ["--filename", "-f"],
                  description:
                    "file to write configuration to (default stdout)",
                  args: [{ name: "filename" }],
                },
              ],
            },
            { name: ["latest"], description: "Get the latest weblog entry" },
            {
              name: ["template"],
              description: "Get your weblog template",
              options: [
                {
                  name: ["--filename", "-f"],
                  description: "file to write template to (default stdout)",
                  args: [{ name: "filename" }],
                },
              ],
            },
          ],
        },
      ],
    },
    {
      name: ["list"],
      description: "List things",
      subcommands: [
        {
          name: ["account"],
          description: "List information about your account",
          subcommands: [
            { name: ["addresses"], description: "List your addresses" },
            { name: ["sessions"], description: "List your sessions" },
          ],
        },
        { name: ["directory"], description: "List the address directory" },
        { name: ["dns"], description: "List your dns records" },
        { name: ["now"], description: "List Now pages" },
        {
          name: ["paste"],
          description: "List pastes",
          options: [
            {
              name: ["--address", "-a"],
              description: "address whose pastes to list",
              args: [{ name: "address" }],
            },
          ],
        },
        {
          name: ["purl"],
          description: "List all PURLs",
          options: [
            {
              name: ["--address", "-a"],
              description: "address whose PURLs to get",
              args: [{ name: "address" }],
            },
          ],
        },
        {
          name: ["status"],
          description: "List statuses",
          options: [
            {
              name: ["--address", "-a"],
              description: "address whose status(es) to get",
              args: [{ name: "address" }],
            },
            {
              name: ["--limit", "-l"],
              description: "how many status(es) to get (default all)",
              args: [{ name: "limit", default: "0" }],
            },
          ],
        },
        {
          name: ["statuslog"],
          description: "List the statuslog",
          options: [
            {
              name: ["--all", "-A"],
              description:
                "get the entire statuslog (default is latest statuses only)",
            },
          ],
        },
        { name: ["theme"], description: "List profile themes" },
        { name: ["weblog"], description: "List all weblog entries" },
      ],
    },
    {
      name: ["update"],
      description: "Update things",
      subcommands: [
        {
          name: ["account"],
          description: "Update information about your account",
          subcommands: [
            { name: ["name"], description: "set the name on your account" },
            {
              name: ["settings"],
              description: "set the settings on your account",
              options: [
                {
                  name: ["--communication", "-c"],
                  description: "communication preference",
                  args: [{ name: "communication" }],
                },
                {
                  name: ["--date-format", "-d"],
                  description: "date format preference",
                  args: [{ name: "date-format" }],
                },
                {
                  name: ["--web-editor", "-w"],
                  description: "web editor preference",
                  args: [{ name: "web-editor" }],
                },
              ],
            },
          ],
        },
        {
          name: ["dns"],
          description: "Update a DNS record",
          options: [
            {
              name: ["--priority", "-p"],
              description: "updated priority",
              args: [{ name: "priority", default: "0" }],
            },
            {
              name: ["--ttl", "-T"],
              description: "updated TTL",
              args: [{ name: "ttl", default: "3600" }],
            },
          ],
        },
        {
          name: ["email"],
          description: "set email forwarding address(es)",
          options: [
            {
              name: ["--destination", "-d"],
              description: "address(es) to forward to",
              args: [{ name: "destination" }],
            },
          ],
        },
        { name: ["preference"], description: "set a preference" },
        {
          name: ["set"],
          description: "set Now page content",
          options: [
            {
              name: ["--filename", "-f"],
              description: "file to read Now page from (default stdin)",
              args: [{ name: "filename" }],
            },
            {
              name: ["--listed", "-l"],
              description: "create Now page as listed (default false)",
            },
          ],
        },
        {
          name: ["status"],
          description: "Update a status",
          options: [
            {
              name: ["--emoji", "-e"],
              description: "emoji to add to status (default sparkles)",
              args: [{ name: "emoji" }],
            },
          ],
        },
        { name: ["status-bio"], description: "Update your status bio" },
        {
          name: ["web"],
          description: "set webpage content",
          subcommands: [
            { name: ["pfp"], description: "set your profile picture" },
          ],
          options: [
            {
              name: ["--filename", "-f"],
              description: "file to read webpage from (default stdin)",
              args: [{ name: "filename" }],
            },
            {
              name: ["--publish", "-p"],
              description: "publish the updated page (default false)",
            },
          ],
        },
        {
          name: ["weblog"],
          description: "set your weblog config",
          subcommands: [
            {
              name: ["config"],
              description: "set your weblog config",
              options: [
                {
                  name: ["--filename", "-f"],
                  description: "file to read config from (default stdin)",
                  args: [{ name: "filename" }],
                },
              ],
            },
            {
              name: ["template"],
              description: "set your weblog template",
              options: [
                {
                  name: ["--filename", "-f"],
                  description: "file to read template from (default stdin)",
                  args: [{ name: "filename" }],
                },
              ],
            },
          ],
        },
      ],
    },
    {
      name: ["help"],
      description: "Help about any command",
      subcommands: [
        {
          name: ["create"],
          description: "Create things",
          subcommands: [
            { name: ["dns"], description: "Create a DNS record" },
            { name: ["paste"], description: "Create or update a paste" },
            { name: ["purl"], description: "Create a PURL" },
            { name: ["status"], description: "Create a status" },
            { name: ["weblog"], description: "Create a weblog entry" },
          ],
        },
        {
          name: ["delete"],
          description: "Delete things",
          subcommands: [
            {
              name: ["account"],
              description: "Delete information about your account",
              subcommands: [
                { name: ["session"], description: "Delete a session" },
              ],
            },
            { name: ["dns"], description: "Delete a DNS record" },
            { name: ["paste"], description: "Delete a paste" },
            { name: ["purl"], description: "Delete a PURL" },
            { name: ["weblog"], description: "Delete a weblog entry" },
          ],
        },
        {
          name: ["get"],
          description: "Get things",
          subcommands: [
            {
              name: ["account"],
              description: "Get information about your account",
              subcommands: [
                { name: ["info"], description: "Get info about your account" },
                { name: ["name"], description: "Get your account name" },
                {
                  name: ["settings"],
                  description: "Get your account settings",
                },
              ],
            },
            {
              name: ["address"],
              description: "Get information about an address",
              subcommands: [
                {
                  name: ["availability"],
                  description: "Get address availability",
                },
                { name: ["expiration"], description: "Get address expiration" },
                {
                  name: ["info"],
                  description: "Get information about an address",
                  subcommands: [
                    {
                      name: ["private"],
                      description: "Get private information about an address",
                    },
                    {
                      name: ["public"],
                      description: "Get public information about an address",
                    },
                  ],
                },
              ],
            },
            {
              name: ["email"],
              description: "Get email forwarding address(es)",
            },
            { name: ["now"], description: "Get a Now page" },
            { name: ["paste"], description: "Get a paste" },
            { name: ["purl"], description: "Get a PURL" },
            { name: ["service"], description: "Get service stats" },
            { name: ["status"], description: "Get status" },
            { name: ["status-bio"], description: "Get status bio" },
            {
              name: ["theme"],
              description: "Get theme information",
              subcommands: [
                { name: ["preview"], description: "Get theme preview" },
              ],
            },
            { name: ["web"], description: "Get your webpage content" },
            {
              name: ["weblog"],
              description: "Get a weblog entry",
              subcommands: [
                { name: ["config"], description: "Get your weblog config" },
                {
                  name: ["latest"],
                  description: "Get the latest weblog entry",
                },
                { name: ["template"], description: "Get your weblog template" },
              ],
            },
          ],
        },
        {
          name: ["list"],
          description: "List things",
          subcommands: [
            {
              name: ["account"],
              description: "List information about your account",
              subcommands: [
                { name: ["addresses"], description: "List your addresses" },
                { name: ["sessions"], description: "List your sessions" },
              ],
            },
            { name: ["directory"], description: "List the address directory" },
            { name: ["dns"], description: "List your dns records" },
            { name: ["now"], description: "List Now pages" },
            { name: ["paste"], description: "List pastes" },
            { name: ["purl"], description: "List all PURLs" },
            { name: ["status"], description: "List statuses" },
            { name: ["statuslog"], description: "List the statuslog" },
            { name: ["theme"], description: "List profile themes" },
            { name: ["weblog"], description: "List all weblog entries" },
          ],
        },
        {
          name: ["update"],
          description: "Update things",
          subcommands: [
            {
              name: ["account"],
              description: "Update information about your account",
              subcommands: [
                {
                  name: ["name"],
                  description: "set the name on your account",
                  args: [{ name: "name", description: "the name to set" }],
                },
                {
                  name: ["settings"],
                  description: "set the settings on your account",
                },
              ],
            },
            {
              name: ["dns"],
              description: "Update a DNS record",
              args: [
                { name: "id", description: "the ID of the record to update" },
                { name: "name", description: "updated name" },
                { name: "type", description: "updated type" },
                { name: "data", description: "updated data" },
              ],
            },
            {
              name: ["email"],
              description: "set email forwarding address(es)",
              args: [
                { name: "address", description: "address(es) to forward to" },
              ],
            },
            {
              name: ["preference"],
              description: "set a preference",
              args: [
                { name: "item", description: "preferences item to set" },
                { name: "value", description: "value to set" },
              ],
            },
            { name: ["set"], description: "set Now page content" },
            {
              name: ["status"],
              description: "Update a status",
              args: [
                { name: "id", description: "the ID of the status to update" },
                { name: "text", description: "new text for the status" },
              ],
            },
            {
              name: ["status-bio"],
              description: "Update your status bio",
              args: [
                { name: "text", description: "new text for the status bio" },
              ],
            },
            {
              name: ["web"],
              description: "set webpage content",
              subcommands: [
                {
                  name: ["pfp"],
                  description: "set your profile picture",
                  args: [
                    {
                      name: "filename",
                      description: "the filename of the image to set",
                    },
                  ],
                },
              ],
            },
            {
              name: ["weblog"],
              description: "set your weblog config",
              subcommands: [
                { name: ["config"], description: "set your weblog config" },
                { name: ["template"], description: "set your weblog template" },
              ],
            },
          ],
        },
      ],
    },
  ],
  options: [
    { name: ["--help", "-h"], description: "Display help", isPersistent: true },
  ],
};
export default completionSpec;
