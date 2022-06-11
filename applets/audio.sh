#!/bin/bash
echo "  "$(amixer get Master |grep % |sed -e 's/\].*//' |sed -e 's/.*\[//')
