Model = require '../lib/model'
Utils = require '../lib/utils'

module.exports = class Signup extends Model
  parser: new Utils.SchemaParser(Registration.Schema)
  
  initialize: ->
    @set 'groups', @parser.fieldGroups()
    @set 'config', Registration.Config
  
  bind: ->
    Utils.Validation.chosen()
    Utils.Validation.phoneMask()
    Utils.Validation.form(@parser, Registration.DataHandlers.Signup)