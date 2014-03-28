module.exports = class AppRouter extends Backbone.Router
  routes:
    '': 'form'
 
  initialize: ->
    @content = $("#content .container")
    Backbone.Events.on 'router:navigate', @navigate, @

  render: (view) ->
    @content.html(view.render().el)
    view.trigger "model:bind"

  form: ->
    @render Registration.Views.Form