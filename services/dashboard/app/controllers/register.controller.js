'use strict';

var _ = require('underscore');
var $ = require('jquery');

RegisterController.$inject = [
  'UserService',
  '$location',
  '$rootScope'
];

function RegisterController(UserService, $location, $rootScope) {
  var _this = this;

  var controllerAlias = 'newUser';

  var formFields = [
    'firstName',
    'lastName',
    'email',
    'password'
  ];

  var user = function() {
    return {
      firstname: _this.firstName,
      lastname: _this.lastName,
      email: _this.email,
      password: _this.password
    };
  };

  var handleSuccessRequest = function(response) {
    if (response.status == 201) {
      $rootScope.flashMsg = {msg: "Successfully created the user! Please use the same credentials to login.", type: 'success'};
      $location.path('/login');
    } else {
      $rootScope.flashMsg = {
        type: 'danger',
        msg: "Oops! something went wrong. Please double check your email and the password."
      };
    }
  };

  var handleErrorRequest = function(response) {
    _this.formErrors = [];

    // TODO: Change the response structure on the backend API
    if (response.data && response.data.data && response.data.data.error) {
      _this.formErrors = response.data.data.error.split(";");
    }
  };

  var formValidation = function() {
    var validForm = false;

    _.each(formFields, function(field){
      var $field = $('[ng-model="'+ controllerAlias + '.' + field + '"]');

      if ($field.val()) {
        $field.removeClass('has-error');
        validForm = true;
      } else {
        validForm = false;
        $field.addClass('has-error');
      }
    });

    return validForm;
  };

  _this.register = function() {
    if (formValidation()) {
      UserService.register(user())
        .then(handleSuccessRequest, handleErrorRequest);
    }
  };
}

module.exports = RegisterController;
