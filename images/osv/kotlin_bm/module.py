from osv.modules import api

api.require('java8')

default = api.run(cmdline="--cwd=/ /java.so -jar hello.jar")
