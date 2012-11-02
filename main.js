var http = require('http');

var options = {
    host: 'news.sdtv.com.tw',
    port: 80,
    path: '/News_menu2.html?today=2009-03-04',
    method: 'POST'
};

var req = http.request(options, function(res) {
    console.log('STATUS: ' + res.statusCode);
    /*
    res.on('data', function(chunk) {
	console.log('BODY: ' + chunk);
    });
    */
});

req.on('error', function(e) {
    console.log('error: ' + e.message);
});

req.end();
