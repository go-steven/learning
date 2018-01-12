// run it with:
//    npm init
//    npm install log4js --registry=https://registry.npm.taobao.org --save
//    node main.js

var log4js = require('log4js');
log4js.configure({
	appenders: [{
		type: 'console'
	}, {
		type: 'file',
		filename: 'test.log',
		maxLogSize: 1024,
		backups: 3,
		category: 'normal'
	}]
});
var logger = log4js.getLogger('normal');
logger.setLevel('INFO');

logger.info('hello world');