# Land Grab

```
import json
import random

import flask
import requests


APP = flask.Flask(__name__)
API = 'http://api.landgrab.jwowillo.com'


@APP.route('/')
def play():
    state = flask.request.args.get('state')
    token = requests.get(url=API+'/token').json()['data']['token']
    resp = requests.get(
        url=API+'/moves?state='+state,
        headers={'Authorization': token}
    ).json()['data']
    play = [random.choice(ms) for ms in resp.values()]
    return flask.jsonify({'data': {'moves': play}})
```
