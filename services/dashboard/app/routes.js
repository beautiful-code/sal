// retrieve already created module
var dashboardApp = angular.module('dashboard');

// Dynamic page title
dashboardApp.run(['$rootScope', '$route', 'AuthService', '$location', function($rootScope, $route, AuthService, $location) {
  $rootScope.$on('$routeChangeSuccess', function() {
    document.title = $route.current.pageTitle + ' | SAL';
  });

  $rootScope.userSignedIn = function() {
    return AuthService.isAuthed();
  };

  $rootScope.logOut = function() {
    AuthService.logout();
    $rootScope.flashMsg = {
      type: 'success',
      msg: 'Logged out user Successfully!'
    };
    $location.path('/');
  };

  $rootScope.currentUser = function() {
    return AuthService.currentUser();
  };
}]);

dashboardApp.run(['$rootScope', '$location', 'AuthService', '$timeout', function ($rootScope, $location, AuthService, $timeout) {
  $rootScope.$on('$routeChangeStart', function (event, next, curr) {
    if (next.$$route) {
      var securePath = next.$$route.loginRequired;

      if (!AuthService.isAuthed() && securePath) { $location.path('/'); }
    }

    $timeout(function() {
      $rootScope.flashMsg = {};
    }, 15000);
  });
}]);

dashboardApp.config(['$routeProvider', function($routeProvider) {
  $routeProvider
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
      pageTitle: 'Dashboard',
      loginRequired: true
    })
    .when('/dashboard/applications/new', {
      templateUrl: 'app/templates/new_application.html',
      controller: 'DashboardController',
      css: 'app/stylesheets/dashboard.css',
      pageTitle: 'New Application | Dashboard',
      loginRequired: true
    })
    .when('/dashboard/applications/:appId', {
      templateUrl: 'app/templates/application.html',
      controller: 'ApplicationController',
      css: 'app/stylesheets/dashboard.css',
      pageTitle: 'Feedbacks | Application | Dashboard',
      loginRequired: true
    })
    .when('/dashboard/applications/:appId/feedbacks/new', {
      templateUrl: 'app/templates/new_feedback.html',
      controller: 'ApplicationController',
      css: 'app/stylesheets/dashboard.css',
      pageTitle: 'New Feedback | Application | Dashboard',
      loginRequired: true
    }).otherwise({redirectTo: '/'});
}]);
