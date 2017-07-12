'use strict';

exports.validateiban = function(args, res, next) {
  /**
   * validate iban endpoint
   *
   * iban String International Bank Account Number (IBAN)
   * returns inline_response_200_1
   **/
  var examples = {};
  examples['application/json'] = {
  "valid" : true,
  "iban" : "DE44500105175407324931"
};
  if (Object.keys(examples).length > 0) {
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(examples[Object.keys(examples)[0]] || {}, null, 2));
  } else {
    res.end();
  }
}

