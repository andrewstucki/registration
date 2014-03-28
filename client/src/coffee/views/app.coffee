View = require '../lib/view'

module.exports = class AppView extends View
  el: 'body'

  events:
    'mouseenter .color-picker': (event) ->
      $(event.currentTarget).animate
        left: "0px"
      , 500
    'mouseleave .color-picker': (event) ->
      $(event.currentTarget).animate
        left: "-170px"
      , 500
    'click .bg-selector': (event) ->
      @$el.removeClass().addClass $(event.currentTarget).data("background")
    'click .close': (event) ->
      $(event.currentTarget).closest(".alert").hide()

  bootstrap: ->
    Registration.Templates = require('../templates')(Handlebars)
    @loadViews()

  loadViews: ->
    FormView = require './form'

    Registration.Views.Form = new FormView()
            
    AppRouter = require '../routers/app'
    Registration.Routers.AppRouter = new AppRouter()
    
    Backbone.history.start()