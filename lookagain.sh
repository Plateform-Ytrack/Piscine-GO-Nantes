#!/bin/bash
find . -name "*.sh"  | cut -d "." -f 2 | rev | cut -d "/" -f1 | rev 	