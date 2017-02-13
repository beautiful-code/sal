'use strict';

ApplicationService.$inject = [
  '$http',
  'AuthService',
  '$location',
  '$rootScope',
  'ApplicationServiceEndpoint'
];

function ApplicationService($http, AuthService, $location, $rootScope, AppServiceEndpoint) {
  var _this = this;

  var redirectOrContinue = function() {
    if (!AuthService.isAuthed()) {
      $location.path("/");
    }
  };

  _this.createApplication = function(application) {
    redirectOrContinue();

    return $http.post(AppServiceEndpoint + '/applications/create', {
      data: application
    });
  };

  _this.getApplications = function() {
    redirectOrContinue();
    return $http.post(AppServiceEndpoint + '/applications');
  };

  _this.getFeedbacks = function(application) {
    redirectOrContinue();
    return $http.post(AppServiceEndpoint + '/feedbacks', {
      data: application
    });
  };

  _this.createFeedback = function(feedback) {
    redirectOrContinue();
    return $http.post(AppServiceEndpoint + '/feedbacks/create', {
      data: feedback
    });
  };
}

module.exports = ApplicationService;
