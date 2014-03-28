# App Namespace
window.Registration ?= {}
Registration.DataHandlers ?= {}
Registration.Routers ?= {}
Registration.Views ?= {}
Registration.Collections ?= {}
Registration.Templates ?= {}
Registration.User ?= {}
Registration.Schema ?= {}
Registration.Config ?= {}

(() ->
  Registration.DataHandlers.Signup = require './handlers/signup'
  Registration.DataHandlers.Schema = require './handlers/schema'
  Registration.DataHandlers.Config = require './handlers/config'

  # Load App Helpers
  require './lib/helpers'
  
  Registration.DataHandlers.Schema.fetch (data) ->
    Registration.Schema = data
    
    Registration.DataHandlers.Config.fetch (data) ->
      Registration.Config = data
      
      AppView = require './views/app'
      Registration.Views.App = new AppView()
)()