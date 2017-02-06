var ngRoute = require('angular-route');
var angularCSS = require('angular-css');
var dashboardApp = angular.module('dashboard', ['ngRoute', 'angularCSS']);

// Dynamic page title
dashboardApp.run(['$rootScope', '$route', function($rootScope, $route) {
  $rootScope.$on('$routeChangeSuccess', function() {
    document.title = $route.current.pageTitle + ' | SAL';
  });
}]);

dashboardApp.config(function($routeProvider) {
  $routeProvider
  // route for homepage
    .when('/', {
      templateUrl: 'app/templates/home.html',
      controller: 'homeController',
      css: 'app/stylesheets/home.css',
      pageTitle: 'Home'
    })
    .when('/login', {
      templateUrl: 'app/templates/login.html',
      controller: 'loginController',
      css: 'app/stylesheets/login.css',
      pageTitle: 'Login'
    })
    .when('/register', {
      templateUrl: 'app/templates/register.html',
      controller: 'registerController',
      css: 'app/stylesheets/login.css',
      pageTitle: 'Register'
    })
    .when('/dashboard', {
      templateUrl: 'app/templates/dashboard.html',
      controller: 'dashboardController',
      css: 'app/stylesheets/dashboard.css',
      pageTitle: 'Dashboard'
    });
});
