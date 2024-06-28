#!/usr/bin/env python3

import os
import utils
import sys

OTHER_ARGS = ' '.join(sys.argv[1:])

utils.run(f"{utils.PYTHON_EXE} {os.path.normpath('./scripts/build.py')}")

utils.run(os.path.normpath(f'./build/bin/scoop{utils.BIN_EXT} {OTHER_ARGS}'))
