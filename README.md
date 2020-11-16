# LaunchDarkly Code Generation

## Running

Until packing is added to compile in the templates you can test the program by running:
```
go run main.go generate --apiToken api-12345 --projectKey default --outFile ./LDFlags.ts --baseUri "https://app.launchdarkly.com"
```

This will generate a new file `LDFlags.ts` in the current directory.

Each flag's default value will be set to the `OffVariation` under `Defaults`. It currently does not handle the case were no `OffVariation` is set.

## Development

This application uses `cobra` to handle commands. `raymond` for rendering out the template files that are located under `templates`. `strcase` for converting flag keys to a consistent format. In the typescript example it converts them all to `lowerKebabCase`.

## Roadmap
* Ability to pass in your own templates with optional default overrides.

## Ideas
* Generate a changelog for any changes.
