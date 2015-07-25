"use strict";
const express = require('express');
const app = express();

app.use(require('./app-engine-handlers'));

module.exports = app;
