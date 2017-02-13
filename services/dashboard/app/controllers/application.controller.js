'use strict';

ApplicationController.$inject = [
  'ApplicationService',
  '$rootScope',
  '$location',
  '$routeParams'
];

function ApplicationController(AppService, $rootScope, $location, $routeParams) {
  var _this = this;
  _this.feedbacks = [];
  _this.applicationId = parseInt($routeParams.appId);

  var application = function() {
    return {id: _this.applicationId};
  };

  var handleGetFeedbacksSuccessRequest = function(response) {
    if (response.status == 200) {
      _this.feedbacks = response.data.data;
    }
  };

  var handleGetFeedbacksErrorRequest = function(response) {
    if (response.data && response.data.data && response.data.data.error) {
      $rootScope.flashMsg = {
        msg: response.data.data.error.split(";"),
        type: 'danger'
      };
    }
  };

  _this.getFeedbacks = function() {
    AppService.getFeedbacks(application())
      .then(handleGetFeedbacksSuccessRequest, handleGetFeedbacksErrorRequest);
  };

  var handleCreateFeedbackErrorRequest = function(response) {
    _this.formErrors = [];

    if (response.data && response.data.data && response.data.data.error) {
      _this.formErrors = response.data.data.error.split(";");
    } else {
      _this.formErrors.push("Oops! something went wrong.");
    }
  };

  var handleCreateFeedbackSuccessRequest = function(response) {
    if (response.status == 201) {
      $rootScope.flashMsg = {
        msg: 'Created feedback Successfully!',
        type: 'success'
      };

      $location.path('/dashboard/applications/' + _this.applicationId);
    }
  };

  var newFeedback = function() {
    return {
      appid: _this.applicationId,
      desc: _this.feedbackDesc,
      email: _this.feedbackEmail
    };
  };

  _this.createFeedback = function() {
    AppService.createFeedback(newFeedback())
      .then(handleCreateFeedbackSuccessRequest, handleCreateFeedbackErrorRequest);
  };

  // Initialisations
  _this.getFeedbacks();
}

module.exports = ApplicationController;
