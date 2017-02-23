'use strict';

(function() {
  if(typeof jQuery === 'undefined') {
    var headTag = document.getElementsByTagName('head')[0];
    var scriptTag = document.createElement('script');
    scriptTag.type = 'text/javascript';
    scriptTag.src = '//cdnjs.cloudflare.com/ajax/libs/jquery/3.1.1/jquery.js';
    scriptTag.onload = initSalWidget;
    headTag.appendChild(scriptTag);
  } else {
    initSalWidget();
  }

  function initSalWidget() {
    // TODO: Fix this
    // var _this = this;
    var _this = {};

    $(
      '<button type="button" class="sal-feedback__toggle--btn">Feedback</button>' +
      '<div class="sal-feedback__container">' +
        '<div class="sal-feedback__overlay">Thank you for your valuable feedback.</div>' +
        '<a class="sal-feedback__container--close">x</a>' +
        '<form>' +
          '<h3 class="sal-feedback__header">Feedback</h3>' +
          '<div class="sal-feedback__textField">' +
            '<input type="email" required name="email" placeholder="user@example.com">' +
          '</div>' +
          '<div class="sal-feedback__textArea">' +
            '<textarea required name="description" cols="40" placeholder="Leave your feedback here ...">' +
            '</textarea>' +
          '</div>' +
          '<div class="sal-feedback__btn">' +
            '<button type="button" class="submit">Send Feedback</button>' +
            '<button type="button" class="clear">Clear</button>' +
          '</div>' +
        '</form>' +
      '</div>'
    ).appendTo('body');

    _this.$container = $('.sal-feedback__container');
    _this.$toggleBtn = $('.sal-feedback__toggle--btn');
    _this.$overlay = _this.$container.find('.sal-feedback__overlay');
    _this.$closeBtn = _this.$container.find('.sal-feedback__container--close');
    _this.$submitBtn = _this.$container.find('.sal-feedback__btn .submit');
    _this.$resetBtn = _this.$container.find('.sal-feedback__btn .clear');
    _this.$form = _this.$container.find('form');
    _this.submitBtnLoadingText = 'Sending ...';
    _this.requiredFields = ['description', 'email'];
    _this.requiredFieldsInputs = [];
    _this.salEndpoint = "http://localhost:8090";

    for(var i = 0; i < _this.requiredFields.length; i++) {
      var $inputEle = _this.$form.find("[name='" + _this.requiredFields[i] + "']");
      _this.requiredFieldsInputs.push($inputEle);
    }

    _this.validateForm = function() {
      for(var i = 0; i < _this.requiredFieldsInputs.length; i++) {
        if(!_this.requiredFieldsInputs[i].val()) {
          return false;
        }
      }

      return true;
    };

    _this.resetForm = function() {
      if(_this.$form[0]) {
        _this.$form[0].reset();

        for(var i = 0; i < _this.requiredFieldsInputs.length; i++) {
          _this.requiredFieldsInputs[i].removeClass('invalid');
        }

        _this.$submitBtn.html('Send Feedback');
      }
    };

    _this.highlightRequiredFields = function() {
      for(var i = 0; i < _this.requiredFieldsInputs.length; i++) {
        var $inputEle = _this.requiredFieldsInputs[i];

        if(!$inputEle.val()) {
          $inputEle.addClass('invalid');
        } else {
          $inputEle.removeClass('invalid');
        }
      }
    };

    _this.inputValOf = function(inputName) {
      var value;

      $.grep(_this.requiredFieldsInputs, function($input, i) {
        if ($input[0].name == inputName) {
          value = $input.val();
          return;
        }
      });

      return value;
    };

    _this.submitForm = function() {
      if(_this.validateForm()) {
        _this.$submitBtn.addClass('disabled');
        _this.$submitBtn.html(_this.submitBtnLoadingText);
        _this.submitFeedback();
      } else {
        _this.highlightRequiredFields();
      }
    };

    _this.$closeBtn.on('click', function(e) {
      e.preventDefault();
      _this.$container.hide();
      _this.resetForm();
      _this.$overlay.hide();
      _this.$toggleBtn.show();
    });

    _this.$resetBtn.on('click', function(e) {
      e.preventDefault();
      _this.resetForm();
    });

    _this.$submitBtn.on('click', function(e) {
      e.preventDefault();
      _this.submitForm();
    });

    _this.$toggleBtn.on('click', function(e) {
      e.preventDefault();
      _this.$toggleBtn.hide();
      _this.$container.show();
    });

    _this.newFeedbackData = function() {
      var params =  {
        data: {
          appid: 4,
          desc: _this.inputValOf('description'),
          email: _this.inputValOf('email')
        }
      };

      return JSON.stringify(params);
    };

    _this.submitFeedback = function() {
      $.ajax({
        url: _this.salEndpoint + '/feedbacks/create',
        method: 'POST',
        contentType: 'application/json',
        data: _this.newFeedbackData()
      }).then(
        function(res) {
          _this.$overlay.show();
        },
        function(err) {
          console.log('Oops! something went wrong.');
        }
      );
    };
  }
})();
