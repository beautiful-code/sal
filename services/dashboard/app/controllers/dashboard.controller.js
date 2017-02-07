'use strict';

Dashboardcontroller.$inject = ['UserService', 'AuthService'];

function Dashboardcontroller(UserService, AuthService) {
  var _this = this;

  _this.logout = function() {
    return AuthService.logout && AuthService.logout();
  };

  _this.userSignedIn = function() {
    return AuthService.isAuthed? AuthService.isAuthed() : false;
  };
}

module.exports = Dashboardcontroller;
