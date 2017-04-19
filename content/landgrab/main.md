# Land Grab

```
import json
import random

import flask
import requests


APP = flask.Flask(__name__)
RESOURCE = 'http://api.landgrab.jwowillo.com/moves'


@APP.route('/')
def play():
    state = flask.request.args.get('state')
    resp = requests.get(url=RESOURCE+'?state='+state).json()['data']
    play = [random.choice(ms) for ms in resp.values()]
    return flask.jsonify({'data': {'moves': play}})
```
