# Templates
`lcg` uses [raymond](https://github.com/aymerick/raymond) for generating templates.

## Required User Defined Variables
```
{{~outNumber "number"}}: "number" types represented
{{~outBool "boolean"}}: "boolean" types represented
{{~outString "string"}}: "string" types represented
{{~outMap "object"}}: "JSON" types represented
{{~boolCase "lower"~}}: Boolean returns, example: true/false
{{~outComment "//" }}: Used to comment local flag markers, example: //LOCAL_LCG_FLAGS_END
```
`boolCase` will accept `lower` or `title`.
## Pre-defined Variables in templates
`flags` - This is a representation passed in from the LaunchDarkly API - [Feature Flags](https://github.com/launchdarkly/api-client-go/blob/master/docs/FeatureFlags.md)
`localFlags` - These are flags passed in from the command line and marshalled to an array of [LocalFlagTemplate](https://github.com/launchdarkly/lcg-private/blob/master/cmd/generate.go#L19-L23)

## Template Helpers
Managing case: `lowerCamelCase`, `snakeCase`

`defaultValue` - parameters, `feature flag`, `quotes`. This helper returns the default off value for the flag, if none exists it uses the value of the last variation. See an example in the node-typescript generator.

`returnCheck` - parameters `feature flag`. Checks for the Default Off Variation for new Environments or returns the last variation.
