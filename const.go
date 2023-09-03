package jsonschema

const (
	keywordSep         = ","
	keywordArraySep    = ";"
	keywordKeyValueSep = "="

	kwDefault          = "default"
	kwFormat           = "format"
	kwPattern          = "pattern"
	kwTitle            = "title"
	kwDescription      = "description"
	kwExample          = "example"
	kwAnchor           = "anchor"
	kwEnum             = "enum"
	kwDeprecated       = "deprecated"
	kwMinLength        = "minLength"
	kwMaxLength        = "maxLength"
	kwExclusiveMinimum = "exclusiveMinimum"
	kwExclusiveMaximum = "exclusiveMaximum"
	kwMinimum          = "minimum"
	kwMaximum          = "maximum"
	kwMinItems         = "minItems"
	kwMaxItems         = "maxItems"
	kwUniqueItems      = "uniqueItems"
	kwMultipleOf       = "multipleOf"
	kwType             = "type"
	kwAnyOfType        = "anyof_type"
	kwOneOfType        = "oneof_type"
	kwAnyOfRequired    = "anyof_required"
	kwOneOfRequired    = "oneof_required"
	kwOmitEmpty        = "omitempty"
	kwNullable         = "nullable"

	kwReadOnly  = "readOnly"
	kwWriteOnly = "writeOnly"

	valueT     = "t"
	valueTrue  = "true"
	valueFalse = "false"

	tagJSON                  = "json"
	tagJSONSchema            = "jsonschema"
	tagJSONSchemaDescription = "jsonschema_description"
	tagJSONSchemaExtras      = "jsonschema_extras"
	tagRequired              = "required"
)

const (
	// TypeString represents the string JSON Schema type.
	// See https://json-schema.org/understanding-json-schema/reference/string.html#string for more information.
	TypeString = "string"

	// TypeNumber represents the numeric JSON Schema type.
	// See https://json-schema.org/understanding-json-schema/reference/numeric.html#number for more information.
	TypeNumber = "number"

	// TypeInteger represents the integer JSON Schema type.
	// See https://json-schema.org/understanding-json-schema/reference/numeric.html#integer for more information.
	TypeInteger = "integer"

	// TypeObject represents the object JSON Schema type.
	// See https://json-schema.org/understanding-json-schema/reference/object.html#object for more information.
	TypeObject = "object"

	// TypeArray represents the object JSON Schema type.
	// See https://json-schema.org/understanding-json-schema/reference/array.html#array for more information.
	TypeArray = "array"

	// TypeBoolean represents the boolean JSON Schema type.
	// See https://json-schema.org/understanding-json-schema/reference/boolean.html#boolean for more information.
	TypeBoolean = "boolean"

	// TypeNull represents the null JSON Schema type.
	// See https://json-schema.org/understanding-json-schema/reference/null.html#null for more information.
	TypeNull = "null"
)

// A built-in string format.
// See https://json-schema.org/understanding-json-schema/reference/string.html#built-in-formats for more information.
const (
	FormatStringDateTime                  = "date-time"
	FormatStringTime                      = "time"
	FormatStringDate                      = "date"
	FormatStringDuration                  = "duration"
	FormatStringEmail                     = "email"
	FormatStringInternationalizedEmail    = "idn-email"
	FormatStringHostname                  = "hostname"
	FormatStringInternationalizedHostname = "idn-hostname"
	FormatStringIPv4                      = "ipv4"
	FormatStringIPv6                      = "ipv6"
	FormatStringUUID                      = "uuid"
	FormatStringURI                       = "uri"
	FormatStringURIReference              = "uri-reference"
	FormatStringURITemplate               = "uri-template"
	FormatStringIRI                       = "iri"
	FormatStringIRIReference              = "iri-reference"
	FormatStringRegex                     = "regex"
	FormatStringJSONPointer               = "json-pointer"
	FormatStringRelativeJSONPointer       = "relative-json-pointer"
)
