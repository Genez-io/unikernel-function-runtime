from osv.modules import api

api.require('node')

default = api.run(cmdline="--cwd=/test_app/ /libnode.so ./hello.js")
