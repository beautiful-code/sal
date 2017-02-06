// Note: Angular is loaded by default for the entire app
// webpack.config.js so no need to use require('angular')
var ngRoute = require('angular-route');
var angularCSS = require('angular-css');

// Init the angular module
angular.module('dashboard', ['ngRoute', 'angularCSS']);

// Import the index.js in these directories
require('./controllers');
require('./directives');
require('./factories');
require('./routes');
require('./services');
// Load the main.scss
require('./stylesheets/main.scss');
