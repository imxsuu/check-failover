import os
import sys
import subprocess

try:
    proc = subprocess.run(["python3", "server.py"], stdout=subprocess.PIPE, stderr=subprocess.PIPE, universal_newlines=True)
    proc.check_returncode()

except subprocess.CalledProcessError as e:

    print(e.stderr)

