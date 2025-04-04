#!/bin/bash

# 1. Capture apt update output and filter for NO_PUBKEY errors
sudo apt update 2>&1 | grep "NO_PUBKEY" > /tmp/apt_key_errors.txt

# 2. Extract the repository URLs from the error messages
if [ -s "/tmp/apt_key_errors.txt" ]; then
  while read line; do
    repo=$(echo "$line" | awk '{print $7}') # Extract the repository URL
    if [[ "$repo" == "http://"* || "$repo" == "https://"* ]]; then
      # 3. List installed packages from the problematic repository
      packages=$(apt-cache policy | grep "$repo" -A 1000 | grep "Installed:" | awk '{print $2}')
      if [[ -n "$packages" ]]; then
        echo "Packages from $repo (signature verification failed):"
        echo "$packages"
      fi
    fi
  done < /tmp/apt_key_errors.txt
#   rm /tmp/apt_key_errors.txt
else
  echo "No signature verification errors detected."
fi
