// Retrieve the angular module
angular.module('dashboard', []);

require('./routes');

// Import the index.js in these directories
require('./directives');
require('./services');
require('./controllers');

// Load the main.scss
require('./stylesheets/main.scss');
