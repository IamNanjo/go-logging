// Logging package allows creating custom loggers with prefixes.
// The pacage also includes default logging functions that can be used without creating a custom logger.
//
// Part of logging package is also logging/loglevel:
//
//	Initialize logging level by reading case-insensitive environment variable LOGLEVEL.
//	Valid values ordered from least verbose to most verbose are CRITICAL, ERROR, WARN, INFO, DEBUG.
//	Level can also be set programmatically by editing loglevel.Level to any of the supported values.
package logging
