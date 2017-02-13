'use strict';

angular.module('dashboard')
  .factory('AuthInterceptor', require('./auth.factory'))
  .config(function($httpProvider) {
    $httpProvider.interceptors.push('AuthInterceptor');
  });
