'use strict';

// retrive the angular module and load the controllers
angular.module('dashboard')
  .controller('HomeController', require('./home.controller'))
  .controller('LoginController', require('./login.controller'))
  .controller('RegisterController', require('./register.controller'))
  .controller('DashboardController', require('./dashboard.controller'))
  .controller('ApplicationController', require('./application.controller'));
