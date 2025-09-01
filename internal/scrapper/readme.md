# Scrapper

The XML and the JSON files contain multiple syntax errors and inconsistencies.

The scrapper reports the errors during the process, and the files need to be corrected manually.

Also, some operations are duplicated; in this case, the file names will have a suffix.
Most of those files are useless.

The file names that contain `-` or `_` are ignored by the generator.

`bidAuctions` is suffixed with `-SKIP` because this is the only one that uses `POST` instead of `GET`.
