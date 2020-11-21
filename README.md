# LaunchDarkly Code Generation

## Basic Usage

Until packing is added to compile in the templates you can test the program by running:
```
go run main.go generate --apiToken api-12345 --projectKey default --outFile ./LDFlags.py --language python
```

This will generate a new file `LDFlags.py` in the current directory.

Each flag's default value will be set to the `OffVariation` under `Defaults`. It currently does not handle the case were no `OffVariation` is set.

## Commands
`generate` - creates a new file with optional local flags.
`validate` - validates that no local flags are present.
## Development

This application uses `cobra` to handle commands. `raymond` for rendering out the template files that are located under `templates`. `strcase` for converting flag keys to a consistent format. In the typescript example it converts them all to `lowerKebabCase`.

## Roadmap
* Ability to pass in your own templates with optional default overrides.

## Ideas
* Generate a changelog for any changes.
