'use strict';

AuthService.$inject = [
  '$window',
  'jwtHelper'
];

function AuthService($window, jwtHelper) {
  var _this = this;

  _this.parseJwt = function(token) {
    return jwtHelper.decodeToken(token);
  };

  _this.saveToken = function(token) {
    $window.localStorage['jwtToken'] = token;
  };

  _this.getToken = function() {
    return $window.localStorage['jwtToken'];
  };

  _this.isAuthed = function() {
    var token = _this.getToken();

    if(token) {
      return jwtHelper.isTokenExpired(token);
    } else {
      return false;
    }
  };

  _this.logout = function() {
    $window.localStorage.removeItem('jwtToken');
  };

  _this.currentUser = function() {
    if (_this.isAuthed()) {
      var token = _this.getToken();
      return _this.parseJwt(token);
    } else {
      return {};
    }
  };
}

module.exports = AuthService;
