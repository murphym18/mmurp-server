const env = process.env;
const port = (env.PORT || env.npm_config_port || env.npm_package_config_port || '8080');
const host = (env.HOST || env.npm_config_host || env.npm_package_config_host || '0.0.0.0');
var app = require('./lib');

app.get('/', function(req, res) {
   res.status(200).send("Hello, world!");
});

var server = app.listen(port, host, function() {
   console.log('App listening at http://%s:%s', server.address().address,
      server.address().port);
   console.log("Press Ctrl+C to quit.");
});
