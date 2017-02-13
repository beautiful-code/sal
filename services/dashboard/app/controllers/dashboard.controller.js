'use strict';

DashboardController.$inject = [
  'ApplicationService',
  '$rootScope',
  '$location'
];

function DashboardController(AppService, $rootScope, $location) {
  var _this = this;
  _this.applications = [];

  var handleGetAppsSuccessRequest = function(response) {
    if (response.status == 200) {
      _this.applications = response.data.data;
    }
  };

  var handleGetAppsErrorRequest = function(response) {
    if (response.data && response.data.data && response.data.data.error) {
      $rootScope.flashMsg = {
        msg: response.data.data.error.split(";"),
        type: 'danger'
      };
    }
  };

  _this.getApplications = function() {
    AppService.getApplications()
      .then(handleGetAppsSuccessRequest, handleGetAppsErrorRequest);
  };

  var handleCreateAppErrorRequest = function(response) {
    _this.formErrors = [];

    if (response.data && response.data.data && response.data.data.error) {
      _this.formErrors = response.data.data.error.split(";");
    } else {
      _this.formErrors.push("Oops! something went wrong.");
    }
  };

  var handleCreateAppSuccessRequest = function(response) {
    if (response.status == 201) {
      $rootScope.flashMsg = {
        msg: 'Created application Successfully!',
        type: 'success'
      };

      $location.path('/dashboard');
    }
  };

  var newApplication = function() {
    return {
      name: _this.applicationName
    };
  };

  _this.createApplication = function() {
    AppService.createApplication(newApplication())
      .then(handleCreateAppSuccessRequest, handleCreateAppErrorRequest);
  };

  // Initialisations
  _this.getApplications();
}

module.exports = DashboardController;
