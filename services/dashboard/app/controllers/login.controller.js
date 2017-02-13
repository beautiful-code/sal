'use strict';

var _ = require('underscore');
var $ = require('jquery');

LoginController.$inject = [
  'UserService',
  '$location',
  '$rootScope'
];

function LoginController(UserService, $location, $rootScope) {
  var _this = this;

  var controllerAlias = 'user';

  var formFields = [
    'email',
    'password'
  ];

  var handleSuccessRequest = function(response) {
    if (response.status == 200) {
      $rootScope.flashMsg = {msg: "Successfully logged in!", type: 'success'};
      $location.path('/dashboard');
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
      var err = response.data.data.error;

      if (err == "crypto/bcrypt: hashedSecret too short to be a bcrypted password") {
        _this.formErrors.push('User does not exist. Please create a new account.');
      } else {
        _this.formErrors = err.split(";");
      }
    }
  };

  var user = function() {
    return {
      email: _this.email,
      password: _this.password
    };
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

  _this.login = function() {
    if (formValidation()) {
      UserService.login(user())
        .then(handleSuccessRequest, handleErrorRequest);
    }
  };
}

module.exports = LoginController;
