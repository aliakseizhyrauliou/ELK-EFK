const express = require('express');
const FluentClient = require('@fluent-org/logger').FluentClient;
const app = express();

const logger = new FluentClient('fluentd', {
  socket: {
    host: 'fluentd',
    port: 24224,
    timeout: 3000, 
  }
});

app.get('/', function(request, response) {
  logger.emit('follow', {from: 'userA', to: 'userB'});
  response.send('Hello World!');
});

app.get('/error', function(req, resp){
    logger.emit("/error was requested", {data: "my_data"});
    resp.send('Hello Error!');
});

const port = process.env.PORT || 3000;
app.listen(port, function() {
  console.log("Listening on " + port);
});