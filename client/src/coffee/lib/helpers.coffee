(() ->
  # IIFE to avoid collisions with other variables
  (->
    # Make it safe to do console.log() always.
    console = window.console = window.console or {}
    method = undefined
    dummy = ->
    methods = ('assert,count,debug,dir,dirxml,error,exception,
               group,groupCollapsed,groupEnd,info,log,markTimeline,
               profile,profileEnd,time,timeEnd,trace,warn').split ','

    console[method] = console[method] or dummy while method = methods.pop()
  )()

  # Put your handlebars.js helpers here.

  Handlebars.registerHelper "render", (field) ->
    switch field.type
      when 'text'
        Registration.Templates.TextField field
      when 'boolean'
        Registration.Templates.CheckboxField field
      when 'radio'
        Registration.Templates.RadioField field
      when 'select'
        Registration.Templates.SelectField field
      else
        ""
  
  Handlebars.registerHelper "get", (property, object) ->
    object[property]

  $.mask.definitions["~"] = "[+-]"
  
  jQuery.validator.addMethod "phone", ((value, element) ->
    @optional(element) or /^\(\d{3}\) \d{3}\-\d{4}( x\d{1,6})?$/.test(value)
  ), "Enter a valid phone number."

  $.fn.serializeObject = ->

    json = {}
    push_counters = {}
    patterns =
      validate  : /^[a-zA-Z][a-zA-Z0-9_]*(?:\[(?:\d*|[a-zA-Z0-9_]+)\])*$/,
      key       : /[a-zA-Z0-9_]+|(?=\[\])/g,
      push      : /^$/,
      fixed     : /^\d+$/,
      named     : /^[a-zA-Z0-9_]+$/

    @build = (base, key, value) ->
      base[key] = value
      base

    @push_counter = (key) ->
      push_counters[key] = 0 if push_counters[key] is undefined
      push_counters[key]++

    $.each $(@).serializeArray(), (i, elem) =>
      return unless patterns.validate.test(elem.name)

      keys = elem.name.match patterns.key
      merge = elem.value
      reverse_key = elem.name

      while (k = keys.pop()) isnt undefined

        if patterns.push.test k 
          re = new RegExp("\\[#{k}\\]$")
          reverse_key = reverse_key.replace re, ''
          merge = @build [], @push_counter(reverse_key), merge

        else if patterns.fixed.test k 
          merge = @build [], k, merge

        else if patterns.named.test k
          merge = @build {}, k, merge

      json = $.extend true, json, merge

    return json
)()