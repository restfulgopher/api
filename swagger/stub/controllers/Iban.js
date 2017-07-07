'use strict';

var url = require('url');

var Iban = require('./IbanService');

module.exports.validateiban = function validateiban (req, res, next) {
  Iban.validateiban(req.swagger.params, res, next);
};
