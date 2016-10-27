#!/usr/bin/env python3
# -*- coding: utf-8 *-*

from bottle import route, get, post, request, run, template, redirect

import os
import time


@get('/')
def index():
    return template("view/main.html")

run(host='0.0.0.0', reloader=True, debug=True, port=8080)
