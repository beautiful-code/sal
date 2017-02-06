// retrieve already created module
var dashboardApp = angular.module('dashboard');

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
      controller: 'HomeController',
      css: 'app/stylesheets/home.css',
      pageTitle: 'Home'
    })
    .when('/login', {
      templateUrl: 'app/templates/login.html',
      controller: 'LoginController',
      css: 'app/stylesheets/login.css',
      pageTitle: 'Login'
    })
    .when('/register', {
      templateUrl: 'app/templates/register.html',
      controller: 'RegisterController',
      css: 'app/stylesheets/login.css',
      pageTitle: 'Register'
    })
    .when('/dashboard', {
      templateUrl: 'app/templates/dashboard.html',
      controller: 'DashboardController',
      css: 'app/stylesheets/dashboard.css',
      pageTitle: 'Dashboard'
    });
});
