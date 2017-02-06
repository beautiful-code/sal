'use strict';

var angular = require('angular');

angular.module('dashboard')
  .controller('homeController', require('./home.controller'))
  .controller('loginController', require('./login.controller'))
  .controller('registerController', require('./register.controller'))
  .controller('dashboardController', require('./dashboard.controller'));
