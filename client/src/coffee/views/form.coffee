View = require '../lib/view'
Signup = require '../models/Signup'

module.exports = class FormView extends View
  template: Registration.Templates.Form
  model: new Signup()