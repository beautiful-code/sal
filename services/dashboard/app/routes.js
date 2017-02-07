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
    $location.path('/')
  };

  $rootScope.currentUser = function() {
    return AuthService.currentUser();
  };
}]);

dashboardApp.run(function ($rootScope, $location, AuthService) {
  $rootScope.$on('$routeChangeStart', function (event, next, curr) {
    if (next.$$route) {
      $rootScope.flashMsg = {};
      var securePath = next.$$route.loginRequired;

      if (!AuthService.isAuthed() && securePath) { $location.path('/'); }
    }
  });
});

dashboardApp.config(function($routeProvider) {
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
    }).otherwise({redirectTo: '/'});
});
