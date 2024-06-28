#!/usr/bin/env python3

import os
import utils

BIN_BUILD_PATH = os.path.normpath('build/bin')

utils.make_dir(BIN_BUILD_PATH)

CMD_PATH = os.path.join('.', os.path.normpath(f'./cmd/main.go'))
BIN_PATH = os.path.join('.', BIN_BUILD_PATH, f'scoop{utils.BIN_EXT}')

utils.run(f"go build -o {BIN_PATH} {CMD_PATH}")
