# EmissionsMessageProcessor
## Overview
Basic Message Processor for Emissions Data

This processor will connect to `data.salad.com:5000`, receive binary data via a TCP connection, and parse
the bytes into a standard go struct. The contents of this struct will be printed to the console.
In the future, this could be extended to hand the data off for further processing

## Prerequisites
This package assumes you have the go environment installed correctly and you have initialized the module

## Installation
an executable can be generated via the standard go command:
```
go build -o emissions_processor
```

or in windows:
```
go build -o emissions_processor.exe
```

once an application has been built, run the installed executable.
