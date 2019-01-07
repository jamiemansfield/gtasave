GtaSave
=======

GtaSave is a library for reading GTA III-era save files, licenced under the MIT License.

## Support

There is currently (partial) support for the following games:

- Grand Theft Auto: San Andreas

  ```go
  var raw []byte
  var save gtasa.Save
  parser.Parse(raw, &save);
  ```
