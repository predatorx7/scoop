'Utilities'

import os
import sys

def run(cmd):
    'Run command'
    print(f'$ {cmd}')
    os.system(cmd)

def make_dir(path):
    'Create directory if not exist'
    os.makedirs(name=os.path.abspath(path), exist_ok=True)

def for_nt_or(for_nt, for_others):
    'Returns value different for windows and others'
    if os.name == 'nt':
        return for_nt
    return for_others

BIN_EXT = for_nt_or('.exe', '')
PYTHON_EXE = for_nt_or('python', 'python3')
